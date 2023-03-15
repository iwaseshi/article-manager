package client

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// https://qiita.com/api/v2/docs#post-apiv2items

func SendQiita() {
	accessToken := "ないしょ"
	url := "https://qiita.com/api/v2/items"

	title := "てす"
	body := "Hello, World!"
	tags := []map[string]string{{"name": "Go"}}

	reqBody, _ := json.Marshal(map[string]interface{}{
		"title": title,
		"body":  body,
		"tags":  tags,
	})

	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+accessToken).
		SetBody(reqBody).
		Post(url)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(resp.Status())
	fmt.Println(resp.Status())
}
