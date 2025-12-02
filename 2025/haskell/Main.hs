module Main where
import Day1 (parse, solve1, solve2)
import qualified Day2 as D2

main :: IO()
main = do
    rawInput <- readFile "../input/day01.txt"
    print $ solve2 (parse rawInput)
    putStrLn "----- Day2 -----"
    rawInput2 <- readFile "../input/day02.txt"
    let input2 = D2.parser rawInput2
    print $ D2.solve1 input2
