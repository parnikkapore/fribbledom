/* 10_sliceop: Slice print function */
// EXPAND DONG ... !

/* Main project! */
package main
    
import "fmt" // stdio

func printer(sl ...string){ // Any number of strings, served in a slice
    for _,s := range sl {
        fmt.Println(s);
    }
}

func main(){
    sl := []string{"Hello", "World,"}

    printer(sl...) // Expand the slice into a bunch of args
    printer("How", "are", "you", "today?") // Expands go on their own
}
