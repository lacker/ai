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

(define (smult s k)
  (cond
    [(stream-empty? s) s]
    [(stream-cons (* k (stream-first s)) (stream-rest s))]
    ))

(for ([x (merge '(1 2 3 5) '(2 4 5 6))])
  (writeln x))
     

