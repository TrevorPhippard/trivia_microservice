package game

type Client struct {
	conn    int
	service int
}

func NewClient(url string) (*Client, error) {
	a := 2
	b := 2
	return &Client{a, b}, nil
}
