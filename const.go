package main

import ("fmt"
"math") 

func main() {
    const a string= "helo"
    //a="whoa"
    fmt.Println(a)//print instead of fmt.print has undefined behaviour
    const b float32= 11.4 
    fmt.Printf("%T\n",float64(b))
    print(math.Sin(float64(b)))
}
