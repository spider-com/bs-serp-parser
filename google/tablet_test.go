package google

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func TestParseTablet(t *testing.T) {
	googleDoc, err := ioutil.ReadFile("./tablet.html")
	require.Nil(t, err)

	res, err := parseTablet(bytes.NewReader(googleDoc))
	require.Nil(t, err)

	require.Equal(t, 13, len(res.OrganicItems))
	require.Equal(t, "https://la.eater.com/maps/best-vegan-vegetarian-restaurants-los-angeles", res.OrganicItems[0].URL)
	require.NotEmpty(t, res.OrganicItems[0].Description)
	require.Equal(t, "la.eater.comla.eater.com", res.OrganicItems[0].DisplayURL)
	require.Equal(t, 0, res.OrganicItems[0].Position)
	require.Equal(t, 0, res.OrganicItems[0].PositionOverall)
	require.Equal(t, "The 24 Best Vegetarian Restaurants in Los Angeles - Eater LA14 Best Vegan Restaurants in Los Angeles - Eater LA", res.OrganicItems[0].Title)

	require.Equal(t, 4, len(res.RelatedQuestions))
	require.Equal(t, "Which country is best for vegans?", res.RelatedQuestions[0])
}