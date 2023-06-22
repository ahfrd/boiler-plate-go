package response

type ReadResponse struct {
	Id          string         `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Category    []CategoryBook `json:"category"`
	Keyword     []KeywordBook  `json:"keyword"`
	Price       string         `json:"price"`
	Stock       int            `json:"stock"`
	Publisher   string         `json:"publisher"`
}

type ReadQueryEntity struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Keyword     string `json:"keyword"`
	Price       string `json:"price"`
	Stock       int    `json:"stock"`
	Publisher   string `json:"publisher"`
}

type CategoryBook struct {
	Data string `json:"data"`
}

type KeywordBook struct {
	Data string `json:"data"`
}
type GeneralResponse struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type GeneralResponseTest struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
