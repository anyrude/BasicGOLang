package main

import "fmt"
import "errors"
func f(a int) (int, error){
if a==42 {
return -1, errors.New(" 42 is baaeeed ")
} else if a == 32 {
return -1, errors.New(" 32 is okayyyish ")
} else {
return a + 42 ,nil
}
}

func main(){
a:=[]int{42,32,43}

for _,e:= range a {
if i,r:=f(e); r != nil{
 fmt.Println("Resulting error: ",r)
} else{ 
  fmt.Println("Working fine: ",i)
}
}
}

