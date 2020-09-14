package Services

import (
	Repository "../../Infrastructure/Repository"
)

type PriceService struct {
	priceRepository *Repository.PriceRepository
}

func (priceService *PriceService) Test() string {
	return priceService.priceRepository.Test()
}

func NewPriceService(_priceRepository *Repository.PriceRepository) *PriceService {
	return &PriceService{
		priceRepository: _priceRepository,
	}
}
