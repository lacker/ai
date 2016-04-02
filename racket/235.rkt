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

; rop is the operator of which our answer is the fixed point
(define (rop s)
  (stream-cons
   1
   (trimerge
    (smult 2 s)
    (smult 3 s)
    (smult 5 s)
    )))

(letrec ([fixed-point (rop fixed-point)])
  (for ([x fixed-point])
    (writeln x)))
     

