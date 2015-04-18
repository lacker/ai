(ns meta.core
  (:gen-class))

; Boson is a Lisp dialect that has eight core built-ins:
; The "data structure" stuff: car, cdr, cons, nil
; The "functional" stuff: lambda, recur, this
; "if" because you need if

(defn beval [expr]
  "beval evaluates a list of Boson code."
  (cons 'beval (cons expr nil)))

; TODO: make blank lines and ^D not die. Make bad syntax just fail.
(defn brepl []
  "brepl runs a Boson repl."
  (print ">>> ")
  (flush)
  (println (beval (read-string (read-line))))
  (recur))

; This just exits immediately. But calling (brepl) from lein repl works. Hm.
(defn -main
  "Let's build a Boson repl."
  [& args]
  (brepl)
  (println "done")
  )
