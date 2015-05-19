(ns regular.core-test
  (:require [clojure.test :refer :all]
            [regular.core :refer :all]))

(deftest regular-test
  (testing "re-to-string"
    (is (= "foo" (re-to-string "foo")))
    (is (= "foobar" (re-to-string ["foo" "bar"])))
    )
  )
