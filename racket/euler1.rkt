#lang racket
(require math/number-theory)

(apply +
       (filter (lambda (x) (or (divides? 3 x)
                               (divides? 5 x)))
               (range 1 1000)
               ))

