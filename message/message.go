package message

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type webhookInfopMessage struct {
	Ok     bool `json:"ok"`
	Result struct {
		URL                  string `json:"url"`
		HasCustomCertificate bool   `json:"has_custom_certificate"`
		PendingUpdateCount   int    `json:"pending_update_count"`
		LastErrorDate        int    `json:"last_error_date"`
		LastErrorMessage     string `json:"last_error_message"`
		MaxConnections       int    `json:"max_connections"`
		IPAddress            string `json:"ip_address"`
	} `json:"result"`
}

type deleteWebHookMessage struct {
	Ok          bool   `json:"ok"`
	Result      bool   `json:"result"`
	Description string `json:"description"`
}

type messageList struct {
	Ok     bool `json:"ok"`
	Result []struct {
		UpdateID int `json:"update_id"`
		Message  struct {
			MessageID int `json:"message_id"`
			From      struct {
				ID           int    `json:"id"`
				IsBot        bool   `json:"is_bot"`
				FirstName    string `json:"first_name"`
				LanguageCode string `json:"language_code"`
			} `json:"from"`
			Chat struct {
				ID        int    `json:"id"`
				FirstName string `json:"first_name"`
				Type      string `json:"type"`
			} `json:"chat"`
			Date int    `json:"date"`
			Text string `json:"text"`
		} `json:"message"`
	} `json:"result"`
}

func request(Apikey string, command string) []byte {
	var url = "https://api.telegram.org/bot" + Apikey + "/" + command

	spaceClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	return (body)

	/*jsonErr := json.Unmarshal(body, &answer)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}*/

}

func getWebhookInfo(Apikey string) {

	var answer webhookInfopMessage

	body := request(Apikey, "getWebhookInfo")

	jsonErr := json.Unmarshal(body, &answer)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	println(answer.Result.URL)

}

func DeleteWebhookInfo(Apikey string) bool {
	//var url = "https://api.telegram.org/bot" + Apikey + "/deleteWebhook"
	var answer deleteWebHookMessage

	body := request(Apikey, "deleteWebhook")

	jsonErr := json.Unmarshal(body, &answer)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	println(answer.Description)
	return answer.Ok

}

func GetUpdates(Apikey string) bool {
	var answer messageList

	body := request(Apikey, "getUpdates")

	jsonErr := json.Unmarshal(body, &answer)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	println(answer.Ok)
	return answer.Ok

}
