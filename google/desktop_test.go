package google

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func TestParse(t *testing.T) {
	googleDoc, err := ioutil.ReadFile("./desktop.html")
	require.Nil(t, err)

	res, err := parse(bytes.NewReader(googleDoc))
	require.Nil(t, err)

	require.Equal(t, int64(1440000000), res.TotalResultCount)

	require.Equal(t, 9, len(res.OrganicItems))
	require.Equal(t, "https://www.pizzahut.com/", res.OrganicItems[0].URL)
	require.NotEmpty(t, res.OrganicItems[0].Description)
	require.Equal(t, "www.pizzahut.comwww.pizzahut.com", res.OrganicItems[0].DisplayURL)
	require.Equal(t, 0, res.OrganicItems[0].Position)
	require.Equal(t, 0, res.OrganicItems[0].PositionOverall)
	require.Equal(t, "Pizza Hut: Pizza Delivery | Pizza Carryout | Coupons | Wings ...", res.OrganicItems[0].Title)

	require.Equal(t, 4, len(res.RelatedQuestions))
	require.Equal(t, "What defines a pizza?", res.RelatedQuestions[0])

	require.Equal(t, 0, len(res.TopPLAItems))

	require.Equal(t, int64(1), res.Pagination.Current)
	require.Equal(t, "https://www.google.com/search?q=pizza&ei=X919X822G9a5tAbXkbDgBg&start=10&sa=N&ved=2ahUKEwiN4obZ5KLsAhXWHM0KHdcIDGwQ8NMDegQIAxBC", res.Pagination.Next)
	require.Equal(t, 9, len(res.Pagination.OtherPages))

	require.Equal(t, 0, len(res.PaidItems))
}