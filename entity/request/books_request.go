package request

type CreateRequest struct {
	Username    string         `json:"username"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Category    []CategoryBook `json:"category"`
	Keyword     []KeywordBook  `json:"keyword"`
	Price       string         `json:"price"`
	Stock       int            `json:"stock"`
	Publisher   string         `json:"publisher"`
}

type UpdateRequest struct {
	Username    string         `json:"username"`
	Id          string         `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Category    []CategoryBook `json:"category"`
	Keyword     []KeywordBook  `json:"keyword"`
	Price       string         `json:"price"`
	Stock       int            `json:"stock"`
	Publisher   string         `json:"publisher"`
}
type IdArrayRequest struct {
	Data string `json:"data"`
}
type DeleteRequest struct {
	Username string           `json:"username"`
	Id       []IdArrayRequest `json:"id"`
}

type ReadRequest struct {
	Username string `json:"username"`
	Id       string `json:"id"`
}

type CategoryBook struct {
	Data string `json:"data"`
}

type KeywordBook struct {
	Data string `json:"data"`
}
