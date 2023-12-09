package year2023
import "math"

type Day08 struct{}

func get_direction(a string, step int) (int){
	wrapped_step := step % len(a)
	if a[wrapped_step] == 'R' {
		return 1
	}
	return -1
}

func read_line(line string) (string, string, string) {
	return line[0:3], line[7:10], line[12:15]
}

func (p Day08) PartA(lines []string) any {
	instruction := lines[0]
	map_left := make(map[string]string)
	map_right := make(map[string]string)
	var key, left, right string
	for _, line := range lines[2:len(lines)-1] {
		key, left, right = read_line(line)
		map_left[key] = left
		map_right[key] = right
	}
  
	current := "AAA"
	step := 0
	for current != "ZZZ" {
		direction := get_direction(instruction, step)
		if direction == 1 {
			current = map_right[current]
		} else {
			current = map_left[current]
		}
		step += 1
	}

	return step
}

type Tuple struct {
    step int
    wrapped_step int
}

const epsilon = 1e-20
// Check if float is integral
func isIntegral(val float64) bool {
    return math.Abs(val-float64(int(val))) < epsilon
}

// Greatest common divisor of two numbers
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Least common multiple of two numbers
func lcm(a, b int) int {
	return int(math.Abs(float64(a*b)) / float64(gcd(a, b)))
}

// Find first common hit of two loops defined by start and length
func find_first_common_hit(start1 int, loop_length1 int, start2 int, loop_length2 int) int {
	n := 0
	found := false
	var test_step int
	for !found {
		test_step = n*loop_length1+start1
		flag := true
		n_2 := float64(test_step-start2)/float64(loop_length2)
		if !isIntegral(n_2) {
			flag = false
		}
		if flag {
			found = true
		} else {
			n += 1
		}
	}
	return test_step
}

func (p Day08) PartB(lines []string) any {
	instruction := lines[0]
	map_left := make(map[string]string)
	map_right := make(map[string]string)
	var key, left, right string
	for _, line := range lines[2:len(lines)-1] {
		key, left, right = read_line(line)
		map_left[key] = left
		map_right[key] = right
	}
	var start_nodes []string
	for key, _ := range map_left {
		if key[2] == 'A' {
			start_nodes = append(start_nodes, key)
		}
	}

	step_hits := make([]int, len(start_nodes))
	loop_lengths := make([]int, len(start_nodes))
	for i, node := range start_nodes {
		wrapped_step_map := make(map[string][]Tuple)
		for key, _ := range map_left {
			wrapped_step_map[key] = []Tuple{}
		}
		var loop_start_step int
		step := 0
		current := node
		for current != node || step==0 {
			wrapped_step := step % len(instruction)

			break_loop := false
			for i:=len(wrapped_step_map[current])-1; i>=0; i-- {
				if wrapped_step_map[current][i].wrapped_step == wrapped_step {
					// loop
					break_loop = true
					// find loop start 
					loop_start_step = wrapped_step_map[current][i].step
					break
				}
			}

			if break_loop {
				break
			}

			if current[2] == 'Z' {
				step_hits[i] = step
			}

			wrapped_step_map[current] = append(wrapped_step_map[current], 
				Tuple{step: step, wrapped_step: wrapped_step})

			direction := get_direction(instruction, step)
			if direction == 1 {
				current = map_right[current]
			} else {
				current = map_left[current]
			}

			step += 1
		}
		loop_lengths[i] = step-loop_start_step
	}

	start_i := step_hits[0]
	loop_i := loop_lengths[0]
	for i:=0; i<=len(loop_lengths)-2; i++ {
		start_j := step_hits[i+1]
		loop_j := loop_lengths[i+1]
		start_i = find_first_common_hit(start_i, loop_i, start_j, loop_j)
		loop_i = lcm(loop_i, loop_j)
	}

	return start_i
}
