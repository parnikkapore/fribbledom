/* 5_forezore: All the fors! */
// 3 types of a classic for loop

/* Main project! */
package main
    
import "fmt" // Println

func main(){
    var i int;

    fmt.Println("Type 1!")
    for i=0; i<5; i++ {
        fmt.Println("i =", i)
    }

    fmt.Println("Type 2!")
    i = 0
    for i<5{
        fmt.Println("i =", i)
        i++
    }

    fmt.Println("Type 3!");
    i = 0
    for{ //ever, evermore!!! BA DUM TSH
        if i>=5 { break; }
        fmt.Println("i =", i)
        i++
    }
}
