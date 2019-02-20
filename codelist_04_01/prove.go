package main

import "fmt"

func prove(n int) {
	k := 0
	fmt.Printf("现在开始证明P(%d)成立。\n", n)
	fmt.Printf("根据步骤1得出P(%d)成立。\n", k)
	for k < n {
		fmt.Printf("根据步骤2可以说\"若P(%d)成立，则P(%d)也成立\"。\n", k, k+1)
		fmt.Printf("所以可以说P(%d)成立。\n", k+1)
		k++
	}
	fmt.Println("证明结束。")
}

func main() {
	prove(2)
}
