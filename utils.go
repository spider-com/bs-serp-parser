package serp

import (
	"strings"
)

func PrependDomainToHRef(domain string, href string) string {
	if strings.HasPrefix(href, "/search") {
		href = domain + href
	}

	return href
}
