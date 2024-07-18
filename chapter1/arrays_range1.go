package main

import "fmt"

func main(){
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a) // [1 2 3 4 5]
	
	for idx, elem := range a {
		fmt.Printf("Элемент с индексом %d: %d\n", idx, elem)
		// Элемент с индексом 0: 1
		// Элемент с индексом 1: 2
		// Элемент с индексом 2: 3
		// Элемент с индексом 3: 4
		// Элемент с индексом 4: 5
	}

	fmt.Println("==================")
	// если не нужно или индекс или сам элемент (копия)
	b := [5]int{1, 2, 3, 4, 5}

	for idx := range b {
		fmt.Println(b[idx])
	}
	fmt.Println("==================")
	for idx, _ := range b {
		// В этом случае следует использовать приведенный выше вариант,
		// хотя технически эти варианты работают одинаково
		fmt.Println(b[idx])
	}
	fmt.Println("==================")
	for _, elem := range b {
		fmt.Println(elem)
	}

}
