package usecase

import (
	"log"
	"main/domain/graphhelper/repository"
	"main/model"
)

type GraphHelper struct {
	GraphRepository model.GraphRepository
}

func NewGraphHelper(usecase *repository.GraphHelper) *GraphHelper {
	return &GraphHelper{
		GraphRepository: usecase,
	}
}

func (g *GraphHelper) InitializeGraphWithPassword(username string, password string) {
	//TODO implement me
	err := g.GraphRepository.InitializeGraphWithPasswordAuth(username, password)
	if err != nil {
		log.Panicf("Error initializing graph with password auth: %v\n", err)
	}
}

func (g *GraphHelper) InitializeGraphWithDevice() {
	//TODO implement me
	err := g.GraphRepository.InitializeGraphWithDeviceAuth()
	if err != nil {
		log.Panicf("Error initializing Graph with device auth: %v\n", err)
	}
}

func (g *GraphHelper) DisplayAccessTokenWithPassword() (*string, error) {
	//TODO implement me
	token, err := g.GraphRepository.GetUserTokenWithPassword()
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (g *GraphHelper) DisplayAccessTokenWithDevice() (*string, error) {
	//TODO implement me
	token, err := g.GraphRepository.GetUserTokenWithDevice()
	if err != nil {
		return nil, err
	}
	return token, nil
}
