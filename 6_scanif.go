/* 6_scanif: Scan and If */

/* Main project! */
package main
    
import "fmt" // stdio

func main(){
    var n int; fmt.Scan(&n);
    
    if n>10 {
        fmt.Println("Bigger than 10")
    }else{
        fmt.Println("Less than or equal to 10")
    }
}
