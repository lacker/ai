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

; It might be useful to add "define", and consider there to be a
; global namespace.

(defn bthrow [message]
  (throw (Exception. message)))

(defn loop-expand
  "Expands a loop macro. f and x are uneval'd Boson code."
  ([f x]
   `(~'call (if (~'cdr ~'this) (~'loop ~f (~'cdr ~'this)) (~'car ~'this))
            (~'call ~f ~x))
   ))


(defn beval
  "Evaluates some Boson code."
  [expr & {:keys [this] :or {this "no binding for 'this'"}}]
  (cond
    (= 'nil expr) nil
    (= 'this expr) (if (string? this)
                     (bthrow this)
                     this)

    (seq? expr)
    (let [op (first expr)
          args (next expr)]
      (cond
        
        (= 'if op) (if (= 3 (count args))
                     (if (beval (first args) :this this)
                       (beval (nth args 1) :this this)
                       (beval (nth args 2) :this this))
                     (bthrow "if must have 3 args"))

        (= 'car op) (if (= 1 (count args))
                      (let [arg (beval (first args) :this this)]
                        (if (seq? arg)
                          (first arg)
                          (bthrow (str "can't car " arg
                                       " of type "
                                       (type arg)))))
                      (bthrow "car must have 1 arg"))

        (= 'cdr op) (if (= 1 (count args))
                      (let [arg (beval (first args) :this this)]
                        (if (seq? arg)
                          (next arg)
                          (bthrow "can only cdr a list")))
                      (bthrow "cdr must have 1 arg"))

        (= 'cons op) (if (= 2 (count args))
                       (let [x (beval (first args) :this this)
                             y (beval (nth args 1) :this this)]
                         (cons x y))
                       (bthrow "can only cons two args"))
        
        (= 'call op) (if (= 2 (count args))
                       (let [func (first args)
                             subthis (beval
                                      (nth args 1) :this this)]
                         (beval func :this subthis))
                       (bthrow "can only call two args"))

        (= 'loop op) (if (= 2 (count args))
                       (beval (apply loop-expand args) :this this)
                       (bthrow "can only loop two args"))
        
        :else (bthrow "unknown op")))

    :else (bthrow "unhandled case"))
  )

(defn safe-beval [expr]
  "Evaluates some Boson code and turns exceptions into strings."
  (try
    (beval expr)
    (catch Exception e (str "exception: " (.getMessage e)))))

; TODO: make blank lines and ^D not die. Make bad syntax just fail.
(defn brepl []
  "Runs a Boson repl."
  (print ">>> ")
  (flush)
  (println (safe-beval (read-string (read-line))))
  (recur))

(defn -main [& args]
  (brepl)
  (println "done")
  )
