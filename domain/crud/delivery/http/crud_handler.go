package http

import (
	"github.com/gin-gonic/gin"
	"main/model"
	"net/http"
)

type Handler struct {
	UserUsecase model.UserUsecase
}

func NewUserHandler(usecase model.UserUsecase) *Handler {
	return &Handler{
		UserUsecase: usecase,
	}
}

// Register godoc
//
//	@Summary		Register
//	@Description	註冊
//	@Tags			CRUD
//	@Accept			json
//	@Produce		json
//	@Param			register_body	body	UserRequest		true	"Register account"
//	@Success		201	{string}	string	"Register Successful"
//	@Failure		400	{string}	string	"ok"
//	@Failure		404	{string}	string	"ok"
//	@Router			/CRUD [post]
func (h *Handler) Register(ctx *gin.Context) {
	body := UserRequest{}
	if err := ctx.BindJSON(&body); err != nil { //BindJSON: 綁定提交的 JSON 參數信息
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	user := LoadUserFromRequest(&body)
	if err := h.UserUsecase.Create(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, body)
}

// Login godoc
//
//	@Summary		Login
//	@Description	登入
//	@Tags			CRUD
//	@Accept			json
//	@Produce		json
//	@Param			login_body	body	UserRequest		true	"Login account"
//	@Success		200	{string}	string	"Register Successful"
//	@Failure		400	{string}	string	"ok"
//	@Failure		404	{string}	string	"ok"
//	@Router			/CRUD/login [post]
func (h *Handler) Login(ctx *gin.Context) {
	body := UserRequest{}
	if err := ctx.BindJSON(&body); err != nil { //BindJSON: 綁定提交的 JSON 參數信息
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	user := LoadUserFromRequest(&body)
	LoginUser, err := h.UserUsecase.Login(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"login status": "success",
		"user":         LoginUser,
	})
}

// GetUser godoc
//
//	@Summary		GetUser
//	@Description	列出欲查詢的使用者
//	@Tags			CRUD
//	@Accept			json
//	@Produce		json
//	@Param			id		path	int		true	"Account ID"
//	@Success		200	{string}	string	"Get User Successful"
//	@Failure		400	{string}	string	"ok"
//	@Router			/CRUD/{id} [get]
func (h *Handler) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := h.UserUsecase.Get(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	aluminumResponse := LoadResponseFromUser(user)
	ctx.JSON(http.StatusOK, aluminumResponse)
}

// GetAllUser godoc
//
//	@Summary		GetAllUser
//	@Description	列出所有
//	@Tags			CRUD
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	"Get All Users Successful"
//	@Failure		400	{string}	string	"ok"
//	@Router			/CRUD [get]
func (h *Handler) GetAllUser(ctx *gin.Context) {
	users, err := h.UserUsecase.GetAll()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	var userResponseSlice []UserResponse
	for _, element := range users {
		crudResponse := LoadResponseFromUser(element)
		userResponseSlice = append(userResponseSlice, crudResponse)
	}
	ctx.JSON(http.StatusOK, userResponseSlice)
}

type UserRequest struct {
	Account  string `json:"account"`  // 帳號
	Password string `json:"password"` // 密碼
}

func LoadUserFromRequest(reqBody *UserRequest) *model.User {
	user := model.User{
		Account:  reqBody.Account,
		Password: reqBody.Password,
	}
	return &user
}

type UserResponse struct {
	ID       string `json:"id"`       // 使用者ID
	Account  string `json:"account"`  // 帳號
	Password string `json:"password"` // 密碼
}

func LoadResponseFromUser(user model.User) UserResponse {
	crudResponse := UserResponse{
		ID:       user.ID,
		Account:  user.Account,
		Password: user.Password,
	}
	return crudResponse
}

// UpdateAccountByID godoc
//
//	@Summary		Update Account By ID
//	@Description	更新使用者 Account
//	@Tags			CRUD
//	@Accept			json
//	@Produce		json
//	@Param			id		path	int		true	"Account ID"
//	@Param			update_account	body	UserRequestUpdateAccount		true	"Account Struct"
//	@Success		200	{string}	string	"Update User Successful"
//	@Failure		404	{string}	string	"ok"
//	@Router			/CRUD/uA/{id} [put]
func (h *Handler) UpdateAccountByID(ctx *gin.Context) {
	id := ctx.Param("id")
	body := UserRequestUpdateAccount{}
	if err := ctx.BindJSON(&body); err != nil { //BindJSON: 綁定提交的 JSON 參數信息
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	updateUser := LoadUserFromUpdateAccountRequest(&body)

	user, err := h.UserUsecase.UpdateAccount(id, updateUser.Account)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	userResponse := LoadUpdateAccountResponseFromUser(user)

	ctx.JSON(http.StatusOK, gin.H{"update": userResponse})
}

type UserRequestUpdateAccount struct {
	NewAccount string `json:"new_account"` // 更新帳號
}

func LoadUserFromUpdateAccountRequest(reqBody *UserRequestUpdateAccount) *model.User {
	user := model.User{
		Account: reqBody.NewAccount,
	}
	return &user
}

type UserUpdateAccountResponse struct {
	NewAccount string `json:"new_account"` // 更新帳號
}

func LoadUpdateAccountResponseFromUser(user model.User) UserUpdateAccountResponse {
	crudResponse := UserUpdateAccountResponse{
		NewAccount: user.Account,
	}
	return crudResponse
}

// UpdatePasswordByID godoc
//
//	@Summary		Update Password By ID
//	@Description	更新使用者 Password
//	@Tags			CRUD
//	@Accept			json
//	@Produce		json
//	@Param			id		path	int		true	"Password ID"
//	@Param			update_password	body	UserRequestUpdatePassword		true	"Password Struct"
//	@Success		200	{string}	string	"Update User Successful"
//	@Failure		404	{string}	string	"ok"
//	@Router			/CRUD/uP/{id} [put]
func (h *Handler) UpdatePasswordByID(ctx *gin.Context) {
	id := ctx.Param("id")
	body := UserRequestUpdatePassword{}
	if err := ctx.BindJSON(&body); err != nil { //BindJSON: 綁定提交的 JSON 參數信息
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	updateUser := LoadUserFromUpdatePasswordRequest(&body)

	user, err := h.UserUsecase.UpdatePassword(id, updateUser.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	userResponse := LoadUpdatePasswordResponseFromUser(user)

	ctx.JSON(http.StatusOK, gin.H{"update": userResponse})
}

type UserRequestUpdatePassword struct {
	NewPassword string `json:"new_password"` // 更新密碼
}

func LoadUserFromUpdatePasswordRequest(reqBody *UserRequestUpdatePassword) *model.User {
	user := model.User{
		Password: reqBody.NewPassword,
	}
	return &user
}

type UserUpdatePasswordResponse struct {
	NewPassword string `json:"new_password"` // 更新密碼
}

func LoadUpdatePasswordResponseFromUser(user model.User) UserUpdatePasswordResponse {
	crudResponse := UserUpdatePasswordResponse{
		NewPassword: user.Password,
	}
	return crudResponse
}

// DeleteUserByID godoc
//
//	@Summary		DeleteUserByID
//	@Description	刪除使用者
//	@Tags			CRUD
//	@Accept			json
//	@Produce		json
//	@Param			id		path	int		true	"Delete Account ID"
//	@Success		200	{string}	string	"Delete User Successful"
//	@Failure		400	{string}	string	"ok"
//	@Router			/CRUD/{id} [delete]
func (h *Handler) DeleteUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.UserUsecase.DeleteBy(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"delete status": "success"})
}
