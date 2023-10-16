// https://support.hashicorp.com/hc/en-us/articles/18221966166291-HCP-Vault-Secrets-get-create-and-delete-secrets-via-API
package vault

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func GetClientToken(hcpClientId string, hcpClientSecret string) (string, error) {

	type RespBody struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int64  `json:"expires_in"`
		TokenType   string `json:"token_type"`
	}

	var url = "https://auth.hashicorp.com/oauth/token"

	payload := fmt.Sprint("{\"audience\": \"https://api.hashicorp.cloud\", \"grant_type\": \"client_credentials\", \"client_id\": \"", hcpClientId, "\", \"client_secret\": \"", hcpClientSecret, "\"}")

	jsonBody := []byte(payload)
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)

	body := RespBody{}

	_ = json.Unmarshal(data, &body)
	return string(body.AccessToken), nil
}

func GetSecret() {

	type CreatedBy struct {
		Name  string `json:"name"`
		Type  string `json:"type"`
		Email string `json:"email"`
	}

	type SyncStatus struct {
	}

	type Version struct {
		Version   string    `json:"version"`
		Type      string    `json:"type"`
		CreatedAt string    `json:"created_at"`
		Value     string    `json:"value"`
		CreatedBy CreatedBy `json:"created_by"`
	}

	type Secret struct {
		Name          string     `json:"name"`
		Version       Version    `json:"version"`
		CreatedAt     string     `json:"created_at"`
		LatestVersion string     `json:"latest_version"`
		CreatedBy     CreatedBy  `json:"created_by"`
		SyncStatus    SyncStatus `json:"sync_status"`
	}

	type SecretMain struct {
		Secrets []Secret `json:"secrets"`
	}

	var hpcOrgID, hcpProjID, hcpClientId, hcpClientSecret = os.Getenv("HCP_ORG_ID"), os.Getenv("HCP_PROJ_ID"), os.Getenv("HCP_CLIENT_ID"), os.Getenv("HCP_CLIENT_SECRET")

	hcpAPIToken, err := GetClientToken(hcpClientId, hcpClientSecret)
	if err != nil {
		panic(err)
	}

	log.Printf("|- \nHCP_ORG_ID: %v\nHCP_PROJ_ID: %v", hpcOrgID, hcpProjID)
	log.Printf("|- \nHCP_CLIENT_ID: %v\nHCP_CLIENT_SECRET: %v", hcpClientId, hcpClientSecret)

	baseUrl := "https://api.cloud.hashicorp.com/secrets/2023-06-13/organizations"
	secret := "test-app"

	url, err := url.JoinPath(baseUrl, hpcOrgID, "projects", hcpProjID, "apps", secret, "open")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("URL: %v", url)

	var secretStruct = SecretMain{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", hcpAPIToken))

	res, _ := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)

	_ = json.Unmarshal(data, &secretStruct)

	log.Println(secretStruct.Secrets[0].Name)
	log.Println(secretStruct.Secrets[0].Version.Value)

}

// {
// 	"secrets": [
// 	  {
// 		"name": "postgres",
// 		"version": {
// 		  "version": "1",
// 		  "type": "kv",
// 		  "created_at": "2023-10-16T16:28:43.644793Z",
// 		  "value": "secret1234",
// 		  "created_by": {
// 			"name": "my.subs.news@live.com",
// 			"type": "TYPE_USER",
// 			"email": "my.subs.news@live.com"
// 		  }
// 		},
// 		"created_at": "2023-10-16T16:28:43.644793Z",
// 		"latest_version": "1",
// 		"created_by": {
// 		  "name": "my.subs.news@live.com",
// 		  "type": "TYPE_USER",
// 		  "email": "my.subs.news@live.com"
// 		},
// 		"sync_status": {}
// 	  }
// 	]
// }

// curl \
// --location "https://api.cloud.hashicorp.com/secrets/2023-06-13/organizations/8d155e8d-4f55-4596-8bec-422f5d18c7d7/projects/4f479c4e-eac1-43f9-8e0b-bd9c8f2d2d9e/apps/test-app/open" \
// --request GET \
// --header "Authorization: Bearer $HCP_API_TOKEN" | jq

// curl \
// --request DELETE \
// --location "https://api.cloud.hashicorp.com/secrets/2023-06-13/organizations/$HCP_ORG_ID/projects/$HCP_PROJ_ID/apps/$VLT_APPS_NAME/secrets/my_secret_name" \
// --header "Authorization: Bearer $HCP_API_TOKEN"

// curl \
// --request POST \
// --location "https://api.cloud.hashicorp.com/secrets/2023-06-13/organizations/$HCP_ORG_ID/projects/$HCP_PROJ_ID/apps/$VLT_APPS_NAME/kv" \
// --header "Authorization: Bearer $HCP_API_TOKEN" \
// --header "Content-Type: application/json" \
// --data-raw '{
// "name": "my_secret_name",
// "value": "my_secret_value"
// }'
