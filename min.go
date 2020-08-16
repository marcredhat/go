package main
import "fmt"

func Min(a ...int) int {
    min := int(^uint(0) >> 1)  // largest int
    for _, i := range a {
        if i < min {
            min = i
        }
    }
    return min
}


func main() {
	fmt.Printf("%d", Min(3,5,7,101))
}
