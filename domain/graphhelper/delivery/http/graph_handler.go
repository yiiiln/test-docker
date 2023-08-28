package http

import (
	"github.com/gin-gonic/gin"
	"main/model"
	"net/http"
)

type GraphHelper struct {
	GraphUsecase model.GraphUsecase
}

func NewGraphHandler(usecase model.GraphUsecase) *GraphHelper {
	return &GraphHelper{
		GraphUsecase: usecase,
	}
}

func (g *GraphHelper) ToDisplayAccessToken(c *gin.Context) {
	g.GraphUsecase.InitializeGraph()
	accessToken, err := g.GraphUsecase.DisplayAccessToken()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"access_ms_graph_token_error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"access_ms_graph_token": accessToken,
	})
	return
}
