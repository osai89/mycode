/* Alta3 Research | RZFeeser
   Variables - Shadowing */


package main

import (
    "fmt"
)

func main() {
    shadow := "world"
    test  := 2
    fmt.Println(shadow)     // world
    fmt.Println(test)
    {
        shadow := "hello"   // outer shadow is inaccessible from this point
        fmt.Println(shadow) // hello
    }

    fmt.Println(shadow)     // world
}

