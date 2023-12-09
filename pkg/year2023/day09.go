package year2023
import "strings"
import "strconv"
//import "fmt"
type Day09 struct{}

func all_zeroes(numbers []int) (bool) {
	for _, num := range numbers {
		if num != 0 {
			return false
		}
	}
	return true
}

func diff_numbers(numbers []int) ([]int) {
	new_numbers := make([]int, len(numbers)-1)
	for i:=1; i<len(numbers); i++ {
		new_numbers[i-1] = numbers[i]-numbers[i-1]
	}
	return new_numbers
}

func find_next(numbers []int) (int) {
	next := numbers[len(numbers)-1]
	loop := true
	for loop {
		if all_zeroes(numbers) {
			return next
		}
		numbers = diff_numbers(numbers)
		next += numbers[len(numbers)-1]
	}
	return -1
}

func read_line09(line string) ([]int) {
	res := make([]int,0)
	for _, char := range strings.Split(line, " ") {
		num, _ := strconv.Atoi(char)
		res = append(res, num)
	}

	return res
}

func (p Day09) PartA(lines []string) any {
	total := 0
	for _, line := range lines[:len(lines)-1] {
		total += find_next(read_line09(line))
	}
	return total
}

func find_previous(numbers []int) (int) {
	next := numbers[0]
	loop := true
	i := 1
	for loop {
		if all_zeroes(numbers) {
			return next
		}
		numbers = diff_numbers(numbers)
		if i % 2 == 0 {
			next += numbers[0]
		} else {
			next -= numbers[0]
		}
		i++
	}
	return -1
}

func (p Day09) PartB(lines []string) any {
	total := 0
	for _, line := range lines[:len(lines)-1] {
		total += find_previous(read_line09(line))
	}
	return total
}
