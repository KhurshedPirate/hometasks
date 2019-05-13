package main
import "fmt"
func main(){
    var x int
    fmt.Print("enter length in m ")
    fmt.Scan(&x)
    fmt.Print("length in km ",x/1000)
  }
