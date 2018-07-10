package main

import "fmt"

func fac(k int) int {
if k== 0{
return 1}
return k  * fac(k-1)
}
func main(){
fmt.Println(fac(5))
}

