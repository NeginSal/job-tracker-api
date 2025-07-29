package job

import (
	"github.com/google/uuid"
)

type Service struct{
	repo *Repository
}

func NewService(r *Repository) *Service{
	return &Service{r}
}

func (s *Service)CreateJob(job *Job)error{
	job.ID = uuid.New()
	return s.repo.Create(job)
}

func (s *Service)GetJobsByUser(userID string)([]Job , error){
	return s.repo.GetAllByUser(userID)
}

funct (s *Service)DeleteJob (id uuid.UUID, userID string)error{
	return s.repo.Delete(id ,userID)
}