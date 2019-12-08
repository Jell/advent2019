(require '[clojure.string :as str])

(def layers
  (->> "../inputs/day08.txt"
       (slurp)
       (str/trim)
       (map (comp read-string str))
       (partition 25)
       (partition 6)))

(def least-zeros-layer
  (->> layers
       (map flatten)
       (sort-by (fn [layer] (get (frequencies layer) 0)))
       (first)))

(def checksum
  (let [freqs (frequencies least-zeros-layer)]
    (* (get freqs 1)
       (get freqs 2))))

(println "Day 08 - Part 1:" checksum)
