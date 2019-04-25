/* 4_index: Item by Index */
// Lists the items of the slice (vector), separated by commas

/* Main project! */
package main
    
import "fmt" // Println

func main(){
    sl := []string{"Hello", "world!", "What", "a", "beautiful", "day."}

    for i:=0; i<len(sl); i++ { fmt.Print(sl[i], ",");};

    fmt.Println("\n", sl); // Print that array!
}
