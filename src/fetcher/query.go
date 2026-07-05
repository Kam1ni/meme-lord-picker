package fetcher

import (
	"fmt"
	"meme-lord-picker/memelord"
	"strings"
)

func stringQueryToMemelordQuery(query string) memelord.Query {
	result := memelord.Query{
		PageSize: 100,
	}
	parts := strings.Split(query, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if len(part) == 0 {
			continue
		}
		if part[0] == '#' {
			tag := part[1:]
			if tag == "" {
				continue
			}
			result.Tags = append(result.Tags, strings.TrimSpace(tag))
		} else {
			result.Title += " " + part
		}
	}
	result.Title = strings.TrimSpace(result.Title)
	fmt.Println(result.Title)

	return result
}
