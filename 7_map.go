/* 7_map: Key-value pairs */

/* Main project! */
package main
    
import "fmt" // stdio

func main(){
    m := make(map[string] int)

    m["foo"] = 42
    m["bar"] = 1337

    fmt.Println(m["foo"], m["bar"], m["baar"])

    for k,v := range m {
        fmt.Printf("m.%s = %d\n", k, v);
    }

    // Check if an item exists
    var q string; fmt.Scan(&q)
    v, have := m[q];
    if have {
        fmt.Printf("%s is %d\n", q, v)
    }else{ // A key with that name doesn't exist (or some other error)
        fmt.Printf("You are a totally useless guesser\n")
    }
}
