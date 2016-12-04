package circleci

import (
	"fmt"
	"os"
)

func ExampleClient_Me() {
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

func ExampleClient_Me_unauthorized() {
	client := defaultClient
	_, apiResp := client.Me()
	fmt.Printf(`[Response]
Success = %v
`, apiResp.Success())
	// Output:
	// [Response]
	// Success = false
}
