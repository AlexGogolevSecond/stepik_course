/*
Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.

Входные данные

Программа получает на вход два числа. Гарантируется, что цифры в числах не повторяются. Числа в пределах от 0 до 10000.

Выходные данные

Программа должна вывести цифры, которые имеются в обоих числах, через пробел. Цифры выводятся в порядке их нахождения в первом
числе.

Внимание! Если вам сложно решить эту задачу, пропустите и изучайте материал дальше, затем вернетесь к этой задаче позже.

Sample Input:

564 8954
Sample Output:

5 4
*/
package main

import "fmt"

func main(){
	var (
		a, b string
	)

	fmt.Scan(&a, &b)

	for i:= 0; i < len(a); i++ {
		for j := 0; j < len(b); j ++ {
			if a[i] == b[j] {
				fmt.Print(string(a[i]) + " ")
			}
		}
	}
}
