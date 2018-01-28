// Package config is golang package, supply config file resolve.
//
// https://github.com/daewood
package config

import (
	"errors"
	"io/ioutil"
	"path"
	"strings"

	"github.com/daewood/config/ini"
	"github.com/daewood/config/json"
	"github.com/daewood/config/toml"
	"github.com/daewood/config/xml"
	"github.com/daewood/config/yaml"
)

func New(f string, c interface{}) error {

	buf, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}

	//fmt.Println(string(buf))

	fileSuffix := getFileSuffix(f)

	switch fileSuffix {
	case "json":
		err = json.Unmarshal(buf, c)
	case "yaml":
		err = yaml.Unmarshal(buf, c)
	case "xml":
		err = xml.Unmarshal(buf, c)
	case "ini":
		err = ini.Unmarshal(buf, c)
	case "toml":
		err = toml.Unmarshal(buf, c)
	default:
		err = errors.New("Not support")
	}
	return err
}

func getFileSuffix(f string) string {
	filePath := path.Base(f)
	fileSuffix := path.Ext(filePath)
	return strings.Trim(fileSuffix, ".")
}
