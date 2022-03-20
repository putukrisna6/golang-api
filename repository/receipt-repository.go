package repository

import (
	"github.com/putukrisna6/golang-api/entity"
	"gorm.io/gorm"
)

type ReceiptRepository interface {
	InsertReceipt(receipt entity.Receipt) entity.Receipt
	UpdateReceipt(receipt entity.Receipt) entity.Receipt
	ShowReceipt(receiptID uint64) entity.Receipt
	DeleteReceipt(receipt entity.Receipt)
	AllReceipts() []entity.Receipt
}

type receiptConnection struct {
	connection *gorm.DB
}

func NewReceiptRepository(db *gorm.DB) ReceiptRepository {
	return &receiptConnection{
		connection: db,
	}
}

func (db *receiptConnection) InsertReceipt(receipt entity.Receipt) entity.Receipt {
	db.connection.Save(&receipt)
	db.connection.Find(&receipt)
	return receipt
}

func (db *receiptConnection) UpdateReceipt(receipt entity.Receipt) entity.Receipt {
	var tempReceipt entity.Receipt
	db.connection.Find(&tempReceipt, receipt.ID)

	if receipt.Amount == 0 {
		receipt.Amount = tempReceipt.Amount
	}
	if receipt.Total == 0 {
		receipt.Total = tempReceipt.Total
	}

	db.connection.Save(&receipt)
	return receipt
}

func (db *receiptConnection) ShowReceipt(receiptID uint64) entity.Receipt {
	var receipt entity.Receipt
	db.connection.Find(&receipt, receiptID)
	return receipt
}

func (db *receiptConnection) DeleteReceipt(receipt entity.Receipt) {
	db.connection.Delete(&receipt)
}

func (db *receiptConnection) AllReceipts() []entity.Receipt {
	var receipts []entity.Receipt
	db.connection.Find(&receipts)
	return receipts
}
