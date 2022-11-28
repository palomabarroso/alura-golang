package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "Paloma"
	age := 29
	version := 1.0

	fmt.Println("Hello ", name)
	fmt.Println("Your age is: ", age)
	fmt.Println("The version of this program is: ", version)
	fmt.Println(reflect.TypeOf(version))

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var cmd int
	fmt.Scan(&cmd)

	fmt.Println("O cmd escolhido foi", cmd)
	fmt.Println("O endereço de memória do cmd é: ", &cmd)
}
