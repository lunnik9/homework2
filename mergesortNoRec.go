package main

import "fmt"

func main() {

	massiv := []string{"abcd", "acacacac", "qweqwe", "a", "b", "c", "d", "bcbc"}
	baseLength := len(massiv)
	temp := make([]string, len(massiv)*2)
	for i := 1; i < baseLength*2; i++ {

		if i%2 != 0 {
			temp[i-1] = massiv[i/2]
		} else {
			temp[i-1] = " "
		}

	}
	temp = append(temp[:len(temp)-1])
	var back []string
	var forward []string
	var temptemp []string
	var temp1 []string
	var temp2 []string
	ii, jj := 0, 0
	var b, f int
	for len(temp) > baseLength {

		for i := 0; i < len(temp); i++ {
			if temp[i] == " " {
				for b = 0; temp[i-b-1] != " " && (i-b-1) >= 0 && i > b; b++ {
					//fmt.Fprintln(b)
					back[b] = temp[i-b+1]
				}

				fmt.Print(back)

				for z := len(back)/2 - 1; z >= 0; z-- {
					opp := len(back) - 1 - z
					back[z], back[opp] = back[opp], back[z]
				}

				for f = 0; temp[i+f] != " " && &temp[i+f] != nil; f++ {
					forward[f] = temp[i+f]
				}

				fmt.Print(forward)

				ii, jj = 0, 0
				for k := 0; k < b+f; k++ {
					if ii > b-1 && jj <= f-1 {
						temptemp[k] = forward[jj]
						jj++
					} else if jj > f-1 && ii <= b-1 {
						temptemp[k] = back[i]
						i++
					} else if back[ii] < forward[jj] {
						temptemp[k] = back[ii]
						ii++
					} else {
						temptemp[k] = forward[jj]
						jj++
					}
				}

				temp1 = temp[0 : i-b]
				temp2 = temp[i+f : len(temp)]
				temp = append(temp1, temptemp...)
				temp = append(temp, temp2...)

			}

		}
	}

	fmt.Printf("%v \n", massiv)
	fmt.Printf("%v ", temp)
}
