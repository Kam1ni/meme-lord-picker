package memelord

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func formatDate(date time.Time) string {
	year, month, d := date.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, int(month), d)
}

func (q Query) toQueryParamets() url.Values {
	result := url.Values{}
	if q.Username != "" {
		result.Set("username", q.Username)
	}
	if q.Title != "" {
		result.Set("title", q.Title)
	}
	for _, t := range q.Tags {
		result.Add("tag", strings.ReplaceAll(strings.ToLower(t), " ", "-"))
	}
	if q.DateFrom != nil {
		result.Set("date_from", formatDate(*q.DateFrom))
	}
	if q.DateTo != nil {
		result.Set("date_to", formatDate(*q.DateTo))
	}
	if q.MediaType != NO_MEDIA_TYPE {
		result.Set("media_type", string(q.MediaType))
	}
	if q.Ordering != NO_ORDERING {
		result.Set("ordering", string(q.Ordering))
	}
	if q.Page != 0 {
		result.Set("page", fmt.Sprintf("%d", q.Page))
	}
	if q.PageSize != 0 {
		result.Set("page_size", fmt.Sprintf("%d", q.PageSize))
	}
	return result
}

func handleApiError(response *http.Response) error {
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("Http request has response status %d and failed to read response body\n%s", response.StatusCode, err.Error())
	}
	errorMessage := ErrorResponse{}
	err = json.Unmarshal(bodyBytes, &errorMessage)
	if err != nil {
		return fmt.Errorf("Http request has response status %d\n%s", response.StatusCode, string(bodyBytes))
	}
	return errors.New(errorMessage.Error)
}

func (c Client) FetchMemes(query Query) (MemesResponse, error) {
	url := fmt.Sprintf("%s/api/memes/?%s", c.ApiUrl, query.toQueryParamets().Encode())
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return MemesResponse{}, fmt.Errorf("Failed to create http request\n%s", err.Error())
	}
	fmt.Println(url)
	request.Header.Set("Authorization", "Bearer "+c.apiToken)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return MemesResponse{}, fmt.Errorf("Failed to do request\n%s", err.Error())
	}
	if response.StatusCode != 200 {
		return MemesResponse{}, handleApiError(response)
	}
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return MemesResponse{}, fmt.Errorf("Api had response 200 but failed to read response body\n%s", err.Error())
	}
	result := MemesResponse{}
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return MemesResponse{}, fmt.Errorf("Api had response 200 but failed to parse response body\n%s", err.Error())
	}
	return result, nil
}

func (c Client) FetcRandomhMeme(query Query) (Meme, error) {
	url := fmt.Sprintf("%s/api/memes/random/?%s", c.ApiUrl, query.toQueryParamets())
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Meme{}, fmt.Errorf("Failed to create http request\n%s", err.Error())
	}
	request.Header.Set("Authorization", "Bearer "+c.apiToken)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return Meme{}, fmt.Errorf("Failed to do request\n%s", err.Error())
	}
	if response.StatusCode != 200 {
		return Meme{}, handleApiError(response)
	}
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return Meme{}, fmt.Errorf("Api had response 200 but failed to read response body\n%s", err.Error())
	}
	result := Meme{}
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return Meme{}, fmt.Errorf("Api had response 200 but failed to parse response body\n%s", err.Error())
	}
	return result, nil
}

// Returns image bytes and mimetype
func DownloadImage(url string) ([]byte, string, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, "", fmt.Errorf("Failed to create request\n%s", err.Error())
	}
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, "", fmt.Errorf("Failed to do http request\n%s", err.Error())
	}
	if res.StatusCode != 200 {
		return nil, "", handleApiError(res)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, "", fmt.Errorf("Failed to read response body\n%s", err.Error())
	}
	return body, http.DetectContentType(body), nil
}

func (m Meme) DownloadImage() ([]byte, string, error) {
	return DownloadImage(m.ImageUrl)
}
