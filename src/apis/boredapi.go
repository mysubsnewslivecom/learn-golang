package apis

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func BoredApi() {

	url := "http://www.boredapi.com/api/activity?participants=1&type=cooking"
	type BodyResponse struct {
		Activity      string  `json:"activity"`
		Type          string  `json:"type"`
		Participants  int64   `json:"participants"`
		Price         float64 `json:"price"`
		Link          string  `json:"link"`
		Key           string  `json:"key"`
		Accessibility float64 `json:"accessibility"`
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Accept", "*/*")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "spacecount-tutorial")

	res, getErr := http.DefaultClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
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
