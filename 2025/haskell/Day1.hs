module Day1 (parse, solve1, solve2) where
import Data.List (isInfixOf, mapAccumL)

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


zeros :: [Int] -> [Int]
zeros = filter (== 0)

solve1 :: [String] -> [Int]
solve1 input = zeros $ applyGC 50 (gcZero input)

solve2 :: [String] -> Int
solve2 input = solve2' $ gcZero input

solve2' :: [Int] -> Int
solve2' = sum . snd . mapAccumL turn2 50

turn :: Int -> Int -> Int
turn a b = (a + b) `mod` 100

zeros' :: Int -> Int -> Int
zeros' a b
    | b > 0 = (a+b) `div` 100
    | otherwise = (a+b) `div` (-100) - a `div` (-100)

turn2 :: Int -> Int -> (Int, Int)
turn2 a b = (turn a b, zeros' a b)
