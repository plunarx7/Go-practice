package main
import (
	"fmt"
	"bufio"
	"os"
	s "strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the string:")
	input, _ := reader.ReadString('\n')
	istr := s.ToLower(s.TrimSpace(input))
	if (s.HasPrefix(istr, "i") && s.HasSuffix(istr, "n") && s.Contains(istr, "a")) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
