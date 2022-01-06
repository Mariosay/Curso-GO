package main 

import "fmt"


func main(){
	x:= 10
	y:= 10
	if (x>y) {
		fmt.Printf("%d es mayor que %d \n",x,y)
	}else if x<y {
		fmt.Printf("%d es mayor que %d \n",y,x)
	}else{
		fmt.Println("son iguales")
	}
}

/*
== igual a 
!= diferente de
< menor que 
> mayon que
>= mayor o igual que
<= menor o igual que
&& AND
|| OR
 */