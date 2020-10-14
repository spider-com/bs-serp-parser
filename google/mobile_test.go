package google

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func TestParseMobile(t *testing.T) {
	googleDoc, err := ioutil.ReadFile("./mobile.html")
	require.Nil(t, err)

	res, err := parseMobile(bytes.NewReader(googleDoc))
	require.Nil(t, err)

	require.Equal(t, 8, len(res.OrganicItems))
	require.Equal(t, "https://www.evian.com/en_us/the-evian-story/", res.OrganicItems[0].URL)
	require.NotEmpty(t, res.OrganicItems[0].Description)
	require.Equal(t, "Evian", res.OrganicItems[0].DisplayURL)
	require.Equal(t, 0, res.OrganicItems[0].Position)
	require.Equal(t, 0, res.OrganicItems[0].PositionOverall)
	require.Equal(t, "The evian® Story", res.OrganicItems[0].Title)
	require.True(t, res.OrganicItems[2].IsAMP)

	require.Equal(t, 4, len(res.RelatedQuestions))
	require.Equal(t, "What is the source of Evian water?", res.RelatedQuestions[0])
}