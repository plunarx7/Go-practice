package main
import "fmt"
func main() {
	var fnum float32
	var inum int
	fmt.Printf("Enter a floating point number: ")
	fmt.Scan(&fnum)
	inum = int(fnum)
	fmt.Printf("Truncated number: %d", inum)
}
