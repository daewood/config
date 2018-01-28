// json file resolve.
//
// https://github.com/daewood/config
package json

import (
	"bytes"
	"encoding/json"
)

func Unmarshal(f []byte, result interface{}) error {

	decoder := json.NewDecoder(bytes.NewReader(f))
	decoder.UseNumber()
	return decoder.Decode(result)
}
