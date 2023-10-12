package apis

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func GetAnything() {

	type Headers struct {
		Accept         string `json:"Accept"`
		AcceptEncoding string `json:"Accept-Encoding"`
		ContentType    string `json:"Content-Type"`
		Host           string `json:"Host"`
		UserAgent      string `json:"User-Agent"`
		XAmznTraceId   string `json:"X-Amzn-Trace-Id"`
	}

	type BodyResponse struct {
		Json    string  `json:"json"`
		Data    string  `json:"data"`
		Origin  string  `json:"origin"`
		Method  string  `json:"method"`
		Url     string  `json:"url"`
		Headers Headers `json:"headers"`
	}

	var sb strings.Builder

	url := "https://httpbin.org/anything/"

	sb.WriteString(url)
	sb.WriteString("hello")

	req, err := http.NewRequest(http.MethodGet, sb.String(), nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Accept", "*/*")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "go-tutorial")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}
	// body, _ := io.ReadAll(res.Body)

	// fmt.Println(res)
	// fmt.Println(string(body))

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	result := BodyResponse{}
	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		log.Fatalf("unable to parse value: %q, error: %s", string(body), jsonErr.Error())
	}
	fmt.Println(result)
}
