package main

import "fmt";

const saysHelloInEnglish = "Hello "

func Hello(name string)string{
    if name == ""{
	 name = "World"
	}
	return saysHelloInEnglish + name
}

func main(){
 var name string;
 fmt.Print("[Your Name]:")
 fmt.Scanln(&name)
 fmt.Println(Hello(name))
}