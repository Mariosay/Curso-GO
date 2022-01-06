package main


import (
	"fmt"
	"bufio"
	"os"
)

func main() {
 	/*var nombre string
	fmt.Println("Escrie tu nombre: " )
	fmt.Scanf("%s\n",&nombre)
	fmt.Printf("Tu nombre es: %s\n",nombre)*/
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingresa tu nombre: ")
	nombre,err :=reader.ReadString('\n')
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Hola "+ nombre)
	}
}