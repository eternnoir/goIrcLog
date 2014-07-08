package Config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	DbType, DbConnection string
	Channels             string
	User                 string
	Password             string
	Nick                 string
	Server               string
	Port                 string
}

type Configslice struct {
	Configs []Config
}

func GetConfigslice(path string) *Configslice {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
	}
	str := string(dat)
	var configs Configslice
	json.Unmarshal([]byte(str), &configs)
	return &configs
}
