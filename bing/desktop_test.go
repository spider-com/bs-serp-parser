package bing

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func TestParseAds(t *testing.T) {
	googleDoc, err := ioutil.ReadFile("./assets/ads.html")
	require.Nil(t, err)

	res, err := parse(bytes.NewReader(googleDoc))
	require.Nil(t, err)

	// some pages with ads don't have result count
	require.Equal(t, int64(0), res.TotalResultCount)

	require.Equal(t, 3, len(res.AdItems))
	require.Equal(t, "https://www.samsung.com/Access", res.AdItems[0].URL)
	require.Equal(t, "Samsung Access For Mobile - Starting from $25.61/mo.", res.AdItems[0].Title)
	require.Equal(t, 0, res.AdItems[0].Position)
	require.Equal(t, "https://www.samsung.com/Access", res.AdItems[0].DisplayURL)
	require.NotEqual(t, "", res.AdItems[0].Description)

	require.Equal(t, 6, len(res.Items))
	require.Equal(t, "www.apple.com/shop/buy-iphone/iphone-11", res.Items[0].URL)
	require.Equal(t, "Buy iPhone 11 - Apple", res.Items[0].Title)
	require.Equal(t, 0, res.Items[0].Position)
	require.Equal(t, "www.apple.com/shop/buy-iphone/iphone-11", res.Items[0].DisplayURL)
	require.NotEqual(t, "", res.Items[0].Description)

	require.Equal(t, 4, len(res.AdBottomItems))
	require.Equal(t, "https://www.apple.com", res.AdBottomItems[0].URL)
	require.Equal(t, "iPhone 11 - From $479 with trade-in", res.AdBottomItems[0].Title)
	require.Equal(t, 0, res.AdBottomItems[0].Position)
	require.Equal(t, "https://www.apple.com", res.AdBottomItems[0].DisplayURL)
	require.NotEqual(t, "", res.AdBottomItems[0].Description)

	require.Equal(t, 0, len(res.RelatedQuestions))

	require.Equal(t, int64(1), res.Pagination.Current)
	require.NotEmpty(t, res.Pagination.Next)
	require.Equal(t, 4, len(res.Pagination.OtherPages))
}
