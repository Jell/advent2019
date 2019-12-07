input = File.read("../inputs/day02.txt").split(",").map(&.to_i)

ADD=1
MULT=2
RETURN=99

def execute!(prog)
  prog = prog.clone
  pos = 0
  loop do
    op_id,a_star,b_star,res_star = prog[pos .. pos+3]
    break if op_id == RETURN
    op = case op_id
         when ADD  then -> (x : Int32, y : Int32) { x + y }
         when MULT then -> (x : Int32, y : Int32) { x * y }
         else raise "oops"
         end
    prog[res_star] = op.call(prog[a_star], prog[b_star])
    pos+=4
  end
  prog
end

input[1]=12
input[2]=2
result = execute!(input)

puts "Day 02 - Part 1: #{result[0]}"
