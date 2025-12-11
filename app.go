package main

import "fmt"

func main() {
	ten := "John Doe"
	tuoi := 30
	chieuCao := 1.75
	daTotNghiep := true
	phanTramHocLuc := 85.5

	fmt.Printf("Kiểu dữ liệu của ten là: %T\n", ten)
	fmt.Printf("Kiểu dữ liệu của tuoi là: %T\n", tuoi)
	fmt.Printf("Kiểu dữ liệu của tuoi là: %T\n", chieuCao)
	fmt.Printf("Kiểu dữ liệu của tuoi là: %T\n", daTotNghiep)
	fmt.Printf("Kiểu dữ liệu của tuoi là: %T\n", phanTramHocLuc)

	fmt.Printf("Giá trị của ten là: %v\n", ten)
	fmt.Printf("Giá trị của tuoi là: %v\n", tuoi)
	fmt.Printf("Giá trị của tuoi là: %v\n", chieuCao)
	fmt.Printf("Giá trị của tuoi là: %v\n", daTotNghiep)
	fmt.Printf("Giá trị của tuoi là: %v\n", phanTramHocLuc)

	fmt.Printf("My name is %s and my age is %d", ten, tuoi)
}
