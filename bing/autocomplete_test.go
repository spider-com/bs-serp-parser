package bing

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func TestAutocomplete(t *testing.T) {
	f, err := ioutil.ReadFile("./autocomplete.html")
	require.Nil(t, err)

	res, err := parseAutocomplete(bytes.NewReader(f))
	require.Nil(t, err)
	require.Equal(t, 8, len(res))
	require.Equal(t, "pizza", res[0])
	require.Equal(t, "pizza delivery near me", res[4])
}
