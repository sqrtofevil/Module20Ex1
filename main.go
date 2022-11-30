package main

import (
	"fmt"
	"math/rand"
	"time"
)

var matrix [3][3]int

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = rand.Intn(10)
		}
	}
	fmt.Println("Входная матрица 3*3:")
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}

	fmt.Println("Определитель матрицы =", det(matrix))
}

func det(m [3][3]int) int {
	a := m[0][0] * m[1][1] * m[2][2]
	b := m[0][1] * m[1][2] * m[2][0]
	c := m[0][2] * m[1][0] * m[2][1]
	d := m[0][2] * m[1][1] * m[2][0]
	e := m[0][0] * m[1][2] * m[2][1]
	f := m[0][1] * m[1][0] * m[2][2]
	return (a + b + c - d - e - f)
}
