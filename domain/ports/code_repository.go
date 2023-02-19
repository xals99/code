package ports

import "github.com/xals99/code/domain/model"

type CodeRepository interface {
	Get(code string) (*model.Code, error)
	Set(model.Code)
}
