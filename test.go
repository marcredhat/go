package main
import "fmt"


func unhex(c byte) byte {
    switch {
    case '0' <= c && c <= '9':
        return c - '0'
    case 'a' <= c && c <= 'f':
        return c - 'a' + 10
    case 'A' <= c && c <= 'F':
        return c - 'A' + 10
    }
    return 0
}


// Compare returns an integer comparing the two byte slices,
// lexicographically.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b
func Compare(a, b []byte) int {
    for i := 0; i < len(a) && i < len(b); i++ {
        switch {
        case a[i] > b[i]:
            return 1
        case a[i] < b[i]:
            return -1
        }
    }
    switch {
    case len(a) > len(b):
        return 1
    case len(a) < len(b):
        return -1
    }
    return 0
}


func main() {
    fmt.Println("hello world")
    for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
    	fmt.Printf("character %#U starts at byte position %d\n", char, pos)
    }
    b :=  unhex('B')
    fmt.Printf("%x",b)
    byteslice1 := []byte("Here is a string")
    byteslice2 := []byte("Here is a string that's longer")
    fmt.Printf("%d", Compare(byteslice1,byteslice2))
    fmt.Println("end")
}
