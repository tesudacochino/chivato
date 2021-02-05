package configure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config configuracion
type Config struct {
	File          string
	Apikey        string `json:"api"`
	SslEnable     bool   `json:"ssl"`
	Sslcrt        string `json:"sslcrt"`
	Sslkey        string `json:"sslkey"`
	WebhookEnable bool   `json:"webhook"`
	Webhookurl    string `json:"webhookurl"`
}

// GetFileConfig get file config
func (config *Config) GetFileConfig() {
	config.File = "conf/development.conf"
	if _, err := os.Stat(config.File); os.IsNotExist(err) {
		config.File = "conf/default.conf"
	}

}

// ReadConfig read the config
func (config *Config) ReadConfig() {
	config.GetFileConfig()

	jsonFile, err := os.Open(config.File)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Open OK")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &config)

	if config.Apikey == "" {
		fmt.Println("Error: undefined Api in " + config.File)
		os.Exit(1)
	}
	if (config.WebhookEnable == true) && (config.Webhookurl == "") {
		fmt.Printf("Error: undefined WebHook url in" + config.File)
		os.Exit(1)
	}

	/*  acceso directo
		var result map[string]interface{}
		json.Unmarshal([]byte(byteValue), &result)

	    fmt.Println(result["api"])
	*/

}
