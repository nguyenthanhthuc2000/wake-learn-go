package repo

type CompanyRepository struct{}

type Company struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewCompanyRepo() *CompanyRepository {
	return &CompanyRepository{}
}

func (cr *CompanyRepository) GetCompanyByID() Company {
	return Company{
		ID:   "1",
		Name: "THK Holding Vietnam",
	}
}
