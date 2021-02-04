package configure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config configuracion
type Config struct {
	file   string
	Apikey string `json:"api"`
	Ssl    bool   `json:"ssl"`
	Sslcrt string `json:"sslcrt"`
	Sslkey string `json:"sslkey"`
}

// GetFileConfig get file config
func (config *Config) GetFileConfig() {
	config.file = "conf/development.conf"
	if _, err := os.Stat("conf/development.conf"); os.IsNotExist(err) {
		config.file = "conf/default.conf"
	}

}

// ReadConfig read the config
func (config *Config) ReadConfig() {
	config.GetFileConfig()

	jsonFile, err := os.Open(config.file)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Open OK")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &config)

	fmt.Println(config.Apikey)

	/*  acceso directo
		var result map[string]interface{}
		json.Unmarshal([]byte(byteValue), &result)

	    fmt.Println(result["api"])
	*/

}
