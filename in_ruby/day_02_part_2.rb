input = File.read("../inputs/day02.txt").split(",").map(&:to_i)

TARGET=19690720

ADD=1
MULT=2
RETURN=99

def execute!(prog,limit)
  prog = prog.clone
  pos = 0
  loop do
    op_id,a_star,b_star,res_star = prog[pos .. pos+3]
    break if op_id == RETURN
    break if prog[0] > limit
    op = case op_id
         when ADD  then :+.to_proc
         when MULT then :*.to_proc
         end
    prog[res_star] = op[prog[a_star], prog[b_star]]
    pos+=4
  end
  prog
end

def find_noun_and_verb(input, target)
  99.times do |noun|
    99.times do |verb|
      prog = input.clone
      prog[1] = noun
      prog[2] = verb
      result = execute!(prog, target)
      if result[0] == target
        return [noun, verb]
      end
    end
  end
end

noun, verb = find_noun_and_verb(input, TARGET)

puts "Day 02 - Part 2: #{100*noun + verb}"
