package year2023
import "strings"
import "strconv"
//import "fmt"
type Day05 struct{}

type slice struct {
	low int
	high int
}

func slices_overlap(slice1 slice, slice2 slice) (bool) {
	if slice1.low > slice2.low {
		slice1, slice2 = slice2, slice1
	}
	if slice1.high > slice2.low {
		return true
	}
	return false
}

func number_in_slice(slice slice, number int) (bool) {
	if number >= slice.low && number < slice.high {
		return true
	}
	return false
}

func map_slice(input slice, map_slice slice, diff int) ([]slice, []slice) {
	if !slices_overlap(input, map_slice) {
		return []slice{}, []slice{input}
	}
	if map_slice.low <= input.low && map_slice.high >= input.high {
		// everything is mapped
		return []slice{slice{input.low+diff, input.high+diff}}, []slice{}
	}
	if map_slice.low <= input.low {
		// lower part is mapped 
		return []slice{slice{input.low+diff, map_slice.high+diff}}, []slice{slice{map_slice.high, input.high}}
	}
	if map_slice.high >= input.high {
		// upper part is mapped 
		return []slice{slice{map_slice.low+diff, input.high+diff}}, []slice{slice{input.low, map_slice.low}}
	}
	// middle part is mapped
	return []slice{slice{map_slice.low+diff, map_slice.high+diff}}, []slice{slice{input.low, map_slice.low}, slice{map_slice.high, input.high}}
}

func get_slice_and_diff(line string) (slice, int) {
	numbers := strings.Split(line, " ")
	start_to, _ := strconv.Atoi(numbers[0])
	start_from, _ := strconv.Atoi(numbers[1])
	length, _ := strconv.Atoi(numbers[2])
	return slice{low:start_from, high:start_from+length}, start_to-start_from
}

func get_min_of_slices(slices []slice) (int) {
	min := -1
	for _, s := range slices {
		if min == -1 || s.low < min {
			min = s.low
		}
	}
	return min
}

func get_min_of_numbers(numbers []int) (int) {
	min := -1
	for _, num := range numbers {
		if min == -1 || num < min {
			min = num
		}
	}
	return min
}

func (p Day05) PartA(lines []string) any {
	seed_ids_strings := strings.Split(lines[0], " ")[1:]
	var numbers []int
	var s slice
	var diff int
	var diffs []int
	var map_slices []slice
	for _, seed_id_string := range seed_ids_strings {
		seed_id, _ := strconv.Atoi(seed_id_string)
		numbers = append(numbers, seed_id)
	}
	current_numbers := numbers
	for line_num := 3; line_num < len(lines); line_num++ {
		line := lines[line_num]
		if len(line) == 0 {
			// Do things
			mapped_numbers := []int{}
			for i:=0; i<len(diffs); i++ {
				unmapped_numbers := []int{}
				for _, num := range current_numbers {
					if number_in_slice(map_slices[i], num) {
						mapped_numbers = append(mapped_numbers, num+diffs[i])
					} else {
						unmapped_numbers = append(unmapped_numbers, num)
					}
				}
				current_numbers = unmapped_numbers
			}
			current_numbers = append(current_numbers, mapped_numbers...)
			line_num++
			map_slices = []slice{}
			diffs = []int{}
			continue
		}
		s, diff = get_slice_and_diff(line)
		map_slices = append(map_slices, s)
		diffs = append(diffs, diff)
	}
	return get_min_of_numbers(current_numbers)
}

func (p Day05) PartB(lines []string) any {
	seed_ids_strings := strings.Split(lines[0], " ")[1:]
	var slices []slice
	var s slice
	var diff int
	var diffs []int
	var map_slices []slice
	for k:=0; k<len(seed_ids_strings); k+=2 {
		start, _ := strconv.Atoi(seed_ids_strings[k])
		length, _ := strconv.Atoi(seed_ids_strings[k+1])
		slices = append(slices, slice{start, start+length})
	}
	current_slices := slices
	for line_num := 3; line_num < len(lines); line_num++ {
		line := lines[line_num]
		if len(line) == 0 {
			// Do things
			mapped_slices := []slice{}
			//current_slices := slices
			for i:=0; i<len(diffs); i++ {
				unmapped_slices := []slice{}
				for _, s := range current_slices {
					mapped, unmapped := map_slice(s, map_slices[i], diffs[i])
					mapped_slices = append(mapped_slices, mapped...)
					unmapped_slices = append(unmapped_slices, unmapped...)
				}
				current_slices = unmapped_slices
			}
			current_slices = append(current_slices, mapped_slices...)
			line_num++
			map_slices = []slice{}
			diffs = []int{}
			continue
		}
		s, diff = get_slice_and_diff(line)
		map_slices = append(map_slices, s)
		diffs = append(diffs, diff)
	}

	//fmt.Println(starts, lengths)
	return get_min_of_slices(current_slices)
}


// FIRST APPROACH FOR PART A => 118 mus

// type seed struct {
// 	id  int
//     soil int
// 	fertilizer int
// 	water int
// 	light int 
// 	temperature int 
// 	humidity int
// 	location int
// }

// func get_start_stop(line string) (int, int, int) {
// 	numbers := strings.Split(line, " ")
// 	start_to, _ := strconv.Atoi(numbers[0])
// 	start_from, _ := strconv.Atoi(numbers[1])
// 	length, _ := strconv.Atoi(numbers[2])
// 	return start_from, start_to, length
// }

// func apply_map(seeds []seed, line string, t int) ([]seed){
// 	var start_from, start_to, length int
// 	start_from, start_to, length = get_start_stop(line)

// 	diff := start_to-start_from
// 	var new_seeds []seed
// 	for _, seed := range seeds {
// 		switch t {
// 		case 0:
// 			if seed.id>=start_from && seed.id<start_from+length {
// 				seed.soil = seed.id+diff
// 			}
// 		case 1:
// 			if seed.soil>=start_from && seed.soil<start_from+length {
// 				seed.fertilizer = seed.soil+diff
// 			}
// 		case 2:
// 			if seed.fertilizer>=start_from && seed.fertilizer<start_from+length {
// 				seed.water = seed.fertilizer+diff
// 			}
// 		case 3:
// 			if seed.water>=start_from && seed.water<start_from+length {
// 				seed.light = seed.water+diff
// 			}
// 		case 4:
// 			if seed.light>=start_from && seed.light<start_from+length{
// 				seed.temperature = seed.light+diff
// 			}
// 		case 5:
// 			if seed.temperature>=start_from && seed.temperature<start_from+length {
// 				seed.humidity = seed.temperature+diff
// 			}
// 		case 6:
// 			if seed.humidity>=start_from && seed.humidity<start_from+length {
// 				seed.location = seed.humidity+diff
// 			}
// 		}
// 		new_seeds = append(new_seeds, seed)
// 	}
// 	return new_seeds
// }

// func add_default(seeds []seed, t int) ([]seed){
// 	var new_seeds []seed
// 	for _, seed := range seeds {
// 		switch t {
// 		case 0:
// 			seed.soil = seed.id
// 		case 1:
// 			seed.fertilizer = seed.soil
// 		case 2:
// 			seed.water = seed.fertilizer
// 		case 3:
// 			seed.light = seed.water
// 		case 4:
// 			seed.temperature = seed.id
// 		case 5:
// 			seed.humidity = seed.temperature
// 		case 6:
// 			seed.location = seed.humidity
// 		}
// 		new_seeds = append(new_seeds, seed)
// 	}
// 	return new_seeds
// }


// func (p Day05) PartA(lines []string) any {
// 	seed_ids_strings := strings.Split(lines[0], " ")[1:]
// 	var seed_ids []int
// 	for _, seed_id_string := range seed_ids_strings {
// 		seed_id, _ := strconv.Atoi(seed_id_string)
// 		seed_ids = append(seed_ids, seed_id)
// 	}
// 	var seeds []seed
// 	for _, seed_id := range seed_ids {
// 		seeds = append(seeds, seed{seed_id, 0, 0, 0, 0, 0, 0, 0})
// 	}
// 	t := 0
// 	seeds = add_default(seeds, 0)
// 	for _, line := range lines[3:] {
// 		if len(line) == 0 {
// 			continue
// 		}
// 		if strings.Contains(line, "map"){
// 			t++
// 			seeds = add_default(seeds, t)

// 			continue
// 		}
// 		seeds = apply_map(seeds, line, t)
// 	}
// 	smallest_loc := -1
// 	for _, seed := range seeds {
// 		if smallest_loc == -1 || seed.location < smallest_loc {
// 			smallest_loc = seed.location
// 		}
// 	}
// 	return smallest_loc
// }