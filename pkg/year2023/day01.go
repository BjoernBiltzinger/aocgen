package year2023
import "unicode"
import "strings"

type Day01 struct{}
func isInt(s string) bool {
    for _, c := range s {
        if !unicode.IsDigit(c) {
            return false
        }
    }
    return true
}
func (p Day01) PartA(lines []string) any {
	value := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		zero_number := int('0')
		var low, high int
		for j := 0; j < len(line); j++ {
			if isInt(string(line[j])) {
				low = int(line[j])-zero_number
				break
			}
		}
		for j := len(line) - 1; j >= 0; j-- {
			if isInt(string(line[j])) {
				high = int(line[j])-zero_number
				break
			}
		}
		value += low*10 + high
	}
	return value
}

func (p Day01) PartB(lines []string) any {
	value := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		zero_number := int('0')
		var low, high int
		for j := 0; j < len(line); j++ {
			if isInt(string(line[j])) {
				low = int(line[j]) - zero_number
				break
			}
			if strings.HasPrefix(line[j:], "one") {
				low = int('1') - zero_number
				break
			}
			if strings.HasPrefix(line[j:], "two") {
				low = int('2') - zero_number
				break
			}
			if strings.HasPrefix(line[j:], "three") {
				low = int('3') - zero_number
				break
			}
			if strings.HasPrefix(line[j:], "four") {
				low = int('4') - zero_number
				break
			}
			if strings.HasPrefix(line[j:], "five") {
				low = int('5') - zero_number
				break
			}
			if strings.HasPrefix(line[j:], "six") {
				low = int('6') - zero_number
				break
			}
			if strings.HasPrefix(line[j:], "seven") {
				low = int('7') - zero_number
				break
			}
			if strings.HasPrefix(line[j:], "eight") {
				low = int('8') - zero_number
				break
			}
			if strings.HasPrefix(line[j:], "nine") {
				low = int('9') - zero_number
				break
			}
			if strings.HasPrefix(line[j:], "zero") {
				low = int('0') - zero_number
				break
			}
		}
		for j := len(line) - 1; j >= 0; j-- {
			if isInt(string(line[j])) {
				high = int(line[j]) - zero_number
				break
			}
			if strings.HasSuffix(line[:j+1], "one") {
				high = int('1') - zero_number
				break
			}
			if strings.HasSuffix(line[:j+1], "two") {
				high = int('2') - zero_number
				break
			}
			if strings.HasSuffix(line[:j+1], "three") {
				high = int('3') - zero_number
				break
			}
			if strings.HasSuffix(line[:j+1], "four") {
				high = int('4') - zero_number
				break
			}
			if strings.HasSuffix(line[:j+1], "five") {
				high = int('5') - zero_number
				break
			}
			if strings.HasSuffix(line[:j+1], "six") {
				high = int('6') - zero_number
				break
			}
			if strings.HasSuffix(line[:j+1], "seven") {
				high = int('7') - zero_number
				break
			}
			if strings.HasSuffix(line[:j+1], "eight") {
				high = int('8') - zero_number
				break
			}
			if strings.HasSuffix(line[:j+1], "nine") {
				high = int('9') - zero_number
				break
			}
			if strings.HasSuffix(line[:j+1], "zero") {
				high = int('0') - zero_number
				break
			}
		}
		value += 10*low + high
	}
	return value
}
