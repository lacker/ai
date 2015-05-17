(ns regular.core-test
  (:require [clojure.test :refer :all]
            [regular.core :refer :all]))

(deftest regular-test
  (testing "find-all"
    (is (seq? (find-all "foo" "fao")))
    )
  )
