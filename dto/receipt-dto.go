package dto

type ReceiptCreateDTO struct {
	Amount uint    `json:"amount" form:"amount" binding:"required"`
	Total  float64 `json:"total" form:"total" binding:"required"`
}

type ReceiptUpdateDTO struct {
	ID     uint64  `json:"id" form:"id" binding:"required"`
	Amount uint    `json:"amount" form:"amount"`
	Total  float64 `json:"total" form:"total"`
}
