package gochatwork

// Client is chatwork api client
type Client struct {
	config *config
	connection   apiConnection
}

type config struct {
	url   string
	token string
}

var chatworkURL = "https://api.chatwork.com"
var kddiChatworkURL = "https://kcw.kddi.ne.jp"

// New return api client for www.chatwork.com
func New(token string) *Client {
	return newClient(token, chatworkURL)
}

// NewKddiChatwork return api client for kcw.kddi.ne.jp
func NewKddiChatwork(token string) *Client {
	return newClient(token, kddiChatworkURL)
}

func newClient(token string, url string) *Client {
	c := &config{
		url:   url,
		token: token,
	}

	return &Client{
		config: c,
		connection:   &httpImp{},
	}
}

// MeRaw return /me response by []byte
func (c *Client) MeRaw() ([]byte, error) {
	return c.connection.Get("me", c.config)
}
