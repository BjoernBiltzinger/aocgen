package year2023
import "strings"
import "strconv"
//import "fmt"
type Day04 struct{}

func read_card(card string) ([]int, []int) {
	card_numbers := strings.Split(card, ":")[1]
	number_blocks := strings.Split(card_numbers, "|")
	var list1, list2 []int
	for _, numStr := range strings.Split(number_blocks[0], " ") {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			continue
		}
		list1 = append(list1, num)
	}
	for _, numStr := range strings.Split(number_blocks[1], " ") {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			continue
		}
		list2 = append(list2, num)
	}
	return list1, list2
}

func get_num_matches(winning_numbers []int, my_numbers []int) (int) {
	matches := 0
	for _, num := range winning_numbers {
		for _, num2 := range my_numbers {
			if num == num2 {
				//fmt.Println(num, num2)
				matches++
				break
			}
		}
	}
	return matches
}

func power(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}

func (p Day04) PartA(lines []string) any {
	var winning_numbers []int
	var my_numbers []int
	//var number_blocks []string
	total_sum := 0
	for _, line := range lines[:len(lines)-1] {
		winning_numbers, my_numbers = read_card(line)
		matches := get_num_matches(winning_numbers, my_numbers)

		//total_sum = 1
		//fmt.Println(matches)
		if matches > 0 {
			total_sum += power(2, matches-1)
		}
	}
	return total_sum
}

func (p Day04) PartB(lines []string) any {
	var winning_numbers []int
	var my_numbers []int
	num_copies := make([]int, len(lines)-1)
    for i := 0; i < len(lines)-1; i++ {
        num_copies[i] = 1
    }
	//var number_blocks []string
	for i, line := range lines[:len(lines)-1] {
		winning_numbers, my_numbers = read_card(line)
		matches := get_num_matches(winning_numbers, my_numbers)
		if matches > 0 {
			for j:=i+1; j<=i+matches; j++ {
				if j>=len(lines)-1{
					break
				}
				num_copies[j]+=num_copies[i]
			}
		}
	}
	// Calculate the sum of the array
    sum := 0
    for _, value := range num_copies {
        sum += value
    }
	return sum
}
