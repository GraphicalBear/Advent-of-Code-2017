package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Checksum2 sums the result of subtracking the smallest from the largest number in each row
func Checksum2() {
	file, _ := os.Open("day2/spreadsheet.txt")
	scan := bufio.NewScanner(file)

	checksum := 0

	for scan.Scan() {
		numbers := strings.Fields(scan.Text())

		for index1, val1 := range numbers {
			num1, _ := strconv.Atoi(val1)

			for index2, val2 := range numbers {
				num2, _ := strconv.Atoi(val2)

				if index1 == index2 {
					continue
				}

				if num1%num2 == 0 {
					checksum += num1 / num2
				}

			}
		}
	}

	file.Close()
	fmt.Println(checksum)
}
