package pubg

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_ClientStatus(t *testing.T) {
	require.True(t, cTest.Status())
}
