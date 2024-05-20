package shortener

type LinkShortener struct {
	host string
	urls map[string]string
}

func New(host string) *LinkShortener {
	return &LinkShortener{
		host: host,
		urls: make(map[string]string),
	}
}
