package task

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateTask(task Task) (Task, error) {
	return s.repo.Create(task)
}

func (s *Service) GetTask(id uint32) (Task, error) {
	return s.repo.Get(id)
}

func (s *Service) ListTasks() ([]Task, error) {
	return s.repo.List()
}

func (s *Service) ListTasksByUser(userID uint32) ([]Task, error) {
	return s.repo.ListByUser(userID)
}

func (s *Service) UpdateTask(task Task) (Task, error) {
	return s.repo.Update(task)
}

func (s *Service) DeleteTask(id uint32) error {
	return s.repo.Delete(id)
}
