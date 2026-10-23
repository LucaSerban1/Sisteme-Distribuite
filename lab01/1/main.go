package main

import (
	"fmt"
)

func main(){
	var n int

	fmt.Scan(&n)

	if n < 2 || ( n > 2 && n % 2 == 0){
		fmt.Print("Nu este prim")
		return
	}

	for i := 3; i*i <= n; i += 2{
		if n % i == 0{
			fmt.Print("Nu este prim")
			return
		}
	}
	fmt.Print("Este prim")
}