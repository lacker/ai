#lang racket
(require memoize)

(define/memo (paths a b)
  (cond
    [(= a 0) 1]
    [(= b 0) 1]
    [(+ (paths (- a 1) b) (paths a (- b 1)))]
    ))

(paths 20 20)
