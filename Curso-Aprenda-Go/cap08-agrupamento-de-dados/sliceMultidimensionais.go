/*
Slices multi-dimensionais são slices que contem slices.
- São como planilhas.
- [][]type
*/

package main

import (
	"fmt"
)

func main() {
	ss := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	sss := [][][]int{
		{
			{1, 2, 3, 4},
		},
		{
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		},
	}

	fmt.Println(ss[1][3]) // 8

	fmt.Println(sss[1]) // [[5 6 7 8] [9 10 11 12]]

	fmt.Println(sss[1][1]) // [9 10 11 12]

	fmt.Println(sss[1][1][3]) // 12
}