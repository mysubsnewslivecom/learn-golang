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
		Host           string `json:"Host"`
		UserAgent      string `json:"User-Agent"`
		XAmznTraceID   string `json:"X-Amzn-Trace-Id"`
	}

	type Args struct {
	}

	type BodyResponse struct {
		Args    Args        `json:"args"`
		Data    string      `json:"data"`
		Files   Args        `json:"files"`
		Form    Args        `json:"form"`
		Headers Headers     `json:"headers"`
		JSON    interface{} `json:"json"`
		Method  string      `json:"method"`
		Origin  string      `json:"origin"`
		URL     string      `json:"url"`
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
