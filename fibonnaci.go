package main

import "fmt"

func fibon() func() int{
a,b:=0,1
sum:=0
return func() int{
sum=a+b
a,b=b,sum
return sum
}


}
func main(){

f:=fibon()
for i:=0;i<=10;i++ {
fmt.Println(f())
}
}
