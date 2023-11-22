package main

import (
    "fmt"
    "math"
)

func roundFloat(val float64, precision uint) float64 {
    ratio := math.Pow(10, float64(precision))
    return math.Round(val*ratio) / ratio
}

func main() {
    number := 12.3456789

    fmt.Println(roundFloat(number, 2))
    fmt.Println(roundFloat(number, 3))
    fmt.Println(roundFloat(number, 4))
    fmt.Println(roundFloat(number, 0))

    number = -12.3456789
    fmt.Println(roundFloat(number, 0))
    fmt.Println(roundFloat(number, 1))
    fmt.Println(roundFloat(number, 10))
}
