package memelord

type Client struct {
	apiToken string
	ApiUrl   string
}

func CreateClient(apiUrl string, apiToken string) Client {
	return Client{
		apiToken: apiToken,
		ApiUrl:   apiUrl,
	}
}
