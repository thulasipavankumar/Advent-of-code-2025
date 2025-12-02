package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/thulasipavankumar/Advent-of-code-2025/util"
)

// 6461988558
const FILE_PATH = "input.txt"

func part2(lines []string) (count int) {

	return
}
func part1(input [][2]int) (count int) {

	for _, pair := range input {

		//fmt.Printf("Index %d: left=%d, right=%d\n", index, pair[0], pair[1])

		for i := pair[0]; i <= pair[1]; i++ {
			if isNumPalindrome(i) {
				count += i
			}
		}
		println(count)
	}

	return
}
func isNumPalindrome(num int) bool {
	if num < 10 || len(strconv.Itoa(num))%2 != 0 {
		return false
	} else {
		return dividePartsAndCheck(num)
	}
}
func dividePartsAndCheck(num int) bool {
	s := strconv.Itoa(num)
	mid := len(s) / 2
	return s[:mid] == s[mid:]

}
func main() {
	input := util.GetFileContent(FILE_PATH)
	parts := strings.Split(input[0], ",")
	println(len(parts))
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
