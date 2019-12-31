package main

import "fmt"

func main(){
	var f1,f2 ,op int
	fmt.Printf("ingrese primer número: ")
	fmt.Scanf("%d",&f1)
	fmt.Printf("ingrese segundo numero: ")
	fmt.Scanf("%d",&f2)
	fmt.Println("eliga la operacion a realizar :")
	fmt.Println("1-producto")
	fmt.Println("2-suma")
	fmt.Println("3-resta")
	fmt.Println("4-division")
	fmt.Scanf("%d",&op)
	menu(f1,f2,op)
}
func menu( f1 ,f2 , op int){
	switch op {
	case 1 :
		fmt.Printf("el productos de los numeros ingresados es : %d \n",producto(f1,f2))

	case 2 :
		
		fmt.Printf("el productos de los numeros ingresados es : %d \n",suma(f1,f2))
		

	case 3 :
		fmt.Printf("el productos de los numeros ingresados es : %d \n",resta(f1,f2))

		
	case 4 :

		if(f2 == 0){
			fmt.Printf("no se puede divir por cero \n")
		}else{

			fmt.Printf("el productos de los numeros ingresados es : %d \n",division(f1,f2))
		}

	
	}
}
func producto(f1,f2 int)int{
	return f1*f2
}

func suma(f1,f2 int) int{
	return f1+f2
}

func resta(f1,f2 int) int{
	
	return f1-f2
}
func division(f1,f2 int) int{
	
	return f1/f2
}