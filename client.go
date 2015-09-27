package gochatwork

type Client struct {
    config *config
}

type config struct{
    url string
    token string
}

var chatworkURL = "www.chatwork.com"
var kddiChatworkURL = "kcw.kddi.ne.jp"

func New(token string) *Client {
    return newClient(token, chatworkURL)
}

func NewKddiChatwork(token string) *Client {
    return newClient(token, kddiChatworkURL)
}

func newClient(token string, url string) *Client {
    c := &config{
        url: url,
        token: token,
    }

    return &Client{
        config: c,
    }
}