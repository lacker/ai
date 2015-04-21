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
    )

  (testing "basic apply/this"
    (is (= '(nil) (beval '(apply (cons this this) nil))))
    (is (nil? (beval '(apply (car (cons this this)) nil))))
    (is (nil? (beval '(apply (cdr (cons this this)) nil))))
    )
  )
