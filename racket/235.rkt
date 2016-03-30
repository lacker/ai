#lang lazy

; merge two lists, keeping them sorted
(define (merge a b)
  (cond
    [(empty? a) b]
    [(empty? b) a]
    [(< (first a) (first b)) (cons (first a) (merge (rest a) b))]
    (cons (first b) (merge a (rest b)))))


(merge '(1 3 5) '(2 4 6))
     

