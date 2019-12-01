(require '[clojure.string :as str])

(defn mass->fuel [mass]
  (let [fuel-mass (- (quot mass 3) 2)]
    (if (> fuel-mass 0)
      (+ fuel-mass (mass->fuel fuel-mass))
      0)))

(->> "../inputs/day01.txt"
     (slurp)
     (str/split-lines)
     (map read-string)
     (map mass->fuel)
     (reduce +)
     (str "Day 01 - Part 2: ")
     (println))
