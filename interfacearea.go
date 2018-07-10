package main

import "fmt"

type geometry interface{
area() int
per() int
}

type rect struct{
height int
width int
}

func (r rect) area() int{
return r.height * r.width
}

func (r rect) per() int{
return 2*(r.height + r.width)
}

func get(g geometry){
fmt.Println(g.area())
fmt.Println(g.per())
}

func main() {
s:=rect{2,3}
fmt.Println(s.height)
fmt.Println(s.width)
fmt.Println(s.area())

get(s)
}
