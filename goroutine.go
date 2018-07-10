package main

import "fmt"

func yo(any string){
for i:=0;i<3;i++ {
fmt.Println(any,i)
}
}

func main(){

yo("bruh: ")

go yo("no bruh: ")

go func(msg string) {
        fmt.Println(msg)
}("going")


fmt.Println("bye")
}
