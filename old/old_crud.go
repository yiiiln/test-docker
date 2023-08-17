package old

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Test godoc
//
//	@Summary		Connect Test
//	@Description	測試在 GCP Cloud Run 是否連得上不帶 pSQL 的映像檔
//	@Tags			OLD
//	@Accept			json
//	@Produce		application/json
//	@Success		200	{string}	string	"Connect Cloud Run Successful"
//	@Failure		404	{string}	string	"ok"
//	@Router			/test_connect [get]
func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "Test Connect Cloud Run",
		"Status":  "successful",
	})
	return
}

// Register godoc
//
//	@Summary		Register
//	@Description	註冊
//	@Tags			OLD
//	@Accept			json
//	@Produce		json
//	@Param			register_body	body	Customer		true	"Register account"
//	@Success		201	{string}	string	"Register Successful"
//	@Failure		400	{string}	string	"ok"
//	@Failure		404	{string}	string	"ok"
//	@Router			/c [post]
func Register(c *gin.Context) {
	var inputUser Customer
	if err := c.ShouldBindJSON(&inputUser); err != nil { //BindJSON: 綁定提交的 JSON 參數信息
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	user := DbUser{UserAccount: inputUser.CustomerAccount, UserPassword: inputUser.CustomerPassword}
	saveUser, err := user.Create() // 若帳號不存在，在資料庫裡新增它
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.AbortWithStatusJSON(http.StatusCreated, gin.H{"user": saveUser})
	return
}

// Login godoc
//
//	@Summary		Login
//	@Description	登入
//	@Tags			OLD
//	@Accept			json
//	@Produce		json
//	@Param			login_body	body	Customer		true	"Login account"
//	@Success		200	{string}	string	"Register Successful"
//	@Failure		400	{string}	string	"ok"
//	@Failure		404	{string}	string	"ok"
//	@Router			/login [post]
func Login(c *gin.Context) {
	var login Customer
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := DbUser{UserAccount: login.CustomerAccount, UserPassword: login.CustomerPassword}
	loginUser, err := user.ReadUserByAccount()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, loginUser)
	return
}

// GetUsers godoc
//
//	@Summary		GetUsers
//	@Description	列出所有
//	@Tags			OLD
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	"Get Users Successful"
//	@Failure		404	{string}	string	"ok"
//	@Router			/r [get]
func GetUsers(c *gin.Context) {

	var users []DbUser

	if result := Database.Debug().Find(&users); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &users)
	return
}

// UpdateAccountByID godoc
//
//	@Summary		Update Account By ID
//	@Description	更新Account
//	@Tags			OLD
//	@Accept			json
//	@Produce		json
//	@Param			ID		path	int		true	"Account ID"
//	@Param			update_account	body	UpdateAccount		true	"Account Struct"
//	@Success		200	{string}	string	"Update User Successful"
//	@Failure		404	{string}	string	"ok"
//	@Router			/ua/{ID} [put]
func UpdateAccountByID(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	accountId := c.Param("ID")
	intID, _ := strconv.Atoi(accountId)

	var user UpdateAccount
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updateData := UpdateUserStruct{ID: uint(intID), NewAccount: user.NewAccount}
	db, err := updateData.UpdateUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // 返回更新錯誤
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, db)
	return
}

// UpdatePasswordByID godoc
//
//	@Summary		Update Password By ID
//	@Description	更新Password
//	@Tags			OLD
//	@Accept			json
//	@Produce		json
//	@Param			ID		path	int		true	"Password ID"
//	@Param			update_password	body	UpdatePassword		true	"Password Struct"
//	@Success		200	{string}	string	"Update User Successful"
//	@Failure		404	{string}	string	"ok"
//	@Router			/up/{ID} [put]
func UpdatePasswordByID(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	accountId := c.Param("ID")
	intID, _ := strconv.Atoi(accountId)

	var user UpdatePassword
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updateData := UpdateUserStruct{ID: uint(intID), NewPassword: user.NewPassword}
	db, err := updateData.UpdateUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // 返回更新錯誤
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, db)
	return
}

// DeleteUserByAccount godoc
//
//	@Summary		DeleteUserByAccount
//	@Description	刪除使用者
//	@Tags			OLD
//	@Accept			json
//	@Produce		json
//	@Param			delete_body		body	Customer		true	"Delete user"
//	@Success		200	{string}	string	"Delete User Successful"
//	@Failure		404	{string}	string	"ok"
//	@Router			/d [delete]
func DeleteUserByAccount(c *gin.Context) {
	var user Customer
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userFromDB := DbUser{UserAccount: user.CustomerAccount, UserPassword: user.CustomerPassword}
	_, readErr := userFromDB.ReadUserByAccount()
	if readErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": readErr.Error()})
		return
	}
	_, deleteErr := userFromDB.DeleteUser()
	if deleteErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": deleteErr.Error()})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "刪除成功"})
	return
}
