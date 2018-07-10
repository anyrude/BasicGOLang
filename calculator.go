package main

import "fmt"

func main(){
var a float32
var b float32
var c string
var d float32
fmt.Scanf("%f" , &a)
fmt.Scanf("%f" , &b)
fmt.Scanf("%s" , &c)
switch c{
case "+" : 
     fmt.Println(a+b)
case "-" :
     fmt.Println(a-b)
case "*" :
     fmt.Println(a*b)
case "/" :
	d=a/b
    
    fmt.Println(d)
}
}


