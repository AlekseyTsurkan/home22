package main

import "fmt"

func main() {
	arr := [12]int{1, 2, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var num int
	fmt.Println("Введите число:")
	fmt.Scan(&num)

	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			fmt.Printf("Индекс первого вхождения числа %d: %d\n", num, i)
			return
		}
	}

	fmt.Println("Число  не найдено в массиве", num)
}
