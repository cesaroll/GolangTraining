package main

import "fmt"

func max(x int) int {
    return x * 2
}

func main() {
    
    max := max(7) //We don;t do this; bad coding practive to shadow variables/functions
    fmt.Println(max)

}

