package oauth2
//package main

import (
	/* Not used
	   "github.com/gruntwork-io/terratest/modules/azure"
	*/
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type IDENTITY struct {
	ClientId       string
	ClientSecret   string
	SubscriptionId string
	TenantId       string
	ApiEndpoint    string
	RedirectUri    string
	Resource       string
	GrantType      string
}

var Identity = IDENTITY {
	os.Getenv("TF2_CLIENT_ID"),
	os.Getenv("TF2_CLIENT_SECRET"),
	os.Getenv("TF2_SUBSCRIPTION_ID"),
	os.Getenv("TF2_TENANT_ID"),
	"https://login.microsoftonline.com/72f988bf-86f1-41af-91ab-2d7cd011db47/oauth2/token",
	"http://BrunoServicePrincipal",
	"https://management.azure.com",
	"client_credentials",
}

// =======================================================================
//
// =======================================================================
func GetToken() string {
	formData := url.Values{
		"grant_type":    {Identity.GrantType},
		"client_id":     {Identity.ClientId},
		"client_secret": {Identity.ClientSecret},
		"redirect_uri":  {Identity.RedirectUri},
		"resource":      {Identity.Resource},
	}
    /*
	fmt.Println("==============================")
	fmt.Println("formData", formData)
	fmt.Println("length of clientId ", len(Identity.ClientId))
	fmt.Println("==============================")
    */
	resp, err := http.PostForm(Identity.ApiEndpoint, formData)
	if err != nil {
		fmt.Printf("Error for http.PostForm() %s\n", err)
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result["access_token"].(string)
}

/*
func main() {
	adToken := getToken2()
	// Uses the Azure REST API
	fmt.Println("token = %s", adToken)}
*/
