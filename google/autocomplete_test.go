package google

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func TestAutocomplete(t *testing.T) {
	f, err := ioutil.ReadFile("./autocomplete.json")
	require.Nil(t, err)

	res, err := parseAutocomplete(bytes.NewReader(f))
	require.Equal(t, 10, len(res))
	require.Equal(t, "pizza house", res[0])
	require.Equal(t, "pizza33", res[1])
}
