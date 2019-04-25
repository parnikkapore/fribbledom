/* 8_defun: Defining a function */
// Because Lisp

/* Main project! */
package main
    
import "fmt" // stdio

func even(a int) (bool, bool){
    return a%2==0, true
}

func main(){
    incr := func(a int) int { // INLINE FUNCTIONS BABYYYYYYYYY!
        return a+1
    }

    isEven, err := even(5)
    incd := incr(9)

    fmt.Println(isEven, err, incd)
}
