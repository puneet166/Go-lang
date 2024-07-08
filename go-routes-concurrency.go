package main

import (  
    "fmt"
    "time"
)

func numbers() {  
    for i := 1; i <= 5; i++ {
        time.Sleep(250 * time.Millisecond)
        fmt.Printf("%d ", i)
    }
}
func alphabets() {  
    for i := 'a'; i <= 'e'; i++ {
        time.Sleep(400 * time.Millisecond)
        fmt.Printf("%c ", i)
    }
}
func main() {  
    // for thread and concureency run use go keywork followed by function name at the time of calling .
    go numbers()
        // for thread and concureency run use go keywork followed by function name at the time of calling .

    go alphabets()
    time.Sleep(3000 * time.Millisecond)
    fmt.Println("main terminated")
}
