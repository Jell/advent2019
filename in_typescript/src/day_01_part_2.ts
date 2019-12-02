import { readFileSync } from 'fs';

let input: number[] = readFileSync('../inputs/day01.txt', "utf8").
    split("\n").
    map(i => parseInt(i)).
    slice(0, -1);

let massToFuel = function(mass: number): number {
    let fuel = Math.floor(mass / 3) - 2
    return fuel > 0 ? fuel + massToFuel(fuel) : 0
}

let fuel = input.map(massToFuel).reduce((a, i) => a + i)

console.log("Day 01 - Part 2:", fuel)
