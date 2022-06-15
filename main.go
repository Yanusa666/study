package main

import "fmt"

func main() {
	var lenMs1 int
	fmt.Println("Введите размер первого массива")
	_, err := fmt.Scanln(&lenMs1)
	if err != nil {
		fmt.Println(fmt.Sprintf("Не получилось считать число: %s", err))
	}

	var lenMs2 int
	fmt.Println("Введите размер второго массива")
	_, err = fmt.Scanln(&lenMs2)
	if err != nil {
		fmt.Println(fmt.Sprintf("Не получилось считать число: %s", err))
	}

	fmt.Println("Введите элементы первого массива")
	map1 := make(map[int]string)
	var element string
	for i := 1; i <= lenMs1; i++ {
				_, err := fmt.Scanln(&element)
		if err != nil {
			fmt.Println(fmt.Sprintf("Не получилось считать элемент: %s", err))
		}
		map1[i]=element
	}

	fmt.Println("Введите элементы первого массива")
	map2 := make(map[int]string)
	var elements string
	for i := 1; i <= lenMs2; i++ {
		_, err := fmt.Scanln(&elements)
		if err != nil {
			fmt.Println(fmt.Sprintf("Не получилось считать элемент: %s", err))
		}
		map2[i]=elements
	}

	var	dup [] string
for _, val1 :=range map1 {
	for _, val2:= range map2 {
		if val1==val2 {
			dup = append (dup,val1)
		}
	}
}


	fmt.Println("Одинаковые элементы:", dup)
}