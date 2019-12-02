(require '[clojure.string :as str])

(->> "../inputs/day01.txt"
     (slurp)
     (str/split-lines)
     (map read-string)
     (map #(- (quot % 3) 2))
     (reduce +)
     (str "Day 01 - Part 1: ")
     (println))
