package message

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type messageWebhookInfo struct {
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

type messageDeleteWebHook struct {
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
			//		Chat chat   `json:"chat"`
			Date int    `json:"date"`
			Text string `json:"text"`
		} `json:"message"`
	} `json:"result"`
}

type chat struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	Type      string `json:"type"`
}

/*
type WebhookReqBody struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}*/

// WebhookReqBody message send by webhook
type WebhookReqBody struct {
	UpdateID int64 `json:"update_id"`
	Message  struct {
		MessageID int64 `json:"message_id"`
		From      struct {
			ID           int64  `json:"id"`
			IsBot        bool   `json:"is_bot"`
			FirstName    string `json:"first_name"`
			LanguageCode string `json:"language_code"`
		} `json:"from"`
		Chat chat   `json:"chat"`
		Date int    `json:"date"`
		Text string `json:"text"`
	} `json:"message"`
}

/*type webhookReqBody struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}*/

type messgeActivate struct {
	Ok          bool   `json:"ok"`
	Result      bool   `json:"result"`
	Description string `json:"description"`
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

	println(string(body))
	return (body)

	/*jsonErr := json.Unmarshal(body, &answer)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}*/

}

// GetWebhookInfo get status
func GetWebhookInfo(Apikey string) {

	var answer messageWebhookInfo

	body := request(Apikey, "getWebhookInfo")

	jsonErr := json.Unmarshal(body, &answer)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

}

// ActivateWebhook configure WebHook
func ActivateWebhook(Apikey string, domain string) bool {
	// https://api.telegram.org/bot[TU_TOKEN]/setWebhook?url=https://[TU_DOMINIO]/[CAMINO_AL_WEBHOOK]
	var answer messgeActivate

	body := request(Apikey, "setWebhook?url="+domain)

	jsonErr := json.Unmarshal(body, &answer)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return answer.Ok
}

// DeleteWebHook remove webhook
func DeleteWebHook(Apikey string) bool {
	//var url = "https://api.telegram.org/bot" + Apikey + "/deleteWebhook"
	var answer messageDeleteWebHook

	body := request(Apikey, "deleteWebhook")

	jsonErr := json.Unmarshal(body, &answer)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return answer.Ok
}

// GetUpdates get messages no webhook mode
func GetUpdates(Apikey string) bool {
	var answer messageList

	body := request(Apikey, "getUpdates")

	jsonErr := json.Unmarshal(body, &answer)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return answer.Ok

}
