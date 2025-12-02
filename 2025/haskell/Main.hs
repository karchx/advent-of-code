module Main where
import Day1 (parse, solve1, solve2)


main :: IO()
main = do
    rawInput <- readFile "../input/day01.txt"
    print $ length $ solve2 (parse rawInput)
