package main
import (
	"fmt"
	"os"
	"bufio"
	"encoding/json"
	s "strings"
)

func main() {
	var person = make(map[string]string)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Print("Enter the address: ")
	address, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	person["name"] = s.TrimSuffix(name, "\n")
	person["address"] = s.TrimSuffix(address, "\n")
	p, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling json:", err)
		os.Exit(1)
	}
	fmt.Println(string(p))
}
