package order

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Checkout(userID int) (int, error) {
	orderID, err := s.repo.CreateOrder(userID)
	if err != nil {
		return 0, err
	}

	items, err := s.repo.GetCart(userID)
	if err != nil {
		return 0, err
	}

	for _, item := range items {
		item.OrderID = orderID

		err := s.repo.AddItem(item)
		if err != nil {
			return 0, err
		}
	}

	err = s.repo.ClearCart(userID)
	if err != nil {
		return 0, err
	}

	return orderID, nil
}
