(ns meta.core-test
  (:require [clojure.test :refer :all]
            [meta.core :refer :all]))

(deftest beval-test
  (testing "basic data expressions."
    (is (= nil (beval '(car (cons nil nil)))))
    (is (= '(nil) (beval '(cons nil nil))))
    (is (= '(nil) (beval '(car (cons (cons nil nil) (cons nil nil))))))
    )
  (testing "things that should throw."
    (is (thrown-with-msg? Exception #"no binding for 'this'"
                          (beval '(cons nil this))))

    (is (thrown? Exception (beval '(cons nil))))
    (is (thrown? Exception (beval '(car nil))))
    (is (thrown? Exception (beval '(cdr nil))))

    ; Nothing should work with no args
    (is (thrown? Exception (beval '(call))))
    (is (thrown? Exception (beval '(car))))
    (is (thrown? Exception (beval '(cdr))))
    (is (thrown? Exception (beval '(cons))))
    (is (thrown? Exception (beval '(if))))
    (is (thrown? Exception (beval '(loop))))
    (is (thrown? Exception (beval '(nil))))
    (is (thrown? Exception (beval '(this))))

    )

  (testing "basic call/this"
    (is (= '(nil) (beval '(call (cons this this) nil))))
    (is (nil? (beval '(call (car (cons this this)) nil))))
    (is (nil? (beval '(call (cdr (cons this this)) nil))))
    (is (nil? (beval '(call (if this nil (cons nil nil)) (cons nil nil)))))
    (is (= '(nil) (beval '(call (if this this nil) (cons nil nil)))))
    (is (nil? (beval '(call (if this nil this) nil))))
    (is (nil? (beval '(call (call this this) nil))))
    )

  (testing "loop expand"
    (is (= '(call (if (cdr this) (loop g (cdr this)) (car this))
                  (call g y))
           (loop-expand 'g 'y)))
    )

  (testing "loop in operation"
    (is (nil? (beval '(loop (cdr this) (cons nil (cons nil nil))))))
    )

  (testing "preventing infinite recursion"
    (is (thrown-with-msg? Exception #"recursive depth exceeded"
                          (beval '(loop (cons nil (cons nil nil)) nil))))
    )

  (testing "cross-product"
    (is (= [[1 3] [1 4] [2 3] [2 4]] (cross-product [1 2] [3 4])))
    (is (= [[1 3 5] [1 3 6] [1 4 5] [1 4 6] [2 3 5] [2 3 6] [2 4 5] [2 4 6]]
           (cross-product [1 2] [3 4] [5 6])))
    )

  (testing "compositions"
    (is (= [[1 2] [2 1]] (compositions 3 2)))
    )

  (testing "bcode-for-size"
    (is (= ['this 'nil] (bcode-for-size 1 {})))
    )
  )
