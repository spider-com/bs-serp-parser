package google

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func TestParseGoogleResult(t *testing.T) {
	googleDoc, err := ioutil.ReadFile("./index.html")
	require.Nil(t, err)

	res, err := ParseGoogleResult(bytes.NewReader(googleDoc))
	require.Nil(t, err)

	require.Equal(t, int64(3530000000), res.TotalResultCount)

	require.Equal(t, 12, len(res.OrganicItems))
	require.Equal(t, "https://www.techradar.com/news/best-iphone", res.OrganicItems[0].URL)
	require.Equal(t, "Jul 15, 2020 - Best iPhone: which one should you buy today? iPhone 11. (Image credit: Apple). 1. iPhone 11.", res.OrganicItems[0].Description)
	require.Equal(t, "www.techradar.com › news › best-iphonewww.techradar.com › news › best-iphone", res.OrganicItems[0].DisplayURL)
	require.Equal(t, 0, res.OrganicItems[0].Position)
	require.Equal(t, 0, res.OrganicItems[0].PositionOverall)
	require.Equal(t, "Best iPhone 2020: which Apple phone is the top choice for ...", res.OrganicItems[0].Title)

	require.Equal(t, 3, len(res.RelatedQuestions))
	require.Equal(t, "Which is the best iPhone to buy?", res.RelatedQuestions[0])

	require.Equal(t, 4, len(res.TopPLAItems))
	require.Equal(t, "Мобильный телефон Apple iPhone Xr 64GB", res.TopPLAItems[0].Title)
	require.Equal(t, "https://rozetka.com.ua/apple_mry52/p141517594/", res.TopPLAItems[0].URL)
	require.Equal(t, "UAH 21,499.00", res.TopPLAItems[0].Price)
	require.Equal(t, "ROZETKA", res.TopPLAItems[0].Source)

	require.Equal(t, int64(1), res.Pagination.Current)
	require.Equal(t, "https://www.google.com/search?q=iphone+which+one&ei=dLwqX9LSO5G53AOaoa-QDw&start=10&sa=N&ved=2ahUKEwjS3P_QnYTrAhWRHHcKHZrQC_IQ8NMDegQIDhBl", res.Pagination.Next)
	require.Equal(t, 9, len(res.Pagination.OtherPages))

	require.Equal(t, 1, len(res.PaidItems))
	require.Equal(t, "Айфоны в Rozetka - Купить, Отзывы, Цена - rozetka.com.ua", res.PaidItems[0].Title)
	require.Equal(t, "https://rozetka.com.ua/mobile-phones/c80003/preset=smartfon;producer=apple/", res.PaidItems[0].URL)
}