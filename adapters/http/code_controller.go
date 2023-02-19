package http

import (
	"code/domain/ports"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UseCodeRequest struct {
	Code string `json:"code"`
}

type UseCodeResponse struct {
	Title string `json:"title"`
	Type  string `json:"type"`
}

type CodeController struct {
	codeService ports.CodeService
}

func NewCodeController(codeService ports.CodeService) CodeController {
	return CodeController{codeService}
}

func (cc CodeController) UseCode(c *gin.Context) {
	var request UseCodeRequest

	if err := c.BindJSON(&request); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Use code: ", request.Code)

	product, err := cc.codeService.UseCode(request.Code)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.IndentedJSON(http.StatusOK, UseCodeResponse{product.Name, product.Type.String()})
}
