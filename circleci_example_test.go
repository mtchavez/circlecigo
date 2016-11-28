package circleci

import "fmt"

func ExampleNewClient() {
	token := "my-circle-token"
	client := NewClient(token)
	fmt.Printf("Client Created: %+s\n", client)
	// Output:
	// Client Created: {"base_url":"https://circleci.com/api/v1/","token":"my-circle-token"}
}
