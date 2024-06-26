package usecase

import (
	"Survialander/clean-arch-challenge/internal/entity"
)

type ListOrdersUseCase struct {
	Repository entity.OrderRepositoryInterface
}

type ListOrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

func NewListOrdersUseCase(
	orderRepository entity.OrderRepositoryInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		Repository: orderRepository,
	}
}

func (c *ListOrdersUseCase) Execute() ([]ListOrderOutputDTO, error) {
	orders, err := c.Repository.List()
	if err != nil {
		return []ListOrderOutputDTO{}, err
	}

	var ordersOutput []ListOrderOutputDTO
	for _, order := range orders {
		ordersOutput = append(ordersOutput, ListOrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return ordersOutput, nil
}
