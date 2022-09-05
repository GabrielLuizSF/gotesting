package main

import(
	"fmt"
	"math/sum"
	"math/div"
)


var userinput string;
var value1, value2 float64;

func GetUserinput(){
	fmt.Println("\nSelect The First Value:")
      fmt.Scanln(&value1)
	fmt.Println("\nSelect The Second Value:")
	  fmt.Scanln(&value2)
}

func main(){
 
   fmt.Println("Calculator:\n\n")
   fmt.Println("Options:{\n\nSum:'sum'\n\n}")
   fmt.Scanln(&userinput)

   switch userinput{
     case "sum":
		    GetUserinput()
		    addition := sum.Addition(value1,value2); 
		    fmt.Println("\n[Result]:",addition)
     default:
	   fmt.Println("[ERROR]:Invalid Value")
	break;
}
}