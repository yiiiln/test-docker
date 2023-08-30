package repository

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	auth "github.com/microsoft/kiota-authentication-azure-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"main/model"
	"os"
	"strings"
)

type GraphHelper struct {
	model.GraphStruct
}

func NewGraphHelper() *GraphHelper {
	return &GraphHelper{}
}

func (g *GraphHelper) InitializeGraphWithPasswordAuth(username string, password string) error {
	//TODO implement me
	clientID := os.Getenv("CLIENT_ID")
	tenantID := os.Getenv("TENANT_ID")
	scopes := os.Getenv("GRAPH_USER_SCOPES")
	g.PasswordScopes = strings.Split(scopes, ",")

	credential, err := azidentity.NewUsernamePasswordCredential(tenantID, clientID, username, password, nil)
	if err != nil {
		return err
	}
	g.UsernamePasswordCredential = credential

	// Create an auth provider using the credential
	authProvider, err := auth.NewAzureIdentityAuthenticationProvider(credential)
	if err != nil {
		return err
	}

	// Create a request adapter using the auth provider
	adapter, err := msgraphsdk.NewGraphRequestAdapter(authProvider)
	if err != nil {
		return err
	}

	// Create a Graph client using request adapter
	client := msgraphsdk.NewGraphServiceClient(adapter)
	g.PasswordClient = client

	return nil
}

func (g *GraphHelper) InitializeGraphWithDeviceAuth() error {
	//TODO implement me
	clientId := os.Getenv("CLIENT_ID")
	tenantId := os.Getenv("TENANT_ID")
	scopes := os.Getenv("GRAPH_USER_SCOPES")
	g.DeviceScopes = strings.Split(scopes, ",")

	// Create the device code credential
	credential, err := azidentity.NewDeviceCodeCredential(&azidentity.DeviceCodeCredentialOptions{
		ClientID: clientId,
		TenantID: tenantId,
		UserPrompt: func(ctx context.Context, message azidentity.DeviceCodeMessage) error {
			fmt.Println(message.Message)
			return nil
		},
	})
	if err != nil {
		return err
	}

	g.DeviceCodeCredential = credential

	// Create an auth provider using the credential
	authProvider, err := auth.NewAzureIdentityAuthenticationProviderWithScopes(credential, g.DeviceScopes)
	if err != nil {
		return err
	}

	// Create a request adapter using the auth provider
	adapter, err := msgraphsdk.NewGraphRequestAdapter(authProvider)
	if err != nil {
		return err
	}

	// Create a Graph client using request adapter
	client := msgraphsdk.NewGraphServiceClient(adapter)
	g.DeviceClient = client

	return nil
}

func (g *GraphHelper) GetUserTokenWithPassword() (*string, error) {
	//TODO implement me
	token, err := g.UsernamePasswordCredential.GetToken(context.Background(), policy.TokenRequestOptions{
		Scopes: g.PasswordScopes,
	})
	if err != nil {
		return nil, err
	}

	return &token.Token, nil
}

func (g *GraphHelper) GetUserTokenWithDevice() (*string, error) {
	token, err := g.DeviceCodeCredential.GetToken(context.Background(), policy.TokenRequestOptions{
		Scopes: g.DeviceScopes,
	})
	if err != nil {
		return nil, err
	}

	return &token.Token, nil
}
