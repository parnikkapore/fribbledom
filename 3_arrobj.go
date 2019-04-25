/* 3_arrobj: The array's moveset */
// ("Array Object", bcause it's totally different from CPP)

/* Main project! */
package main
    
import "fmt" // Println

func main(){
    sl := []string{"Hello", "world!"}

    sl = append(sl, "What", "a", "beautiful", "day.") // strcat but it returns a value
    // alas, first entry has to be a "slice", a la our array

    fmt.Println(sl); // Print that array!
}
