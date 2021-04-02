package main

import (
	"fmt"
)

func main() {
//	var f1 float32
//	var f2 float64
	var maxFloat32 = 16777216
	fmt.Println(maxFloat32 == maxFloat32+10) //retun yang akan didapatkan adalah false

	// akan menghasilkan return true
	fmt.Println(maxFloat32+10)
	fmt.Println(maxFloat32+2000000)

//	convert int ke float
	var myint int
	myint = 6
	myfloat := float64(myint)
	fmt.Println(myfloat) // 6

//	convert float ke int
	var myfloat2 float64
	myfloat2 = 6
	myint2 := int(myfloat2)
	fmt.Println(myint2) // 6
}
