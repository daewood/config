// ini file resolve.
//
// https://github.com/daewood/config
package ini

import (
	"gopkg.in/gcfg.v1"
)

// resolve ini
func Unmarshal(f []byte, result interface{}) error {
	return gcfg.ReadStringInto(result, string(f))
}
