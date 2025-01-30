```go
package main

import "fmt"

func main() {
    var i interface{} = 10
    j := i.(int)
    fmt.Println(j)

    var k interface{} = "hello"
    l := k.(int) 
    fmt.Println(l)
}
```