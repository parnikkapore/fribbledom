/* 11_typedef: struct */
// A breath of fresh air!

/* Main project! */
package main
    
import "fmt" // stdio
import "math" // pi

type Circle struct{
    radius float64
}

type Rect struct{
    width,height float64
}

type Cuboid struct{
    Rect // extends Rect
    length float64
}

func (c Circle) area() float64{ // Give Circle an area() method
    return math.Pi * c.radius * c.radius;
}

func (r Rect) area() float64{
    return r.width * r.height
}

func (c Cuboid) volume() float64{
    return c.width * c.height * c.length
}

func main(){
    r := Rect{width:2 ,height:3}
    q := Cuboid{Rect: r, length: 5}

    fmt.Println(r.area(), q.volume())
}
