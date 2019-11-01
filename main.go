package firepower

import (
	"context"
	"fmt"
)

func main() {
	client := NewClient(nil)
	objects, _, _ := client.Object.listObjects(context.Background(), "github", nil)

	if objects != nil {
		fmt.Println("dsfsdfdsf")
	}
}
