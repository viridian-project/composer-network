package main
import "fmt"
type Relationship struct {
   Namespace string `json:"namespace"`
   Type string `json:"type"`
   Identifier string `json:"identifier"`
}
func main() {
   fmt.Printf("Hello, world.")
}
