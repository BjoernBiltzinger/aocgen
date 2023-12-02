package year2023
//import "fmt"
import "strings"
import "strconv"
type Day02 struct{}

func (p Day02) PartA(lines []string) any {
	id_sum := 0
	for i, line := range lines[:len(lines)-1] {
		max_red := 0
		max_green := 0
		max_blue := 0
		line = strings.ReplaceAll(line, ":", "")
		line = strings.ReplaceAll(line, ";", "")
		line = strings.ReplaceAll(line, ",", "")
		game_info := strings.Split(line, " ")
		for i := 2; i < len(game_info)-1; i+=2 {
			number, _ := strconv.Atoi(game_info[i])
			color := game_info[i+1]
			switch color {
			case "blue":
				if number > max_blue {
					max_blue = number
				}
			case "green":
				if number > max_green {
					max_green = number
				}
			case "red":
				if number > max_red {
					max_red = number
				}
			}
		}
		if max_blue<=14 && max_green<=13 && max_red<=12 {
			id_sum += i+1
		}
	}

	return id_sum
}

func (p Day02) PartB(lines []string) any {
	power_sum := 0
	for _, line := range lines[:len(lines)-1] {
		max_red := 0
		max_green := 0
		max_blue := 0
		line = strings.ReplaceAll(line, ":", "")
		line = strings.ReplaceAll(line, ";", "")
		line = strings.ReplaceAll(line, ",", "")
		game_info := strings.Split(line, " ")
		for i := 2; i < len(game_info)-1; i+=2 {
			number, _ := strconv.Atoi(game_info[i])
			color := game_info[i+1]
			switch color {
			case "blue":
				if number > max_blue {
					max_blue = number
				}
			case "green":
				if number > max_green {
					max_green = number
				}
			case "red":
				if number > max_red {
					max_red = number
				}
			}
		}
		power_sum += max_red*max_green*max_blue
	}

	return power_sum
}
