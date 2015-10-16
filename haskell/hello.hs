#!/usr/bin/env runhaskell

-- Solving the interview problem of "print out only numbers that have
-- only 2,3,5 as factors"

-- merge two sorted lists to get another sorted list
merge :: (Ord a) => [a] -> [a] -> [a]
merge xs [] = xs
merge [] ys = ys
merge (x:xs) (y:ys) = if x < y then x:(merge xs (y:ys)) else y:(merge (x:xs) ys)
