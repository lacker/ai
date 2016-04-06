#lang racket
(require math/number-theory)
(require racket/stream)

; Returns a number in possible that divides n, or #f if there is none
(define (find-divisor possibles n)
  (findf (lambda (x) (divides? x n)) possibles))

(define memo (make-hash))

; Returns an ascending list of primes up to and possibly including n
; Uses memo for memoization
(define (primes-up-to n)
  (let ([answer
         (cond
           [(hash-ref memo n #f) (hash-ref memo n)]
           [(< n 2) '()]
           [(let* ([front (primes-up-to (- n 1))]
                   [divisor (find-divisor front n)])
              (if divisor front (append front (list n)))
              )])])
    (hash-set! memo n answer)
    answer))

; Returns some possible divisors for n.
; If n has any divisors, something in this list will be a divisor.
(define (possible-divisors n)
  (append (primes-up-to (integer-sqrt n)) (list n)))

; Return a list of the factorization of n
(define (factor n)
  (if (< n 2)
      '()
      (let ([divisor (find-divisor (possible-divisors n) n)])
        (cons divisor (factor (/ n divisor)))
        )))

; The nth triangle number
(define (trianglify x)
  (/ (* x (+ x 1)) 2))

; Streams the triangle numbers
(define (triangles)
  (stream-map trianglify (in-naturals 1)))

; Returns a hash from item to # of times it occurs
(define (count-duplicates items)
  (let ([h (make-hash)])
    (for ([item items])
      (hash-set! h item (+ 1 (hash-ref h item 0))))
    h))

(define (num-divisors n)
  (for/product ([(p exponent) (count-duplicates (factor n))])
    (+ 1 exponent)))


(for/first ([triangle (triangles)]
            #:when (< 500 (num-divisors triangle)))
  triangle)

