package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	//var number []int
	var n, m int
	fmt.Println("enter the limit")
	fmt.Scanln(&n)
	//var number []int
	fmt.Println("Enter Number To check even or odd")
	var number = make([]int, 0, n)
	for i := 0; i < n; i++ {

		fmt.Scan(&m)
		number = append(number, m)
		sort.Ints(number)
		//fmt.Println(number)

	}
	fmt.Println(number)
	c := make(chan int)

	for i, m := range number {

		go odd(m, c, i, n)

	}
	for i := 0; i < n; i++ {
		fmt.Println(<-c)
	}
}

func odd(m int, c chan int, i int, n int) {

	if m%2 == 0 {
		for r := 0; r <= n; r++ {

			time.Sleep(time.Second)

			//c <- m
		}
		c <- m

	}
	if m%2 != 0 {
		for r := 0; r < n; r++ {
			time.Sleep(time.Second)
		}

		time.Sleep(4 * time.Second)
		for r := 0; r <= i; r++ {
			time.Sleep(time.Second)
			//c<-m
		}

		c <- m
		//time.Sleep(5 * time.Second)
	}
	//}
	return

}
