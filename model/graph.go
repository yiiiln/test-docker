package model

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
)

type GraphStruct struct {
	GraphDevice
	GraphPassword
}

type GraphDevice struct {
	DeviceCodeCredential *azidentity.DeviceCodeCredential //它會自動打開瀏覽器訪問登錄頁面，讓用戶瀏覽到 Azure Active Directory URL、輸入代碼並進行身份驗證。
	DeviceClient         *msgraphsdk.GraphServiceClient
	DeviceScopes         []string
}

type GraphPassword struct {
	UsernamePasswordCredential *azidentity.UsernamePasswordCredential
	PasswordClient             *msgraphsdk.GraphServiceClient
	Username                   string
	Password                   string
	PasswordScopes             []string
}

type GraphRepository interface {
	InitializeGraphWithDeviceAuth() error
	InitializeGraphWithPasswordAuth(username string, password string) error
	GetUserTokenWithDevice() (*string, error)
	GetUserTokenWithPassword() (*string, error)
}

type GraphUsecase interface {
	InitializeGraphWithDevice()
	InitializeGraphWithPassword(username string, password string)
	DisplayAccessTokenWithDevice() (*string, error)
	DisplayAccessTokenWithPassword() (*string, error)
}
