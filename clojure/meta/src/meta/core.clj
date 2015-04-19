(ns meta.core
  (:gen-class))

; Boson is a Lisp dialect that has eight core built-ins:
; The "data structure" stuff: car, cdr, cons, nil
; The "functional" stuff: lambda, recur, this
; "if" because you need if

; It might be useful to add "define", and consider there to be a
; global namespace.

(defn bthrow [message]
  (throw (Exception. message)))

; TODO: (car (cons (cons nil nil) (cons nil nil))) should work but doesn't
; TODO: have real tests
(defn beval [expr]
  "Evaluates some Boson code."
  (cond
    (= 'nil expr) nil
    (list? expr) (let [op (first expr)
                       args (rest expr)]
                   (cond

                     (= 'if op) (if (= 3 (count args))
                                  (if (beval (first args))
                                    (beval (nth args 1))
                                    (beval (nth args 2)))
                                  (bthrow "if must have 3 args"))

                     (= 'car op) (if (= 1 (count args))
                                   (let [arg (beval (first args))]
                                     (if (list? arg)
                                       (first arg)
                                       (bthrow (str "can't car " arg))))
                                   (bthrow "car must have 1 arg"))

                     (= 'cdr op) (if (= 1 (count args))
                                   (let [arg (beval (first args))]
                                     (if (list? arg)
                                       (rest arg)
                                       (bthrow "can only cdr a list")))
                                   (bthrow "cdr must have 1 arg"))

                     (= 'cons op) (if (= 2 (count args))
                                    (let [x (beval (first args))
                                          y (beval (nth args 1))]
                                      (cons x y)))
                                        
                     :else (bthrow "unknown op")))
    :else (bthrow "unhandled case")))

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
