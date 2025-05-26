package response

type CustomerInfoEntity struct {
	CustomerName string `json:"customerName"`
	Address      string `json:"address"`
}

type ListCustomerResponse struct {
	ListData   []CustomerInfoEntity `json:"listData"`
	Pagination CountData            `json:"pagination"`
}
