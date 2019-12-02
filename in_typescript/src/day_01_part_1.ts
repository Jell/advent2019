import { readFileSync } from 'fs';

let input: number[] = readFileSync('../inputs/day01.txt', "utf8").
    split("\n").
    map(i => parseInt(i)).
    slice(0, -1);

let fuel = input.map(i => Math.floor(i / 3) - 2).reduce((a, i) => a + i)

console.log("Day 01 - Part 1:", fuel)
