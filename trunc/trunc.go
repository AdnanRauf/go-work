// Golang program to show how 
// to take input from the user a float and display trucated number
package main 
  
import (
    "fmt"
    "strconv"
)
  
// main function 
func main() { 
  
    // Println function is used to 
    // display output in the next line 
    fmt.Println("Enter a floating point number: ") 
  
    // var then variable name then variable type 
    var num float64 
  
    // Taking input from user 
    fmt.Scanln(&num) 
	num_formatted := strconv.FormatFloat(num, 'g', 5, 64)
	
  
    // Print function is used to 
    // display output in the same line 
    fmt.Print("float number entered is: "+num_formatted+ "\n") 
	fmt.Printf("truncated number is %.0f\n", num) 
  
    
} 
