module Day2 (parser, solve1, solve2) where

type Input = [(Int, Int)]

parser :: String -> Input
parser = map (range) . lines .
            map (\c -> if c == ',' then '\n' else c) . filter (/= '\n')
        where
            range s =
                let (a, b) = span (/= '-') s
                    b' = drop 1 b
                in (read a, read b')


exprRange :: (Int, Int) -> [Int]
exprRange (a, b) = [a..b]

invalid :: Int -> Bool
invalid n = even len && (f == b)
    where
        cs = (show n)
        len = length cs
        (f, b) = splitAt (len `div` 2) cs


solve1 :: Input -> Int
solve1 s = sum $ filter invalid $ concatMap exprRange s

solve2 :: Input -> Int
solve2 s = sum $ filter invalid2 $ concatMap exprRange s

takes :: Int -> [a] -> [[a]]
takes n = takeWhile (not . null) . map (take n) . iterate (drop n)

factor :: Int -> [Int]
factor n = [f | f <- [1..n `div` 2], n `mod` f == 0]

allSame :: Eq a => [a] -> Bool
allSame [] = True
allSame (x:xs) = all (== x) xs

invalid2 :: Int -> Bool
invalid2 n =
       not $ null [f | f <- factor len, allSame (takes f cs)]
    where
        cs = show n
        len = length cs
