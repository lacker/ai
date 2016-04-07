#lang racket
(require math/number-theory)

; Maps integer to the length of its Collatz sequence
(define memo (make-hash))

(define (collatz-length n)
  (cond [(< n 2) 1]
        [(hash-ref memo n #f)
         (hash-ref memo n)]
        [(let ([answer
                (if
                 (divides? 2 n)
                 (+ 1 (collatz-length (/ n 2)))
                 (+ 1 (collatz-length (+ 1 (* n 3))))
                 )])
           (hash-set! memo n answer)
           answer)]))

(argmax collatz-length (range 1000000))

             
    
