package main

import "fmt"
import "bufio"
import "os"

func main() {

    s := bufio.NewScanner(os.Stdin)
    s.Split(bufio.ScanLines)
    s.Scan()

    fmt.Println("Hello, World.")
    fmt.Println(s.Text())
}
