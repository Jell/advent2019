#!/usr/bin/env stack --resolver lts-14.14 runghc
import Data.List (group)

splitOn :: Char -> String -> (String, String)
splitOn c "" = ("","")
splitOn c (x : xs)
        | x == c = ("", xs)
        | otherwise = (x : before, after)
        where (before, after) = splitOn c xs

parseInts :: String -> [Int]
parseInts = map (read . pure) . filter (`elem` "-0123456789")

digitRange :: Int -> [Int]
digitRange n = [n..9]

hasDoubles :: [Int] -> Bool
hasDoubles = any ((== 2) . length) . group

validPasswords :: [Int] -> [Int] -> [[Int]]
validPasswords start end = do
  p1 <- digitRange 0
  p2 <- digitRange p1
  p3 <- digitRange p2
  p4 <- digitRange p3
  p5 <- digitRange p4
  p6 <- digitRange p5
  let v = [p1, p2, p3, p4, p5, p6]
  if v >= start && v <= end && hasDoubles v
  then return v
  else []

main :: IO ()
main = do
  raw <- readFile "../inputs/day04.txt"
  let (startStr, endStr) = splitOn '-' raw
      start = parseInts startStr
      end = parseInts endStr
  putStr "Day 04 - Part 2: "
  putStrLn $ show $ length $ validPasswords start end
