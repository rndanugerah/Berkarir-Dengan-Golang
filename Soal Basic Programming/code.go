package main

import "fmt"

func main() {
	bilanganPrima()
	kelipatan7()
	luasTrapesium()
}

func bilanganPrima() {
	var prima int
	prima = 9

	if prima%2 == 0 {
		fmt.Println(prima, "bukan bilangan prima")
	} else {
		fmt.Println(prima, "adalah bilangan prima")
	}
}

func kelipatan7() {
	var number int
	number = 15

	if number%7 == 0 {
		fmt.Println(number, "adalah kelipatan 7")
	} else {
		fmt.Println(number, "bukan kelipatan 7")
	}
}

func luasTrapesium() {
	var tinggi, panjangA, panjangB int
	tinggi = 25
	panjangA = 10
	panjangB = 20

	luas := (panjangA + panjangB) * tinggi / 2
	fmt.Println("Luas trapesium:", luas)
}
