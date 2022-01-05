package main


import (
	"fmt"
	""
)

func main() {
 	var nombre string
	fmt.Println("Escrie tu nombre: " )
	fmt.Scanf("%s\n",&nombre)
	fmt.Printf("Tu nombre es: %s\n",nombre)
}