package core

type Option struct {
	SecretKey string // Notion api secret key
}

type Client struct {
	option *Option
}

func NewClient(opt *Option) (*Client, error) {
	return &Client{option: opt}, nil
}
