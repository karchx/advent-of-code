module Day3 where

import Data.Array
import Data.Char

type Input = [[Int]]

parser :: String -> Input
parser = map (map digitToInt) . lines

solve1 :: Input -> Int
solve1 input = sum $ map (maxSubNumber 2) input

solve2 :: Input -> Int
solve2 input = sum $ map (maxSubNumber 12) input

maxSubNumber :: Int -> [Int] -> Int
maxSubNumber k ds = maxSubNumbers ds !! (k-1)

maxSubNumbers :: [Int] -> [Int]
maxSubNumbers ds = foldr addDigit [] ds

addDigit :: Int -> [Int] -> [Int]
addDigit d best =
    zipWith max (best++[0]) (zipWith (+) (iterate (*10) d) (0:best))
