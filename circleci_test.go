package circleci

import "testing"

const TestToken = "abcd-1234-abcd-1234"

func TestNewClient(t *testing.T) {
	client := NewClient(TestToken)
	if client.Token != TestToken {
		t.Errorf("Expected Token to be set to %s but got %s", TestToken, client.Token)
	}
}
