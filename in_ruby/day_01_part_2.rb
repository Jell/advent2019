input = File.read("../inputs/day01.txt").lines.map(&:to_i)

mass_to_fuel = -> (mass) {
  fuel_mass = mass.div(3) - 2
  if fuel_mass > 0
    fuel_mass + mass_to_fuel[fuel_mass]
  else
    0
  end
}

fuel = input.map(&mass_to_fuel).reduce(&:+)

#--
puts "Day 01 - Part 2: #{fuel}"
