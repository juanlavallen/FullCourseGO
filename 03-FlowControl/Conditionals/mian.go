package main

import "fmt"

func main() {

	var consumption float64
	var discount float64
	var data string

	fmt.Print("Insert consumption: ")
	fmt.Scanln(&consumption)

	if consumption >= 0 {
		data = "10%"
		discount = consumption * 0.10
	} else {
		fmt.Println("Error")
	}

	mount := consumption - discount
	igv := mount * 0.19
	total := mount + igv

	fmt.Println("##### APP #####")
	fmt.Println("Discount: ", data)
	fmt.Println("Consumption: ", consumption)
	fmt.Println("IGV: ", igv)
	fmt.Println("Total: ", total)
}
