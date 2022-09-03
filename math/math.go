package main

import(
	"fmt"
	"math/sum"
)


var userinput string;
var value1, value2 int;


func main(){
 
   fmt.Println("Calculator:\n\n")
   fmt.Println("Options:{\n\nSum:'sum'\n\n}")
   fmt.Scanln(&userinput)

   switch userinput{
     case "sum":
			fmt.Println("\nSelect The First Value:")
			fmt.Scanln(&value1)
			fmt.Println("\nSelect The Second Value:")
			fmt.Scanln(&value2)
			var addition int =	sum.Addition(value1,value2); 
		    fmt.Println("\n[Result]:",addition)
     default:
	   fmt.Println("[ERROR]:Invalid Value")
	break;
}
}