package category

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateCategory(name string) (*Category, error) {
	return s.repo.Create(name)
}

func (s *Service) GetCategories() ([]Category, error) {
	return s.repo.GetAll()
}
