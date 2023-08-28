package usecase

import (
	"main/model"
)

type GraphHelper struct {
	GraphRepository model.GraphRepository
}

func NewGraphHelper(usecase model.GraphRepository) *GraphHelper {
	return &GraphHelper{
		GraphRepository: usecase,
	}
}

func (g *GraphHelper) InitializeGraph() {
	//TODO implement me
	g.GraphRepository.InitializeGraphForUserAuth()
}

func (g *GraphHelper) DisplayAccessToken() (*string, error) {
	//TODO implement me
	token, err := g.GraphRepository.GetUserToken()
	if err != nil {
		return nil, err
	}
	return token, nil
}
