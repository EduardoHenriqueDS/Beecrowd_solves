package main

import (
	"fmt"
)

var n int

func main() {

	fmt.Scan(&n)
	fmt.Println(n)

	if n >= 100 {
		onehundred := (n / 100)
		fmt.Printf("%v nota(s) de R$ 100,00\n", onehundred)
		n -= onehundred * 100
	} else {
		onehundred := 0
		fmt.Printf("%v nota(s) de R$ 100,00\n", onehundred)
	}
	if n >= 50 && n < 100 {
		fifty := (n / 50)
		fmt.Printf("%v nota(s) de R$ 50,00\n", fifty)
		n -= fifty * 50
	} else {
		fifty := 0
		fmt.Printf("%v nota(s) de R$ 50,00\n", fifty)
	}
	if n >= 20 && n < 50 {
		twenty := (n / 20)
		fmt.Printf("%v nota(s) de R$ 20,00\n", twenty)
		n -= twenty * 20
	} else {
		twenty := 0
		fmt.Printf("%d nota(s) de R$ 20,00\n", twenty)
	}
	if n >= 10 && n < 20 {
		ten := (n / 10)
		fmt.Printf("%v nota(s) de R$ 10,00\n", ten)
		n -= ten * 10
	} else {
		ten := 0
		fmt.Printf("%v nota(s) de R$ 10,00\n", ten)
	}
	if n >= 5 && n < 10 {
		five := (n / 5)
		fmt.Printf("%v nota(s) de R$ 5,00\n", five)
		n -= five * 5
	} else {
		five := 0
		fmt.Printf("%v nota(s) de R$ 5,00\n", five)
	}
	if n >= 2 && n < 5 {
		two := (n / 2)
		fmt.Printf("%v nota(s) de R$ 2,00\n", two)
		n -= two * 2
	} else {
		two := 0
		fmt.Printf("%v nota(s) de R$ 2,00\n", two)
	}
	if n >= 1 && n < 2 {
		one := (n / 1)
		fmt.Printf("%v nota(s) de R$ 1,00\n", one)
		n -= one * 1
	} else {
		one := 0
		fmt.Printf("%v nota(s) de R$ 1,00\n", one)
	}

}
