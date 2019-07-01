package main

import "fmt"

func main() {
	words := []string{"abcd", "acacacac", "qweqwe", "a", "b", "c", "d", "bcbc"}
	fmt.Printf("%v", words)
	mergeSort(words, len(words))
	fmt.Printf("%v", words)
}

func mergeSort(data []string, lenD int) []string {
	if lenD > 1 {
		middle := lenD / 2
		rem := lenD - middle
		L := make([]string, middle)
		R := make([]string, rem)
		for i := 0; i < lenD; i++ {
			if i < middle {
				L[i] = string(data[i])
			} else {
				R[i-middle] = string(data[i])
			}
		}
		mergeSort(L, middle)
		mergeSort(R, rem)
		merge(data, lenD, L, middle, R, rem)
	}
	return data
}

func merge(merged []string, lenD int, L []string, lenL int, R []string, lenR int) []string {
	i := 0
	j := 0
	for i < lenL || j < lenR {
		if i < lenL && j < lenR {
			if L[i] <= R[i] {
				merged[i+j] = L[i]
				i++
			} else {
				merged[i+j] = R[j]
			}
		} else if i < lenL {
			merged[i+j] = L[i]
			i++
		} else if j < lenR {
			merged[i+j] = R[j]
			j++
		}
	}
	return merged
}
