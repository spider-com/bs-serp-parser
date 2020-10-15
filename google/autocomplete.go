package google

import (
	"encoding/json"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"strings"
)

func parseAutocomplete(r io.Reader) (res []string, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}

	var response []interface{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return
	}
	if len(response) != 3 {
		return res, errors.New("bad response")
	}

	for _, s := range response[1].([]interface{}) {
		suggestions, ok := s.([]interface{})
		if !ok || len(suggestions) != 2 {
			continue
		}

		sRaw, ok := suggestions[0].(string)
		if !ok {
			continue
		}

		sug, err := goquery.NewDocumentFromReader(strings.NewReader(sRaw))
		if err != nil {
			return res, err
		}
		res = append(res, sug.Text())
	}

	return
}