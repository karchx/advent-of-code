module Day2 (parser, solve1) where

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
