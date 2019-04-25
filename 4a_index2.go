/* 4a_index: Item by Index (but with forrange) */
// Lists the items of the slice (vector), separated by commas

/* Main project! */
package main
    
import "fmt" // Println

func main(){
    sl := []string{"Hello", "world!", "What", "a", "beautiful", "day."}

    for _, x := range sl { fmt.Print(x, ",");} // index, element

    fmt.Println("\n", sl[:2], sl[2], sl[1:3], sl[1:]); // Guess what this does!
} // items i[1] thru i[2], because lazy ^
  //                                [1] thru end ^
