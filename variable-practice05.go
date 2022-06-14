/* Alta3 Research | RZFeeser
   Variables - Shadowing */


package main

import (
    "fmt"
)

func main() {
    shadow := "world"
    able := 7
    test  := 2
    fmt.Println(shadow)     // world
    fmt.Println(test)
    fmt.Println(able)
    {
        shadow := "hello"// outer shadow is inaccessible from this point
        able := "welcome"
        fmt.Println(able)
        fmt.Println(shadow) // hello
    }

    fmt.Println(shadow)     // world
}

