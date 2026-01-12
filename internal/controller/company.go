package controller

import (
	"golang/internal/service"
	"golang/pkg/response"

	"github.com/gin-gonic/gin"
)

type CompanyController struct {
	companyService *service.CompanyService
}

func NewCompanyController() *CompanyController {
	return &CompanyController{
		companyService: service.NewCompanyService(),
	}
}

func (cc *CompanyController) GetCompanyByID(c *gin.Context) {
	result := cc.companyService.GetCompanyByID()

	response.SuccessResponse(c, response.Success, result)
}
