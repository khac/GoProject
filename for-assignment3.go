// task: foo bar
// print all numbers from 1 to 100.
// along with each number,
// print "foo" if the number is divisible by 3
// print "bar" if the number is divisible by 5
// print "foobar" if the number is divisible by 3 and 5

/*something like
1
2
3 foo
4
5 bar
6 foo
7
8
9 foo
10 bar
11
12 foo
13
14
15 foobar
.
.
.
*/

package main

import "fmt"

func main() {
	// there is only for loop in go
	// do-while and while loop do not exist in golang
	var string = ""
	for i := 1; i < 100; i++ {
		fmt.Printf("%v", i)
		if i%3 == 0 {
			string += "foo"
		}
		if i%5 == 0 {
			string += "bar"
		}
		fmt.Printf(" %v\n", string)
		string = ""

	}
	fmt.Println()

}
