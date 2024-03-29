package sugoku

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// Difficulty will be used in a const to specify different types of difficulties
type Difficulty int

const (
	limitNumber  int = 9
	gridQuantity int = 9
)

// This block will have the different type of difficulties availables
const (
	Easy   Difficulty = 1
	Medium Difficulty = 2
	Hard   Difficulty = 3
)

// Sugoku represents a solved grid
type Sugoku [limitNumber][limitNumber]int

var gridPos = map[int]map[string]int{
	1: map[string]int{
		"cs": 0, // Col Start
		"ce": 2, // Col End
		"rs": 0, // Row Start
		"re": 2, // Row End
	},
	2: map[string]int{
		"cs": 3, // Col Start
		"ce": 5, // Col End
		"rs": 0, // Row Start
		"re": 2, // Row End
	},
	3: map[string]int{
		"cs": 6, // Col Start
		"ce": 8, // Col End
		"rs": 0, // Row Start
		"re": 2, // Row End
	},
	4: map[string]int{
		"cs": 0, // Col Start
		"ce": 2, // Col End
		"rs": 3, // Row Start
		"re": 5, // Row End
	},
	5: map[string]int{
		"cs": 3, // Col Start
		"ce": 5, // Col End
		"rs": 3, // Row Start
		"re": 5, // Row End
	},
	6: map[string]int{
		"cs": 6, // Col Start
		"ce": 8, // Col End
		"rs": 3, // Row Start
		"re": 5, // Row End
	},
	7: map[string]int{
		"cs": 0, // Col Start
		"ce": 2, // Col End
		"rs": 6, // Row Start
		"re": 8, // Row End
	},
	8: map[string]int{
		"cs": 3, // Col Start
		"ce": 5, // Col End
		"rs": 6, // Row Start
		"re": 8, // Row End
	},
	9: map[string]int{
		"cs": 6, // Col Start
		"ce": 8, // Col End
		"rs": 6, // Row Start
		"re": 8, // Row End
	},
}

// GenerateSugoku returns a generated grid
func GenerateSugoku() Sugoku {
	start := time.Now()

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	var sudoku Sugoku
	regenerate := false
	retries := 0
	for g := 1; g <= gridQuantity; g++ {
		if regenerate == true {
			retries++
			if g = g - retries; (g - retries) < 1 {
				g = 1
			}
			cleanGrid(&sudoku, g)
			regenerate = false
		} else {
			retries = 0
		}
		conf := gridPos[g]
		var valsGenerated []int
		regen := false
		for r := conf["rs"]; r <= conf["re"]; r++ {
			for c := conf["cs"]; c <= conf["ce"]; c++ {
				value := 0
				possibleValues := getPossibleValues(r, c, valsGenerated, &sudoku)
				if len(possibleValues) == 0 {
					regen = true
					break
				}

				pos := r1.Intn(len(possibleValues))
				value = possibleValues[pos]

				if regen == true {
					break
				}
				valsGenerated = append(valsGenerated, value)
				sudoku[r][c] = value
			}

			if regen == true {
				g = g - 1
				regenerate = true
				break
			}
		}
	}

	elapsed := time.Since(start)
	log.Printf("Sugoku took %s", elapsed)
	return sudoku
}

func getPossibleValues(r int, c int, vG []int, s *Sugoku) []int {
	var vA []int // Values Available
	for v := 1; v <= limitNumber; v++ {
		if contains(vG, v) == true {
			continue
		}

		if validateCondition(s, r, c, v) == true {
			continue
		}

		vA = append(vA, v)
	}

	return vA
}

func cleanGrid(s *Sugoku, sp int) {

	for p := sp; p <= gridQuantity; p++ {
		conf := gridPos[p]
		for r := conf["rs"]; r <= conf["re"]; r++ {
			for c := conf["cs"]; c <= conf["ce"]; c++ {
				s[r][c] = 0
			}
		}
	}
}

func validateCondition(s *Sugoku, row int, col int, v int) bool {
	cond := false

	rows := s[row]
	cond = (contains(rows[:], v) || cond)

	if cond == true {
		return cond
	}

	var cols []int
	for r := 0; r < limitNumber; r++ {
		cols = append(cols, s[r][col])
	}
	cond = (contains(cols, v) || cond)
	return cond
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Solve will return the solved Sugoku
func (s Sugoku) Solve() {
	fmt.Println()
	for r := 0; r < limitNumber; r++ {
		for c := 0; c < limitNumber; c++ {
			if c%3 == 0 {
				fmt.Print(" ")
			}
			fmt.Print(s[r][c], " ")
		}
		fmt.Println()
		if (r+1)%3 == 0 {
			fmt.Println()
		}
	}
}

// Prepare will return a Sugoku ready to be played
func (s Sugoku) Prepare(d Difficulty) {
	if d == 0 {
		d = Medium
	}
	fmt.Println(s)
}
