package service

import "golang/internal/repo"

type CompanyService struct {
}

type Company struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewCompanyService() *CompanyService {
	return &CompanyService{}
}

func (cs *CompanyService) GetCompanyByID() Company {
	result := repo.NewCompanyRepo().GetCompanyByID()
	return Company{
		ID:   result.ID,
		Name: result.Name,
	}
}
