package model

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
)

type GraphStruct struct {
	DeviceCodeCredential *azidentity.DeviceCodeCredential //它會自動打開瀏覽器訪問登錄頁面，讓用戶瀏覽到 Azure Active Directory URL、輸入代碼並進行身份驗證。
	UserClient           *msgraphsdk.GraphServiceClient
	GraphUserScopes      []string
}

type GraphRepository interface {
	InitializeGraphForUserAuth() error
	GetUserToken() (*string, error)
}
type GraphUsecase interface {
	InitializeGraph()
	DisplayAccessToken() (*string, error)
}
