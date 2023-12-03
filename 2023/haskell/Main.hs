module Main where

import Day1 (trebuchetPartone)

main :: IO()
main = do
  rawInput <- readFile "../inputs/day01.txt"
  print $ trebuchetPartone rawInput
