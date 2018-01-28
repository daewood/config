// toml file resolve.
//
// https://github.com/daewood/config
package toml

import (
	"github.com/BurntSushi/toml"
)

func Unmarshal(f []byte, result interface{}) error {
	_, err := toml.Decode(string(f), result)
	return err
}
