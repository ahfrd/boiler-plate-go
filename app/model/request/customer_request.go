package request

type GetListCustomerRequest struct {
	PageNumber string `json:"pageNumber" binding:"required"`
	PageSize   string `json:"pageSize" binding:"required"`
}

type AddCustomerRequest struct {
	CustomerName string `json:"customerName"`
	Address      string `json:"address"`
}
