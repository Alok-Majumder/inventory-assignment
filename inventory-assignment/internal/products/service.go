package products

import "context"

type Service struct {
	productRepository ProductRepository
}

func NewService(productRepository ProductRepository) *Service {
	return &Service{
		productRepository,
	}
}

func (s *Service) GetAllProducts(ctx context.Context) ([]*ProductQty, error) {

	return s.productRepository.GetAllProductsWithQuantity(ctx)

}

func (s *Service) SellProduct(ctx context.Context, product []byte) (*Response, error) {
	res, err := s.productRepository.SellProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	return res, nil

}
