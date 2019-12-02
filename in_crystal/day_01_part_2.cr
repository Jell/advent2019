input = File.read("../inputs/day01.txt").lines.map(&.to_i)

def mass_to_fuel (mass)
  fuel_mass = (mass / 3) - 2
  fuel_mass > 0 ? fuel_mass + mass_to_fuel(fuel_mass) : 0
end

fuel = input.map{ |mass| mass_to_fuel(mass) }.sum

#--
puts "Day 01 - Part 2: #{fuel}"
