(ns regular.core
  (:gen-class))

; Tools for finding regular expressions.

; Interesting note:
; http://www.regexmagic.com/autogenerate.html
; claims it is impossible to intelligently figure out a regex given
; some strings that a human wants to match.

; For working with them, we represent regexes as Clojure expressions,
; calling this the "expressive" form.
; A string represents a literal matcher.
; A vector represents matching things consecutively.
; A set represents a set of characters.
; A list (repeat n m regex) represents repeating a regex.
; These four types of regular expressions we call the "basic" ones.

(defn re-to-string [re]
  "Creates a typical regex string from the expressive form."
  (cond
    (string? re) re
    (vector? re) (clojure.string/join "" (map re-to-string re))
    :else (throw (Exception. "cannot re-to-string this"))
    ))

; TODO: implement
(defn tight-basics [s]
  "Finds all the basic regular expressions that tightly map to this
   string. A tight mapping is one that is locally minimal - no
   character could be removed from a group and have it still match,
   and the repeat operator can only occur in comma-less form."
  [])

(defn -main
  "Doesn't do anything."
  [& args]
  (println "Hello regex world."))
