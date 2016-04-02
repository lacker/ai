#lang racket

; merge two streams, removing dupes, keeping them sorted
(define (merge a b)
  (cond
    [(stream-empty? a) b]
    [(stream-empty? b) a]
    [(= (stream-first a) (stream-first b))
     (stream-cons (stream-first a) (merge (stream-rest a) (stream-rest b)))]
    [(< (stream-first a) (stream-first b))
     (stream-cons (stream-first a) (merge (stream-rest a) b))]
    [(stream-cons (stream-first b) (merge a (stream-rest b)))]
    ))

(define (trimerge a b c) (merge (merge a b) c))

(define (smult k s)
  (cond
    [(stream-empty? s) s]
    [(stream-cons (* k (stream-first s)) (smult k (stream-rest s)))]
    ))

(define (solve)
  (stream-cons
   1
   (trimerge
    (smult 2 (solve))
    (smult 3 (solve))
    (smult 5 (solve))
    )))

(for ([x (solve)])
  (writeln x))
     

