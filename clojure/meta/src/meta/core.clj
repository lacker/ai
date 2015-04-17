(ns meta.core
  (:gen-class))

; Boson is a Lisp dialect that has eight core built-ins:
; The "data structure" stuff: car, cdr, cons, nil
; The "functional" stuff: lambda, recur, this
; "if" because you need if

(defn beval [expr]
  "beval evaluates a list of Boson code."
  (cons 'beval (cons expr nil)))

(defn brepl []
  "brepl runs a Boson repl."
  nil)

(defn -main
  "Let's build a Boson repl."
  [& args]
  (println ">>> "))
