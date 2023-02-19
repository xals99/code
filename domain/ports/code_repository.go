package ports

import "code/domain/model"

type CodeRepository interface {
	Get(code string) (*model.Code, error)
	Set(model.Code)
}
