package types_test

import (
	"testing"

	"github.com/lucaspari/go-api/types"
)

func TestServer(t *testing.T) {
	t.Run("Test Server struct creation", func(t *testing.T) {
		port := ":8080"
		server := types.NewServer(port)
		if server == nil {
			t.Fatalf("Error creating server")
		}
	})
}
