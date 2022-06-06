/* Alta3 Research | Author: RZFeeser
   Types - inference   */

package main

import "fmt"

func main() {    
    // inference 
    i := 42              // int
    fmt.Printf("i is of type %T\n", i)
    // inference
    b := 45.89
    fmt.Printf("b is of type %T\n", b)
    // inference
    f := 3.142			 // float64
    fmt.Printf("f is of type %T\n", f)

    // inference
    v := 0.867 + 0.5i    // complex128
    fmt.Printf("v is of type %T\n", v)
}

