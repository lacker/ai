(ns meta.core
  (:gen-class))

; Boson is a Lisp dialect that has eight core built-ins:
; The "data structure" stuff: car, cdr, cons, nil
; The "functional" stuff: lambda, recur, this
; "if" because you need if

(defn bthrow [message]
  (throw (Exception. message)))

(defn beval [expr]
  "Evaluates some Boson code."
  (cond
    (= 'nil expr) nil
    (list? expr) (let [op (first expr)
                       args (rest expr)]
                   (cond
                     (= 'if op) (if (= 3 (count args))
                                  (if (beval (first args))
                                    (beval (first (rest args)))
                                    (beval (first (rest (rest args)))))
                                  (bthrow "if must have 3 args"))
                     true (bthrow "unknown op")))
    true (bthrow "unhandled case")))

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
