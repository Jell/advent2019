(require '[clojure.string :as str])

(def input
  (-> "../inputs/day02.txt"
      (slurp)
      (str/trim)
      (str/split #",")
      (->> (mapv read-string))))

(def ops {1 +, 2 *, 99 :return})

(defn execute [program-init]
  (loop [pos 0 program program-init]
    (let [op-id (get program pos)
          op (get ops op-id)]
      (if (= op :return)
        program
        (let [a*   (get program (+ 1 pos))
              b*   (get program (+ 2 pos))
              res* (get program (+ 3 pos))
              res  (op (get program a*)
                       (get program b*))]
          (recur (+ pos 4) (assoc program res* res)))))))

(def result
  (first
   (execute (assoc input
                   1 12
                   2 2))))

(println "Day 02 - Part 1:" result)
