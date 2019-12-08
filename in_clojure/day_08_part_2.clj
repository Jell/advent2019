(require '[clojure.string :as str])

(def layers
  (->> "../inputs/day08.txt"
       (slurp)
       (str/trim)
       (map (comp read-string str))
       (partition 25)
       (partition 6)))

(defn color [& [p & others]]
  (case p
    0 " "
    1 "#"
    2 (recur others)))

(println "Day 08 - Part 2: ")
(->> layers
     (map flatten)
     (apply map color)
     (partition 25)
     (map #(str/join "" %))
     (str/join "\n")
     (println))
