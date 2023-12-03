module Day1 (trebuchetPartone) where

import Data.Char (isDigit, digitToInt)

parse :: String -> [String]
parse = lines


trebuchetPartone :: String -> Int
trebuchetPartone x = calibrationValues linesData 
  where
    linesData = parse x

calibrationValues = sum . map (read . parseLine . filter isDigit)
  where
    parseLine l = [head l, last l]
