package product

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateProduct(name, desc string, price float64, categoryID int) (*Product, error) {
	product := &Product{
		Name:        name,
		Description: desc,
		Price:       price,
		CategoryID:  categoryID,
	}

	err := s.repo.Create(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Service) GetProducts() ([]Product, error) {
	return s.repo.GetAll()
}
