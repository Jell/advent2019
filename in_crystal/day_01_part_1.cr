input = File.read("../inputs/day01.txt").lines.map(&.to_i)

fuel = input.map { |mass| (mass / 3) - 2 }.sum

#--
puts "Day 01 - Part 1: #{fuel}"
