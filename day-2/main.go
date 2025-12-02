package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/thulasipavankumar/Advent-of-code-2025/util"
)

// 6461988558
const FILE_PATH = "input.txt"

func part2(input [][2]int) (count int) {

	return
}
func part1(input [][2]int) (count int) {

	for _, pair := range input {

		//fmt.Printf("Index %d: left=%d, right=%d\n", index, pair[0], pair[1])

		for i := pair[0]; i <= pair[1]; i++ {
			if isNumMirrorPalindrome(i) {
				count += i
			} else {
				//println("not equal parts", i)
			}
		}
	}

	return
}
func isNumMirrorPalindrome(num int) bool {
	if num < 10 {
		return false
	} else {
		return dividePartsAndCheck2(num)
	}
}
func dividePartsAndCheck(num int) bool {
	s := strconv.Itoa(num)
	mid := len(s) / 2
	parts := []string{s[:mid], s[mid:]}
	return checkIfSliceHasEqualValues(parts)

}
func dividePartsAndCheck2(num int) bool {
	var equalParts bool = false
	s := splitNumberAll(num)
	for _, group := range s {

		equalParts = checkIfSliceHasEqualValues(group)

		if equalParts {
			break
		}

	}
	return equalParts
}
func checkIfSliceHasEqualValues(input []string) bool {
	equalParts := false
	if len(input) <= 1 {
		return false
	}
	first := input[0]
	for _, v := range input[1:] {
		if v != first {
			equalParts = false
			break
		} else {
			equalParts = true
		}
	}
	return equalParts
}
func splitNumberAll(n int) [][]string {
	s := strconv.Itoa(n)
	L := len(s)
	half := L / 2

	results := [][]string{}

	for size := 1; size <= half; size++ {
		if L%size != 0 {
			continue
		}
		parts := []string{}
		for i := 0; i < L; i += size {
			parts = append(parts, s[i:i+size])
		}
		results = append(results, parts)
	}

	return results
}
func main() {
	input := util.GetFileContent(FILE_PATH)
	parts := strings.Split(input[0], ",")
	sequenceInput := make([][2]int, 0)
	for _, p := range parts {
		lr := strings.SplitN(p, "-", 2)
		left, err1 := strconv.Atoi(lr[0])
		right, err2 := strconv.Atoi(lr[1])
		if err1 != nil || err2 != nil {
			log.Fatal("wrong input")
		}

		sequenceInput = append(sequenceInput, [2]int{left, right})
	}
	println(part1(sequenceInput))
}
