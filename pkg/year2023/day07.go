package year2023
import "strings"
//import "fmt"
import "strconv"
import "math"
import "sort"
type Day07 struct{}

func get_map() (map[byte]int) {
	m := make(map[byte]int)
	m['2'] = 1
	m['3'] = 2
	m['4'] = 3
	m['5'] = 4
	m['6'] = 5
	m['7'] = 6
	m['8'] = 7
	m['9'] = 8
	m['T'] = 9
	m['J'] = 10
	m['Q'] = 11
	m['K'] = 12
	m['A'] = 13
	return m
}
func get_map_B() (map[byte]int) {
	m := make(map[byte]int)
	m['2'] = 1
	m['3'] = 2
	m['4'] = 3
	m['5'] = 4
	m['6'] = 5
	m['7'] = 6
	m['8'] = 7
	m['9'] = 8
	m['T'] = 9
	m['J'] = 0
	m['Q'] = 11
	m['K'] = 12
	m['A'] = 13
	return m
}


func get_type_value(s string) int {
	frequency := make(map[rune]int)

	for _, char := range s {
		frequency[char]++
	}
	var numbers []int
	for _, count := range frequency {
		numbers = append(numbers, count)
	}
	for i, count := range numbers {
		switch count {
		case 5:
			return 7
		case 4:
			return 6
		case 3:
			for _, count2 := range numbers[i+1:] {
				if count2 == 2 {
					return 5
				}
			}
			return 4
		case 2:
			for _, count2 := range numbers[i+1:] {
				switch count2 {
				case 3:
					return 5
				case 2:
					return 3
				}
			}
			return 2
		}
	}
	return 1
}

type Pair struct {
	Value int
	Index int
}

func sort_by_value(values []int) []int {
	pairs := make([]Pair, len(values))
	for i, v := range values {
		pairs[i] = Pair{Value: v, Index: i}
	}

	// Sort
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value < pairs[j].Value
	})

	// Extract the sorted indices
	sortedIndices := make([]int, len(values))
	for i, pair := range pairs {
		sortedIndices[i] = pair.Index
	}
	return sortedIndices
}

func (p Day07) PartA(lines []string) any {
	val_map := get_map()
	values := make([]int, len(lines)-1)
	bids := make([]int, len(lines)-1)
	for n, line := range lines[:len(lines)-1] {
		content := strings.Split(line, " ")
		hand := content[0]
		bid, _ := strconv.Atoi(content[1])
		value := 0
		// value of type
		t := get_type_value(hand)
		// working in a 15 base system because the cards have 13 different values
		value += t*int(math.Pow(15, 5))
		for i:=0; i<len(hand); i++ {
			value += val_map[hand[i]]*int(math.Pow(15, float64(4-i)))
		}
		values[n] = value
		bids[n] = bid
	}

	sortedIndices := sort_by_value(values)

	total := 0
	for i:=0; i<len(sortedIndices); i++ {
		total += bids[sortedIndices[i]]*(i+1)
	}
	return total
}


func most_frequent_not_J(s string) (byte) {
	frequency := make(map[byte]int)

	for _, char := range s {
		frequency[byte(char)]++
	}
	max := 0
	var max_char byte
	max_char = byte('J')
	for char, count := range frequency {
		if char == 'J' {
			continue
		}
		if count > max {
			max = count
			max_char = char
		}
	}
	return max_char
}

func (p Day07) PartB(lines []string) any {
	val_map := get_map_B()
	values := make([]int, len(lines)-1)
	bids := make([]int, len(lines)-1)
	for n, line := range lines[:len(lines)-1] {
		content := strings.Split(line, " ")
		hand := content[0]
		// replace J with most frequent char for hand type calc
		hand_corr := strings.ReplaceAll(hand, "J", string(most_frequent_not_J(hand)))
		bid, _ := strconv.Atoi(content[1])
		value := 0
		t := get_type_value(hand_corr)
		value += t*int(math.Pow(15, 5))
		// use normal hand for this part
		for i:=0; i<len(hand); i++ {
			value += val_map[hand[i]]*int(math.Pow(15, float64(4-i)))
		}
		values[n] = value
		bids[n] = bid
	}

	sortedIndices := sort_by_value(values)

	total := 0
	for i:=0; i<len(sortedIndices); i++ {
		total += bids[sortedIndices[i]]*(i+1)
	}
	return total
}
