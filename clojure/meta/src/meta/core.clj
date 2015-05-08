(ns meta.core
  (:gen-class))

; Boson is a Lisp dialect that is designed so that it is simple to
; automatically generate valid Boson code.

; Each list of Boson code starts with a keyword that determines how
; that list is evaluated.

; Boson has a small number of core keywords:
; The "data structure" stuff: car, cdr, cons, nil
; The "functional" stuff: call, this, loop
; "if" because you need if

; The only particularly tricky one is "loop".
; (loop f x) macro-expands to
; (call (if (cdr this) (loop f (cdr this)) (car this)) (call f x))

(defn bthrow [message]
  (throw (Exception. message)))

(defn loop-expand
  "Expands a loop macro. f and x are uneval'd Boson code."
  ([f x]
   `(~'call (if (~'cdr ~'this) (~'loop ~f (~'cdr ~'this)) (~'car ~'this))
            (~'call ~f ~x))
   ))

(defn beval
  "Evaluates some Boson code.
  Takes the code itself, the data bound to 'this', and the recursive
  depth to permit (for infinite loop prevention)."
  ([expr]
   (beval expr "no binding for 'this'"))
  ([expr this]
   (beval expr this 100))
  ([expr this depth]
   (cond
     (= 'nil expr) nil
     (= 'this expr) (if (string? this)
                      (bthrow this)
                      this)
     (<= depth 0) (bthrow "recursive depth exceeded")

     (seq? expr)
     (let [op (first expr)
           args (next expr)
           subdepth (- depth 1)
           ]
       (cond
         
         (= 'if op) (if (= 3 (count args))
                      (if (beval (first args) this subdepth)
                        (beval (nth args 1) this subdepth)
                        (beval (nth args 2) this subdepth))
                      (bthrow "if must have 3 args"))

         (= 'car op) (if (= 1 (count args))
                       (let [arg (beval (first args) this subdepth)]
                         (if (seq? arg)
                           (first arg)
                           (bthrow (str "can't car " arg
                                        " of type "
                                        (type arg)))))
                       (bthrow "car must have 1 arg"))

         (= 'cdr op) (if (= 1 (count args))
                       (let [arg (beval (first args) this subdepth)]
                         (if (seq? arg)
                           (next arg)
                           (bthrow "can only cdr a list")))
                       (bthrow "cdr must have 1 arg"))

         (= 'cons op) (if (= 2 (count args))
                        (let [x (beval (first args) this subdepth)
                              y (beval (nth args 1) this subdepth)]
                          (cons x y))
                        (bthrow "can only cons two args"))
         
         (= 'call op) (if (= 2 (count args))
                        (let [func (first args)
                              subthis (beval
                                       (nth args 1) this subdepth)]
                          (beval func subthis subdepth))
                        (bthrow "can only call two args"))
         
         (= 'loop op) (if (= 2 (count args))
                        (beval (apply loop-expand args) this subdepth)
                        (bthrow "can only loop two args"))
         
         :else (bthrow "unknown op")))
     
     :else (bthrow "unhandled case"))
   ))

(defn safe [f & args]
  "Turns exceptions into strings in the arg."
  (try
    (apply f args)
    (catch Exception e (str "exception: " (.getMessage e)))))

(defn cross-product
  "Lazily lists all valid Boson expressions that are cons'd with one
  expression from each of the expression-list arguments."
  ([xs]
   (for [x xs] [x]))
  ([xs ys]
   (for [x xs y ys] [x y]))
  ([xs ys zs]
   (for [x xs y ys z zs] [x y z]))
   )

(defn compositions [n len]
  "Lazily lists all length-len lists of numbers summing to n."
  (cond
    (< len 1) (bthrow "can't compose with no length")
    (< n len) []
    (= len 1) [[n]]
    (> len 1) (for [k (range 1 n)
                    subcomp (compositions (- n k) (- len 1))]
                (cons k subcomp))
    ))

; TODO: test
(defn bcode-with-fragment [fragment others]
  "Lazily lists valid Boson expressions that contain the provided
  fragment as an immediate child of the root. 'others' is a sequence
  of other expressions that can be immediate children."
  (concat
   [(list 'car fragment) (list 'cdr fragment)] ; 1-arg cases
   (bthrow "TODO: more than 1 arg cases")
   ))

(defn bcode-for-size [size lookup]
  "Lazily lists all valid Boson expressions of a particular size.
   Boson expressions are ordered lexicographically on:
   1. size (number of keywords)
   2. the first token, alphabetically
   3. recursively on the args, in order
   'lookup' provides a lookup table of which bcode of smaller sizes to
   use as component pieces."
  (cond
    (< size 1) []
    (= size 1) ['this 'nil]
    :else (mapcat
           (fn [[keyword arglen]]
             (for [comp (compositions (- size 1) arglen)
                   args (apply cross-product (map #(get lookup % []) comp))]
               (cons keyword args)
               )
             )
           [['call 2]
            ['car 1]
            ['cdr 1]
            ['cons 2]
            ['if 3]
            ['loop 2]
            ])
    ))

(defn all-bcode []
  "Lazily lists all valid Boson expressions."
  (mapcat first
          (iterate
           (fn [[prev-results size lookup]]
             (let [results (bcode-for-size size lookup)]
               [results (+ 1 size) (assoc lookup size results)]
               ))
           [[] 1 {}]
           )))

(defn all-bcode-sorted-by [sortfn]
  "Lazily lists all valid Boson expressions, sorted by sortfn."
  (bthrow "TODO: implement")
  )

(defn now
  "Current time as a float."
  [] (/ (System/currentTimeMillis) 1000.0))

(defn bfind
  "Finds some bcode that satisfies the predicate.
  Returns a map with
  :bcode - the code
  :time  - the amount of time spent to find it
  :count - the number of expressions tested to find it"
  ([pred] (bfind pred false))
  ([pred verbose] (bfind pred (all-bcode) 0 (now) verbose))
  ([pred codegen num-tested start-time verbose]
   (let [new-num-tested (+ 1 num-tested)]
     (if (pred (first codegen))
       {:bcode (first codegen)
        :time (- (now) start-time)
        :count new-num-tested}
       (recur pred (rest codegen) new-num-tested start-time verbose))
   )))

; TODO: limit by time, print out debug info
(defn solve-io [iolist]
  "iolist is a vector of 2-vectors listing pairs that we want a function
  to implement. solve-io finds a function that takes the first
  element of each of these pairs to the second."
  (bfind
   (fn [expr]
     (every?
      (fn [[input output]] (= output (safe beval expr input)))
      iolist)
  )))

(def gauntlet
  "The gauntlet is a collection of input-output puzzles that define
  common functions. We should keep adding to the gauntlet to make the
  system more intelligent."
  (let [t [nil]]
    (sorted-map
     :always-false [[nil nil]
                    [t nil]
                    [(cons t nil) nil]
                    [(cons nil t) nil]]
     :always-true [[nil t]
                   [t t]
                   [(cons t nil) t]
                   [(cons nil t) t]]
     :append-nil [[nil [nil]]
                  [[nil] [nil nil]]
                  [[t] [t nil]]]
     :prepend-nil [[nil [nil]]
                   [[nil] [nil nil]]
                   [[t] [nil t]]]
     :reverse [[nil nil]
               [t t]
               [[nil nil] [nil nil]]
               [[t nil] [nil t]]]
     )))

(defn brepl []
  "Runs a Boson repl."
  (print ">>> ")
  (flush)
  (println (safe beval (read)))
  (recur))

(defn -main [& args]
  (println "running the gauntlet")
  (doall (for [[name iolist] gauntlet]
           (do (println name)
               (println (solve-io iolist))
               )))
  (println "done"))
