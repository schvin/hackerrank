package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func main() {
    var i uint64 = 4
    var d float64 = 4.0
    var s string = "HackerRank "

    scanner := bufio.NewScanner(os.Stdin)

    // Declare second integer, double, and String variables.
    var i2 uint64
    var d2 float64
    var s2 string
    
    // Read and save an integer, double, and String to your variables.
    scanner.Split(bufio.ScanLines)
    scanner.Scan()
    i2 = scanner.Text()
    
    // Print the sum of both integer variables on a new line.
    
    // Print the sum of the double variables on a new line.
    
    // Concatenate and print the String variables on a new line
    // The 's' variable above should be printed first.

}
