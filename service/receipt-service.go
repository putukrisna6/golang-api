package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/putukrisna6/golang-api/dto"
	"github.com/putukrisna6/golang-api/entity"
	"github.com/putukrisna6/golang-api/repository"
)

type ReceiptService interface {
	Insert(r dto.ReceiptCreateDTO) entity.Receipt
	Update(r dto.ReceiptUpdateDTO) entity.Receipt
	Show(receiptID uint64) entity.Receipt
	Delete(r entity.Receipt)
	All() []entity.Receipt
}

type receiptService struct {
	receiptRepository repository.ReceiptRepository
}

func NewReceiptService(receiptRepository repository.ReceiptRepository) ReceiptService {
	return &receiptService{
		receiptRepository: receiptRepository,
	}
}

func (service *receiptService) Insert(r dto.ReceiptCreateDTO) entity.Receipt {
	newReceipt := entity.Receipt{}
	err := smapping.FillStruct(&newReceipt, smapping.MapFields(&r))
	if err != nil {
		log.Fatalf("failed to map: %v", err.Error())
	}

	return service.receiptRepository.InsertReceipt(newReceipt)
}

func (service *receiptService) Update(r dto.ReceiptUpdateDTO) entity.Receipt {
	receipt := entity.Receipt{}
	err := smapping.FillStruct(&receipt, smapping.MapFields(&r))
	if err != nil {
		log.Fatalf("failed to map: %v", err.Error())
	}

	return service.receiptRepository.UpdateReceipt(receipt)
}

func (service *receiptService) Show(receiptID uint64) entity.Receipt {
	return service.receiptRepository.ShowReceipt(receiptID)
}

func (service *receiptService) Delete(receipt entity.Receipt) {
	service.receiptRepository.DeleteReceipt(receipt)
}

func (service *receiptService) All() []entity.Receipt {
	return service.receiptRepository.AllReceipts()
}
