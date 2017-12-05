package day2

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Checksum1 sums the result of subtracting the smallest from the largest number in each row
func Checksum1() {
	file, _ := os.Open("day2/spreadsheet.txt")
	scan := bufio.NewScanner(file)

	checksum := 0

	for scan.Scan() {
		numbers := strings.Fields(scan.Text())

		small := math.MaxInt32
		large := -math.MaxInt32

		for _, val := range numbers {
			num, _ := strconv.Atoi(val)

			if num < small {
				small = num
			}

			if num > large {
				large = num
			}
		}

		checksum += large - small
	}

	file.Close()
	fmt.Println(checksum)
}
