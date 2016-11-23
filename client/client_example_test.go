package client

import "fmt"

func ExampleNewClient() {
	token := "my-circle-token"
	client := NewClient(token)
	fmt.Printf("Client Created: %+s\n", client)
	// Output:
	// Client Created: {"token":"my-circle-token","host":"https://circleci.com","port":443,"version":"v1"}
}
