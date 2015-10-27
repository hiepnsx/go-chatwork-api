package gochatwork

// Client is chatwork api client
type Client struct {
	config     *config
	connection apiConnection
}

type config struct {
	url   string
	token string
}

var chatworkURL = "https://api.chatwork.com"

// New return api client for www.chatwork.com
func New(token string) *Client {
	c := &config{
		url:   chatworkURL,
		token: token,
	}

	return &Client{
		config:     c,
		connection: &httpImp{},
	}
}
