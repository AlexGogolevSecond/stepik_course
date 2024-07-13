/*
Определите является ли билет счастливым. Счастливым считается билет, в шестизначном номере которого сумма первых трёх цифр
совпадает с суммой трёх последних.

Формат входных данных

На вход подается номер билета - одно шестизначное  число.

Формат выходных данных
Выведите "YES", если билет счастливый, в противном случае - "NO".

Sample Input:

613244
Sample Output:

YES
*/
package main

import "fmt"

func main(){
	var number string
	fmt.Scan(&number)
	
	s1 := number[0] + number[1] + number[2]
	s2 := number[3] + number[4] + number[5]

	// //по идее 
	// fmt.Println(number[0])
	// fmt.Println(number[1])
	// fmt.Println(number[2])
	// fmt.Println(number[3])
	// fmt.Println(number[4])
	// fmt.Println(number[5])
	if s1 == s2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
