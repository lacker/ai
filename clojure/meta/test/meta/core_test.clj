(ns meta.core-test
  (:require [clojure.test :refer :all]
            [meta.core :refer :all]))

(deftest beval-test
  (testing "bevaluating some expressions."
    (is (= nil (beval '(car (cons nil nil)))))
    (is (= '(nil) (beval '(cons nil nil))))
    (is (= '(nil) (beval '(car (cons (cons nil nil) (cons nil nil))))))

    (is (thrown? Exception (beval '(cons nil this))))
    ))
