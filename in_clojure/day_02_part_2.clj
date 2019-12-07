(require '[clojure.string :as str])

(def input
  (-> "../inputs/day02.txt"
      (slurp)
      (str/trim)
      (str/split #",")
      (->> (mapv read-string))))

(def target 19690720)

(def ops {1 +, 2 *, 99 :return})

(defn execute [program-init breaking-point]
  (loop [pos 0 program program-init]
    (let [op-id (get program pos)
          op (get ops op-id)]
      (if (or (= op :return)
              (>= (first program) breaking-point))
        program
        (let [a*   (get program (+ 1 pos))
              b*   (get program (+ 2 pos))
              res* (get program (+ 3 pos))
              res  (op (get program a*)
                       (get program b*))]
          (recur (+ pos 4) (assoc program res* res)))))))

(def result
  (first
   (for [noun (range 100)
         verb (range 100)
         :let [tape (execute (assoc input 1 noun 2 verb) target)]
         :when (= target (first tape))]
     {:noun noun
      :verb verb})))

(println "Day 02 - Part 2:" (+ (* 100 (:noun result))
                               (:verb result)))
