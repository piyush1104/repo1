package use

import (
	"fmt"

	"github.com/piyush1104/repo2"
	"github.com/piyush1104/repo2/pkg/client"
)

func main() {
	client.GetApp()
	fmt.Println(repo2.App{ID: "hello"})
}
