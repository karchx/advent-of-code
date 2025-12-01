module Day1 (parse, gcZero, applyGC, filterSecret) where
import Data.List (isInfixOf)

splitChar :: Char -> String -> [String]
splitChar _ [] = [""]
splitChar delim (c:cs)
    | c == delim = "" : rest
    | otherwise = (c : head rest) : tail rest
    where rest = splitChar delim cs

parseNumber :: String -> Int
parseNumber s = read s :: Int

parse :: String -> [String]
parse = lines

lastSplit :: Char -> String -> String
lastSplit c x =
    case splitChar c x of
        [] -> ""
        xs -> last xs

dialValidate :: Int -> Int
dialValidate x
    | x < 0 = x `mod` 100
    | x > 99 = dialValidate (x - 100)
    | otherwise = x

gcZero :: [String] -> [Int]
gcZero = map s
    where
        s x
            | "L" `isInfixOf` x = negate (parseNumber (lastSplit 'L' x))
            | "R" `isInfixOf` x = parseNumber (lastSplit 'R' x)

applyGC :: Int -> [Int] -> [Int]
applyGC base xs = scanl step base xs
    where
        step acc x = dialValidate (acc + x)


filterSecret :: [Int] -> Int
filterSecret result = length $ (filter (\x -> x == 0) result)
