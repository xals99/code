package ports

import "github.com/xals99/code/domain/model"

type CodeService struct {
	codeRepository CodeRepository
}

func NewCodeService(codeRepository CodeRepository) CodeService {
	return CodeService{codeRepository}
}

func (cs *CodeService) UseCode(code string) (*model.Product, error) {
	codeDb, err := cs.codeRepository.Get(code)

	if err != nil {
		return nil, err
	}

	return codeDb.Product, nil
}
