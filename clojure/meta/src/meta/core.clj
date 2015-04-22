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

(defn beval*
  "Evaluates some Boson code.
  Takes the code itself, the data bound to 'this', and the recursive
  depth to permit (for infinite loop prevention)."
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
                      (if (beval* (first args) this subdepth)
                        (beval* (nth args 1) this subdepth)
                        (beval* (nth args 2) this subdepth))
                      (bthrow "if must have 3 args"))

         (= 'car op) (if (= 1 (count args))
                       (let [arg (beval* (first args) this subdepth)]
                         (if (seq? arg)
                           (first arg)
                           (bthrow (str "can't car " arg
                                        " of type "
                                        (type arg)))))
                       (bthrow "car must have 1 arg"))

         (= 'cdr op) (if (= 1 (count args))
                       (let [arg (beval* (first args) this subdepth)]
                         (if (seq? arg)
                           (next arg)
                           (bthrow "can only cdr a list")))
                       (bthrow "cdr must have 1 arg"))

         (= 'cons op) (if (= 2 (count args))
                        (let [x (beval* (first args) this subdepth)
                              y (beval* (nth args 1) this subdepth)]
                          (cons x y))
                        (bthrow "can only cons two args"))
         
         (= 'call op) (if (= 2 (count args))
                        (let [func (first args)
                              subthis (beval*
                                       (nth args 1) this subdepth)]
                          (beval* func subthis subdepth))
                        (bthrow "can only call two args"))
         
         (= 'loop op) (if (= 2 (count args))
                        (beval* (apply loop-expand args) this subdepth)
                        (bthrow "can only loop two args"))
         
         :else (bthrow "unknown op")))
     
     :else (bthrow "unhandled case"))
   ))

(defn beval
  "Evaluates some Boson code, setting sane defaults."
  [expr]
  (beval* expr "no binding for 'this'" 100))

(defn safe-beval [expr]
  "Evaluates some Boson code and turns exceptions into strings."
  (try
    (beval expr)
    (catch Exception e (str "exception: " (.getMessage e)))))

(defn cross-product
  "Lists all valid Boson expressions that are cons'd with one
  expression from each of the expression-list arguments."
  ([xs]
   (for [x xs] [x]))
  ([xs ys]
   (for [x xs y ys] [x y]))
  ([xs ys zs]
   (for [x xs y ys z zs] [x y z]))
   )

(defn compositions [n len]
  "Lists all length-len lists of numbers summing to n."
  (cond
    (< len 1) (bthrow "can't compose with no length")
    (< n len) []
    (= len 1) [[n]]
    (> len 1) (for [k (range 1 n)
                    subcomp (compositions (- n k) (- len 1))]
                (cons k subcomp))
    ))

(defn bcode-for-size [size lookup]
  "Lists all valid Boson expressions of a particular size.
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
    

; TODO: make blank lines and ^D not die. Make bad syntax just fail.
(defn brepl []
  "Runs a Boson repl."
  (print ">>> ")
  (flush)
  (println (safe-beval (read)))
  (recur))

(defn -main [& args]
  (brepl)
  (println "done")
  )
