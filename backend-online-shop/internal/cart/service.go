package cart

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Add(userID, productID, quantity int) error {
	return s.repo.Add(userID, productID, quantity)
}

func (s *Service) Get(userID int) ([]CartItem, error) {
	return s.repo.GetByUser(userID)
}
