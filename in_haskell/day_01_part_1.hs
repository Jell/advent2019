#!/usr/bin/env stack --resolver lts-14.14 runghc

parseInt :: String -> Int
parseInt x = read $ filter (`elem` "-0123456789") x

loadInput :: IO ([Int])
loadInput = do
  raw <- readFile "../inputs/day01.txt"
  return $ map parseInt $ lines raw

massToFuel :: Int -> Int
massToFuel m = (m `div` 3) - 2

main :: IO ()
main = do
  input <- loadInput
  putStr "Day 01 - Part 1: "
  putStrLn $ show $ foldl (+) 0 $ map massToFuel input
