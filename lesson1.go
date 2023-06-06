package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 10

func main() {
	rand.Seed(time.Now().UnixNano())

	arr := make([]int, size)

	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(100)
	}

	fmt.Println("Массив:", arr)

	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)

	found := false
	count := 0
	for _, n := range arr {
		if n == num {
			found = true
		}
		if found {
			count++
		}
	}
	if found {
		fmt.Printf("Число %d найдено в массиве. Количество чисел после него: %d\n", num, count-1)
	} else {
		fmt.Println("Число не найдено в массиве попробуйте еще раз :с.")
	}
}
