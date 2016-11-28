package circleci

import (
	"fmt"
	"os"
)

func ExampleMe() {
	token := os.Getenv("CIRCLE_TOKEN")
	client := NewClient(token)
	_, apiResp := client.Me()
	fmt.Printf(`[Response]
Success = %v
`, apiResp.Success())
	// Output:
	// [Response]
	// Success = true
}

func ExampleMe_Unauthorized() {
	client := defaultClient
	_, apiResp := client.Me()
	fmt.Printf(`[Response]
Success = %v
`, apiResp.Success())
	// Output:
	// [Response]
	// Success = false
}
