package bootstrap

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type URLRules struct {
	Layout  string   `json:"layout"`
	Routers []Router `json:"routers"`
}

type Router struct {
	URL      string `json:"url"`
	Template string `json:"template"`
	Layout   string `json:"lauout"`
}

var URLRule *URLRules
var I18N map[string]map[string]string

func init() {
	// init url routers
	file, err := os.Open("conf/url.conf")
	if err != nil {
		panic("open url.conf err :" + err.Error())
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic("read url.conf err :" + err.Error())
	}
	URLRule = &URLRules{}
	err = json.Unmarshal(data, URLRule)
	if err != nil {
		panic("json.Unmarshal url.conf err :" + err.Error())
	}

	// init i18n config
	I18N = make(map[string]map[string]string)
	file, err = os.Open("conf/i18n.conf")
	if err != nil {
		panic("open i18n.conf err :" + err.Error())
	}
	data, err = ioutil.ReadAll(file)
	if err != nil {
		panic("read i18n.conf err :" + err.Error())
	}
	err = json.Unmarshal(data, &I18N)
	if err != nil {
		panic("json.Unmarshal i18n.conf err :" + err.Error())
	}
}
