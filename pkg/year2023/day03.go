package year2023
import "strconv"
import "strings"
type Day03 struct{}

func getNumber(s string, idx int) (int, int, int) {
	var low_end, high_end int
	zero_value := int('0')
	if idx == 0 {
		low_end = 0
	}
	if idx == len(s)-1 {
		high_end = idx
	}
	for i := idx; i <= len(s); i++ {
		if i == len(s) {
			high_end = len(s)
		} else {
			val := int(s[i])-zero_value
			if val < 0 || val > 9 {
				high_end = i
				break
			}
		}
	}
	for i := idx; i >=-1; i-- {
		if i == -1 {
			low_end = 0
		} else {
			val := int(s[i])-zero_value
			if val < 0 || val > 9 {
				low_end = i+1
				break
			}
		}
	}
	number, _ := strconv.Atoi(s[low_end:high_end])
	return number, low_end, high_end
}

func get_char_type(s byte) string {
	zero_value := int('0')
	val := int(s)-zero_value
	if val >= 0 && val <= 9 {
		return "number"
	} else if s == '.' {
		return "dot"
	}
	return "symbol"
}

func next_to_symbol(i int, previous string, current string, next string) bool {
	if i > 0 {
		if get_char_type(current[i-1]) == "symbol" {
			return true
		}
		if get_char_type(previous[i-1]) == "symbol" {
			return true
		}
		if get_char_type(next[i-1]) == "symbol" {
			return true
		}
	}
	if i < len(current)-1 {
		if get_char_type(current[i+1]) == "symbol" {
			return true
		}
		if get_char_type(previous[i+1]) == "symbol" {
			return true
		}
		if get_char_type(next[i+1]) == "symbol" {
			return true
		}
	}
	if get_char_type(previous[i]) == "symbol" {
		return true
	}
	if get_char_type(next[i]) == "symbol" {
		return true
	}
	return false
}

func (p Day03) PartA(lines []string) any {

	//symbol_positions := make(map[int][]int)
	//zero_value := int('0')
	var previous, current, next string
	var high_end int
	var low_end int
	var number int
	var total_sum int
	var ignore bool
	for i := 0; i < len(lines)-1; i++ {
		if i==0{
			previous = strings.Repeat(".", 140)
			current = lines[i]
			next = lines[i+1]
		} else if i==len(lines)-2 {
			next = strings.Repeat(".", 140)
		} else {
			next = lines[i+1]
		}
		for j := 0; j < len(current); j++ {
			char_type := get_char_type(current[j])
			if char_type == "number" {
				number, low_end, high_end = getNumber(current, j)
				ignore = true
				for k := low_end; k < high_end; k++ {
					if next_to_symbol(k, previous, current, next) {
						ignore = false
						break
					}
				}
				if !ignore {
					total_sum += number
				}
				j = high_end
			}
		}
		previous = current
		current = next
	}
		
	return total_sum
}

func get_char_type_partB(s byte) string {
	zero_value := int('0')
	val := int(s)-zero_value
	if val >= 0 && val <= 9 {
		return "number"
	} else if s == '*' {
		return "star"
	}
	return "rest"
}

func getNumber_partB(s string, idx int) (int) {
	var low_end, high_end int
	zero_value := int('0')
	if idx == 0 {
		low_end = 0
	}
	if idx == len(s)-1 {
		high_end = idx
	}
	for i := idx; i <= len(s); i++ {
		if i == len(s) {
			high_end = len(s)
		} else {
			val := int(s[i])-zero_value
			if val < 0 || val > 9 {
				high_end = i
				break
			}
		}
	}
	for i := idx; i >=-1; i-- {
		if i == -1 {
			low_end = 0
		} else {
			val := int(s[i])-zero_value
			if val < 0 || val > 9 {
				low_end = i+1
				break
			}
		}
	}
	number, _ := strconv.Atoi(s[low_end:high_end])
	return number
}

func adj_numbers(i int, previous string, current string, next string) ([]int) {
	var numbers []int
	// current first
	if i > 0 {
		if get_char_type_partB(current[i-1]) == "number" {
			numbers = append(numbers, getNumber_partB(current, i-1))
		}
	}
	if i < len(current)-1 {
		if get_char_type_partB(current[i+1]) == "number" {
			numbers = append(numbers, getNumber_partB(current, i+1))
		}
	}

	// previous
	var number_indices []int
	if i > 0 {
		if get_char_type_partB(previous[i-1]) == "number" {
			number_indices = append(number_indices, i-1)
		}
	}
	if i < len(previous)-1 {
		if get_char_type_partB(previous[i+1]) == "number" {
			number_indices = append(number_indices, i+1)
		}
	}
	if get_char_type_partB(previous[i]) == "number" {
		number_indices = append(number_indices, i)
	}

	if len(number_indices) == 1 {
		numbers = append(numbers, getNumber_partB(previous, number_indices[0]))
	} else if len(number_indices) == 2 {
		if number_indices[0] == i-1 && number_indices[1] == i+1 {
			numbers = append(numbers, getNumber_partB(previous, number_indices[0]))
			numbers = append(numbers, getNumber_partB(previous, number_indices[1]))
		} else {
			numbers = append(numbers, getNumber_partB(previous, number_indices[0]))
		}
	} else if len(number_indices) == 3 {
		numbers = append(numbers, getNumber_partB(previous, number_indices[0]))
	}

	var number_indices2 []int
	// next
	if i > 0 {
		if get_char_type_partB(next[i-1]) == "number" {
			number_indices2 = append(number_indices2, i-1)
		}
	}
	if i < len(next)-1 {
		if get_char_type_partB(next[i+1]) == "number" {
			number_indices2 = append(number_indices2, i+1)
		}
	}
	if get_char_type_partB(next[i]) == "number" {
		number_indices2 = append(number_indices2, i)
	}

	if len(number_indices2) == 1 {
		numbers = append(numbers, getNumber_partB(next, number_indices2[0]))
	} else if len(number_indices2) == 2 {
		if number_indices2[0] == i-1 && number_indices2[1] == i+1 {
			numbers = append(numbers, getNumber_partB(next, number_indices2[0]))
			numbers = append(numbers, getNumber_partB(next, number_indices2[1]))
		} else {
			numbers = append(numbers, getNumber_partB(next, number_indices2[0]))
		}
	} else if len(number_indices2) == 3{
		numbers = append(numbers, getNumber_partB(next, number_indices2[0]))
	}

	return numbers
}

func (p Day03) PartB(lines []string) any {
	var previous, current, next string
	var total_sum int
	for i := 0; i < len(lines)-1; i++ {
		if i==0{
			previous = strings.Repeat(".", 140)
			current = lines[i]
			next = lines[i+1]
		} else if i==len(lines)-2 {
			next = strings.Repeat(".", 140)
		} else {
			next = lines[i+1]
		}
		for j := 0; j < len(current); j++ {
			char_type := get_char_type_partB(current[j])
			if char_type == "star" {
				numbers := adj_numbers(j, previous, current, next)
				if len(numbers) == 2 {
					total_sum += numbers[0]*numbers[1]
				}
			}
		}
		previous = current
		current = next
	}
		
	return total_sum
}
