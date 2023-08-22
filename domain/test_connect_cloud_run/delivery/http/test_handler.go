package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// TestConnect godoc
//
//	@Summary		Test Connect Cloud Run
//	@ID 			Test Connect Cloud Run
//
// @Security		BearerOAuth2
//
//	@Description	測試在 GCP Cloud Run 是否連得上不帶 pSQL 的映像檔
//	@Tags			Test Connect
//	@Accept			json
//	@Produce		application/json
//	@Success		200	{string}	string	"Connect Cloud Run Successful"
//	@Router			/Test [post]
func TestConnect(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "Test Connect Cloud Run",
		"Status":  "successful",
	})
	return
}
