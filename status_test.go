package pubg

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func Test_ClientStatus(t *testing.T) {
	c := NewClient(os.Getenv("APIKEY"))
	require.True(t, c.Status())
}
