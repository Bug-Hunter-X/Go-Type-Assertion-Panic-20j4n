```go
package main

import "fmt"

func main() {
    var i interface{} = 10
    j, ok := i.(int)
    if ok {
        fmt.Println(j)
    } else {
        fmt.Println("i is not an integer")
    }

    var k interface{} = "hello"
    l, ok := k.(int)
    if ok {
        fmt.Println(l)
    } else {
        fmt.Println("k is not an integer")
    }
    //Alternative using type switch
    switch v := k.(type) {
    case int:
        fmt.Println(v)
    case string:
        fmt.Println("k is a string", v) 
    default:
        fmt.Println("k is of unknown type")
    }
}
```