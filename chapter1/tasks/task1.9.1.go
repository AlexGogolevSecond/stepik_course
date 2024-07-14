/*
На ввод подается целое число. Если число положительное - вывести сообщение "Число положительное",
если число отрицательное - "Число отрицательное". Если подается ноль - вывести сообщение "Ноль". Выводить сообщение без кавычек.

Sample Input:

5
Sample Output:

Число положительное
*/
package main

import "fmt"

func main(){
	var a int
	fmt.Scan(&a)

	// решение с if
	/*
	if a > 0 {
		fmt.Println("Число положительное")
	} else if a < 0 {
		fmt.Println("Число отрицательное")
	} else {
		fmt.Println("Ноль")
	}
	*/

	//решение с switch
	switch {
	case a > 0: fmt.Println("Число положительное")
	case a < 0: fmt.Println("Число отрицательное")
	case a == 0: fmt.Println("Ноль")
	}

}




