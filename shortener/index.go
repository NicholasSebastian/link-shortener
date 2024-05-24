package shortener

type LinkShortener struct {
	urls map[string]string
}

func New() *LinkShortener {
	return &LinkShortener{
		urls: make(map[string]string),
	}
}
