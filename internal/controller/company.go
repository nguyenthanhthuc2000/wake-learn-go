package controller

import (
	"net/http"

	"golang/internal/service"

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

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success",
		"data":    result,
	})
}
