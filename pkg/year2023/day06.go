package year2023
import "strings"
//import "fmt"
import "strconv"
import "math"
type Day06 struct{}

func find_roots_quadratic(a int, b int, c int) (float64, float64) {
	a_float := float64(a)
	b_float := float64(b)
	c_float := float64(c)
	sqrt := math.Sqrt(math.Pow(b_float, 2)-4*a_float*c_float)
	return (-b_float-sqrt)/(2*a_float), (-b_float+sqrt)/(2*a_float)
}

func (p Day06) PartA(lines []string) any {
	time_strings := strings.Split(lines[0], " ")[1:]
	distance_strings := strings.Split(lines[1], " ")[1:]
	var times []int
	var distances []int


	total := 1.0
	for _, time_string := range time_strings {
		time, err := strconv.Atoi(time_string)
		if err != nil {
			continue
		}
		times = append(times, time)
	}
	for _, distance_string := range distance_strings {
		distance, err := strconv.Atoi(distance_string)
		if err != nil {
			continue
		}
		distances = append(distances, distance)
	}
	for i := 0; i < len(times); i++ {
		// from below
		j_low_float, j_high_float := find_roots_quadratic(1, -times[i], distances[i])
		total *= math.Floor(j_high_float)-math.Ceil(j_low_float)+1
	}
	return total
}

func (p Day06) PartB(lines []string) any {
	time_string := strings.Split(strings.ReplaceAll(lines[0], " ", ""), ":")[1]
	distance_string := strings.Split(strings.ReplaceAll(lines[1], " ", ""), ":")[1]
	time, _ := strconv.Atoi(time_string)
	distance, _ := strconv.Atoi(distance_string)
	j_low_float, j_high_float := find_roots_quadratic(1, -time, distance)
	return math.Floor(j_high_float)-math.Ceil(j_low_float)+1
}
