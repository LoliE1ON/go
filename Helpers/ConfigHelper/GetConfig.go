package ConfigHelper

import (
	"log"

	"github.com/LoliE1ON/go/Types"
)

// Get config file
func GetConfig() (config Types.ServerConfig, err error) {
	config, err = ParseConfig()
	if err != nil {
		log.Println(err)
		return
	}
	return
}
