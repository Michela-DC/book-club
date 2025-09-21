package controller

type CreateBookRequest struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Year   int    `json:"year"`
	Status string `json:"status"`
}

func (b *CreateBookRequest) validate() error {
	// TODO: IMPLEMENT
	return nil
}
