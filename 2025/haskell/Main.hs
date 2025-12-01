module Main where
import Day1 (parse, gcZero, applyGC, filterSecret)


main :: IO()
main = do
    rawInput <- readFile "../input/day01.txt"
    print $ filterSecret $ applyGC 50 (gcZero $ parse rawInput)
