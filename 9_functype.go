/* 9_functype: Returning funcs! */

/* Main project! */
package main
    
import "fmt" // stdio

func math(x int) func(int) int {
    if x==1 { return func(a int) int{ return a+1 } } // incr
    
    return func(a int) int{ return a*a } // sq
}

func main(){
    f1 := math(1)
    //f2 := math(2)

    fmt.Println(f1(1336), math(2)(5))
}
