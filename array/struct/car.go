package main

import (
	"fmt"
)

type Car struct {
	Name, Model, Color string
	WeightInKg         float64
}

func main() {
	var (
		Name, Model, Color string
		Weight             float64
		car                [10]Car
		jumlah             int
	)
	fmt.Print("Masukkan Jumlah Data : ")
	fmt.Scanln(&jumlah)

	for i := 0; i < jumlah; i++ {
		fmt.Scanln(&Name, &Model, &Color, &Weight)
		car[i].Color = Color
		car[i].Name = Name
		car[i].Model = Model
		car[i].WeightInKg = Weight
	}

	fmt.Println("")

	for i := 0; i < jumlah; i++ {
		fmt.Println("No: ", i+1)
		fmt.Println("Car Name: ", car[i].Name)
		fmt.Println("Car Color: ", car[i].Color)
		fmt.Println("Car Model: ", car[i].Model)
		fmt.Println("Car Weight: ", car[i].WeightInKg)
		fmt.Println("")
	}

}
