package main

import (
	"fmt"
	"strconv"
)

// https://play.golang.org/p/Qg_uv_inCek
// contains checks if a string is present in a slice
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func countContains(s []string, str string) int {
	var c = 0
	for _, v := range s {
		if v == str {
			c++
		}
	}

	return c
}

func main() {

	goodCouple := []string{"22", "23", "32", "24", "42", "26", "62", "29", "92", "36", "63", "24", "42", "36", "63", "66", "28", "82", "15", "51", "35", "53", "45", "54", "89", "98", "99", "35", "53", "49", "94", "15", "51", "55", "49", "94", "95", "59", "99"}
	fmt.Println(goodCouple)

	for i := 1000; i <= 9999; i++ {
		sTabian := strconv.Itoa(i)
		sec1 := sTabian[0:2]
		sec2 := sTabian[1:3]
		sec3 := sTabian[2:4]

		//fmt.Printf("%v -> %v,%v,%v\n", sTabian, sec1, sec2, sec3)
		//fmt.Println(contains(goodCouple, sec1))
		//fmt.Println(contains(goodCouple, sec2))
		//fmt.Println(contains(goodCouple, sec3))

		//fmt.Printf("\r%d", i)

		if contains(goodCouple, sec1) {
			if contains(goodCouple, sec2) {
				if contains(goodCouple, sec3) {

					//fmt.Printf("%v\n", sTabian)

					g := []string{}

					//create slice g
					for j := 0; j <= 3; j++ {
						//fmt.Println(sTabian[j:(j + 1)])
						n := sTabian[j:(j + 1)]
						g = append(g, n)
					}

					//check existing
					check := true
					for k := 0; k <= 3; k++ {
						cc := countContains(g, sTabian[k:(k+1)])
						//fmt.Printf("%v->%d\n", sTabian[k:(k+1)], cc)
						if cc < 2 {
							check = true
						} else {
							check = false
							break
						}
					}

					if check {
						fmt.Println(sTabian)
					}

				}
			}
		}

	}

}
