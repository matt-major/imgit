package tinypng

// Response struct
type Response struct {
	Input struct {
		Size int
		Type string
	}

	Output struct {
		Size   int
		Type   string
		Width  int
		Height int
		Ratio  float64
		URL    string
	}
}

// Request struct
type Request struct {
	body []byte
}

// NewRequest return a new Request pointer
func NewRequest() (r *Request) {
	return &Request{}
}

// SetBody - set the request body value
func (r *Request) SetBody(b []byte) {
	r.body = b
}

// GetBody - get the request body value
func (r Request) GetBody() []byte {
	return r.body
}
