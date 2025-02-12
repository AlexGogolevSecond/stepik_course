/*
Последовательность состоит из натуральных чисел и завершается числом 0. Определите количество элементов этой последовательности,
которые равны ее наибольшему элементу.

Формат входных данных

Вводится непустая последовательность натуральных чисел, оканчивающаяся числом 0 (само число 0 в последовательность не входит,
а служит как признак ее окончания).

Формат выходных данных

Выведите ответ на задачу.

Sample Input:

1
3
3
1
0
Sample Output:

2
*/
package main

import "fmt"

func main(){
	var (
		a int
		max_elem int
		count_max_elem int
	)
	a = -1
	max_elem = 0
	count_max_elem = 0
	for a != 0 {
		fmt.Scan(&a)
		if a > max_elem {
			count_max_elem = 0
			max_elem = a
			count_max_elem += 1
		} else if a == max_elem {
			count_max_elem += 1
		}
	}
	fmt.Println(count_max_elem)
}