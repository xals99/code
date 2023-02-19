package main

import (
	"code/adapters/db"
	"code/adapters/http"
	"code/domain/model"
	"code/domain/ports"

	"github.com/gin-gonic/gin"
)

func main() {
	code_repository := db.NewInMemoryCodeRepository()
	product := model.Product{Name: "Hans Kloss", Type: model.Series}
	code_repository.Set(model.Code{Code: "xxx", Product: &product, User: nil})
	code_repository.Set(model.Code{Code: "yyy", Product: &product, User: nil})
	code_repository.Set(model.Code{Code: "zzz", Product: &product, User: nil})
	code_service := ports.NewCodeService(code_repository)
	code_controller := http.NewCodeController(code_service)

	router := gin.Default()
	router.POST("/use", code_controller.UseCode)

	router.Run("localhost:8080")
}
