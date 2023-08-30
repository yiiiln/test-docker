package http

import (
	"github.com/gin-gonic/gin"
	"main/domain/graphhelper/usecase"
	"main/model"
	"net/http"
)

type GraphHelper struct {
	GraphUsecase model.GraphUsecase
}

func NewGraphHandler(usecase *usecase.GraphHelper) *GraphHelper {
	return &GraphHelper{
		GraphUsecase: usecase,
	}
}

// DeviceCodeLogin godoc
//
//	@Summary		DeviceCodeLogin
//	@ID 			DeviceCodeLogin
//	@Description	登入
//	@Tags			Login With Device Code
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	"Get Access Token Successful"
//	@Failure		400	{string}	string	"ok"
//	@Router			/loginAzure/device [post]
func (g *GraphHelper) DeviceCodeLogin(c *gin.Context) {
	g.GraphUsecase.InitializeGraphWithDevice()
	accessToken, err := g.GraphUsecase.DisplayAccessTokenWithDevice()
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

// AccountPasswordLogin godoc
//
//	@Summary		AccountPasswordLogin
//	@ID 			AccountPasswordLogin
//	@Description	使用帳號、密碼登入 Azure
//	@Tags			Login With Password
//	@Accept			json
//	@Produce		json
//	@Param			account-password		formData	UserRequest		true	"Request-Login"
//	@Success		200	{string}	string	"Get Access Token Successful"
//	@Failure		400	{string}	string	"ok"
//	@Router			/loginAzure/password [post]
func (g *GraphHelper) AccountPasswordLogin(ctx *gin.Context) {
	body := UserRequest{}
	body.Account = ctx.PostForm("account")
	body.Password = ctx.PostForm("password")
	if err := ctx.ShouldBind(&body); err != nil { //BindJSON: 綁定提交的 JSON 參數信息
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	g.GraphUsecase.InitializeGraphWithPassword(body.Account, body.Password)
	accessToken, err := g.GraphUsecase.DisplayAccessTokenWithPassword()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"access_ms_graph_token_error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"access_ms_graph_token": accessToken,
	})
	return
}

type UserRequest struct {
	Account  string `json:"account"  form:"account" binding:"required"`  // 帳號
	Password string `json:"password" form:"password" binding:"required"` // 密碼
}
