/*
По данному трехзначному числу определите, все ли его цифры различны.

Формат входных данных
На вход подается одно натуральное трехзначное число.

Формат выходных данных
Выведите "YES", если все цифры числа различны, в противном случае - "NO".

Sample Input 1:

237
Sample Output 1:

YES
Sample Input 2:

117
Sample Output 2:

NO
*/

package main

import "fmt"

func main(){
	var s string
	fmt.Scan(&s)

	var a, b, c bool

	a = s[0] == s[1]
	b = s[0] == s[2]
	c = s[1] == s[2]

	
	if a || b || c {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}



