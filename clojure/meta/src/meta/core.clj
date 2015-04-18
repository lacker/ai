(ns meta.core
  (:gen-class))

; Boson is a Lisp dialect that has eight core built-ins:
; The "data structure" stuff: car, cdr, cons, nil
; The "functional" stuff: lambda, recur, this
; "if" because you need if

(defn beval [expr]
  "beval evaluates a list of Boson code."
  (cond
    (= 'nil expr) nil
    true (cons 'could-not-beval (cons expr nil))))

; TODO: make blank lines and ^D not die. Make bad syntax just fail.
(defn brepl []
  "brepl runs a Boson repl."
  (print ">>> ")
  (flush)
  (println (beval (read-string (read-line))))
  (recur))

(defn -main
  "Let's build a Boson repl."
  [& args]
  (brepl)
  (println "done")
  )
