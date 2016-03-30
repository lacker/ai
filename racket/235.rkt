#lang racket

; merge two streams, removing dupes, keeping them sorted
(define (merge a b)
  (cond
    [(stream-empty? a) b]
    [(stream-empty? b) a]
    [(= (stream-first a) (stream-first b))
     (cons (stream-first a) (merge (stream-rest a) (stream-rest b)))]
    [(< (stream-first a) (stream-first b))
     (cons (stream-first a) (merge (stream-rest a) b))]
    [(cons (stream-first b) (merge a (stream-rest b)))]
    ))



(for ([x (merge '(1 2 3 5) '(2 4 5 6))])
  (writeln x))
     

