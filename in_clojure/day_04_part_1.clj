(require '[clojure.string :as str])

(defn str->digits [in]
  (mapv read-string (str/split in #"")))

(let [[input-start input-end]
      (-> "../inputs/day04.txt"
          (slurp)
          (str/trim)
          (str/split #"-"))]

  (def start (str->digits input-start))
  (def end (str->digits input-end)))

(defn in-bound? [v]
  (and (<= (compare v end) 0)
       (>= (compare v start) 0)))

(defn has-doubles? [v]
  (->> v
       (partition-by identity)
       (map count)
       (some #(>= % 2))))

(def valid-passwords
  (for [p1 (range 0 10)
        p2 (range p1 10)
        p3 (range p2 10)
        p4 (range p3 10)
        p5 (range p4 10)
        p6 (range p5 10)
        :let [v [p1 p2 p3 p4 p5 p6]]
        :when (in-bound? v)
        :when (has-doubles? v)]
    [p1 p2 p3 p4 p5 p6]))

(println "Day 04 - Part 1:" (count valid-passwords))
