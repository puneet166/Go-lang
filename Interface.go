// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import "fmt"

type inter interface{
     age(age int) int
     height(height int) int
}
func age(age int) int{
    return age;
}
func height(height int) int{
    return height;
}
func main() {
    fmt.Println("-----line no 18-----",age(5));
    fmt.Println("-----line no 19-----",height(10));
}
