package main

import (
    "fmt"
    "os"
)

func main() {
    var i int 
    fmt.Print("Enter an osd number: ")
    fmt.Scan(&i)
    if i == 0 {
        fmt.Println(i, "is not a valid osd number.")
        os.Exit(1)
    }

    fmt.Println("You entered:", i)
}
