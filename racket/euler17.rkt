#lang racket
(require math/number-theory)

(define (wordify n)
  (cond
    [(= n 1) "one"]
    [(= n 2) "two"]
    [(= n 3) "three"]
    [(= n 4) "four"]
    [(= n 5) "five"]
    [(= n 6) "six"]
    [(= n 7) "seven"]
    [(= n 8) "eight"]
    [(= n 9) "nine"]
    [(= n 10) "ten"]
    [(= n 11) "eleven"]
    [(= n 12) "twelve"]
    [(= n 13) "thirteen"]
    [(= n 14) "fourteen"]
    [(= n 15) "fifteen"]
    [(= n 16) "sixteen"]
    [(= n 17) "seventeen"]
    [(= n 18) "eighteen"]
    [(= n 19) "nineteen"]
    [(= n 20) "twenty"]
    [(= n 30) "thirty"]
    [(= n 40) "forty"]
    [(= n 50) "fifty"]
    [(= n 60) "sixty"]
    [(= n 70) "seventy"]
    [(= n 80) "eighty"]
    [(= n 90) "ninety"]
    [(= n 1000) "one thousand"]
    [(divides? 100 n) (string-append (wordify (/ n 100)) " hundred")]
    [(< n 100) (string-append (wordify (* 10 (quotient n 10))) " "
                              (wordify (remainder n 10)))]
    [(string-append (wordify (* 100 (quotient n 100))) " and "
                    (wordify (remainder n 100)))]
    ))

;; doesn't count spaces
(define (num-letters phrase)
  (string-length (string-replace phrase " " "")))

(for/sum ([n (in-range 1 1001)])
  (num-letters (wordify n)))

