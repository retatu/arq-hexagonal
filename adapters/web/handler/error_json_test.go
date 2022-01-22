package handler

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHandlerJsonError(t *testing.T) {
	msg := "Hello JSON"
	result := jsonError(msg)
	require.Equal(t, []byte(`{"message":"Hello JSON"}`), result)
}
