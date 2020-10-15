package bing

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func TestParseMobile(t *testing.T) {
	f, err := ioutil.ReadFile("./mobile.html")
	require.Nil(t, err)

	res, err := parseMobile(bytes.NewReader(f))
	require.Nil(t, err)

	// some pages with ads don't have result count
	require.Equal(t, int64(0), res.TotalResultCount)

	require.Equal(t, 0, len(res.AdItems))
	require.Equal(t, 4, len(res.Items))
	require.Equal(t, "https://www.yelp.com/nearme/vegan-food", res.Items[0].URL)
	require.Equal(t, "Best Vegan Food Near Me - October 2020: Find Nearby Ve…", res.Items[0].Title)
	require.Equal(t, 0, res.Items[0].Position)
	require.Equal(t, "https://www.yelp.com/nearme/vegan-food", res.Items[0].DisplayURL)
	require.Empty(t, res.Items[0].Description)

	require.Equal(t, 0, len(res.RelatedQuestions))
}
