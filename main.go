package main

import (
	"bytes"
	"chivato/configure"
	"chivato/message"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Create a struct that mimics the webhook response body
// https://core.telegram.org/bots/api#update
type webhookReqBody struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

// Handler is called everytime telegram sends us a webhook event
func Handler(res http.ResponseWriter, req *http.Request) {
	// First, decode the JSON response body
	body := &webhookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		fmt.Println("could not decode request body", err)
		return
	}

	// Check if the message contains the word "marco"
	// if not, return without doing anything
	if !strings.Contains(strings.ToLower(body.Message.Text), "borja") {
		return
	}

	// If the text contains marco, call the `sayPolo` function, which
	// is defined below
	if err := sayPolo(body.Message.Chat.ID); err != nil {
		fmt.Println("error in sending reply:", err)
		return
	}

	// log a confirmation message if the message is sent successfully
	fmt.Println("reply sent")
}

//The below code deals with the process of sending a response message
// to the user

// Create a struct to conform to the JSON body
// of the send message request
// https://core.telegram.org/bots/api#sendmessage
type sendMessageReqBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

// sayPolo takes a chatID and sends "polo" to them
func sayPolo(chatID int64) error {
	// Create the request body struct
	reqBody := &sendMessageReqBody{
		ChatID: chatID,
		Text:   "el rey del odoo https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Goatse_Security_Logo.png/220px-Goatse_Security_Logo.png",
	}
	// Create the JSON body from the struct
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	// Send a post request with your token
	res, err := http.Post("https://api.telegram.org/bot"+config.Apikey+"/sendMessage", "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	return nil
}

var config configure.Config

// FInally, the main funtion starts our server on port 3000
func main() {
	// TODO remove global variable
	//	var config Config
	config.ReadConfig()
	fmt.Println("---" + config.File + " ---")
	//fmt.Println("https://api.telegram.org/bot" + config.Apikey + "/sendMessage")

	if config.WebhookEnable == true {
		println("Activado: " + strconv.FormatBool(message.ActivateWebhook(config.Apikey, config.Webhookurl)))
		if config.SslEnable == false {
			fmt.Println("http")
			http.ListenAndServe(":"+config.Port, http.HandlerFunc(Handler))
		} else {
			fmt.Println("https")
			log.Fatal(http.ListenAndServeTLS(":"+config.Port, config.Sslcrt, config.Sslkey, http.HandlerFunc(Handler)))
		}
	} else {

		/*
			println("voy yo a por los datos https://api.telegram.org/bot" + config.Apikey + "/getWebhookInfo")
			// https://api.telegram.org/bot<config.Apikey/getWebhookInfo
			resp, err := http.Get("https://api.telegram.org/bot" + config.Apikey + "/getWebhookInfo")
			println(err)
			println(resp)
		*/
		message.GetWebhookInfo(config.Apikey)
		println(message.DeleteWebhook(config.Apikey))
		println(message.GetUpdates(config.Apikey))
		message.GetWebhookInfo(config.Apikey)
	}
}
