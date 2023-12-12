package year2023
//import "fmt"
type Day11 struct{}

func get_hits_in_row_and_col(lines []string) ([]int, []int) {
	dim := len(lines)-1

	hits_in_row := make([]int, dim)
	hits_in_col := make([]int, dim)

	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			if lines[i][j] == '#' {
				hits_in_row[i] += 1
				hits_in_col[j] += 1
			}
		}
	}
	return hits_in_row, hits_in_col
}

func get_distance_sum(hits_in_col []int, hits_in_row []int, add_empty int) int {
	total_distance := 0
	dim := len(hits_in_col)
	var extra_col int
	var extra_row int
	for i := 0; i < dim-1; i++ {
		extra_row = 0
		extra_col = 0
		for i2 := i; i2 < dim; i2++ {
			if hits_in_row[i2] == 0 {
				extra_row += 1
			} else {
				total_distance += hits_in_row[i]*hits_in_row[i2]*(i2-i+extra_row*add_empty)
			}
			if hits_in_col[i2] == 0 {
				extra_col += 1
			} else {
				total_distance += hits_in_col[i]*hits_in_col[i2]*(i2-i+extra_col*add_empty)
			}
		}
	}
	
	return total_distance
}

func (p Day11) PartA(lines []string) any {

	var hits_in_row []int
	var hits_in_col []int
	hits_in_row, hits_in_col = get_hits_in_row_and_col(lines)
	return get_distance_sum(hits_in_col, hits_in_row, 1)
}

func (p Day11) PartB(lines []string) any {

	var hits_in_row []int
	var hits_in_col []int
	hits_in_row, hits_in_col = get_hits_in_row_and_col(lines)
	return get_distance_sum(hits_in_col, hits_in_row, 1000000-1)
}
