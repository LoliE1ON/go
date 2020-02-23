package Types

import "github.com/LoliE1ON/go/Net/Db/Mongo"

type ServerConfig struct {
	Port      int
	Mongo     Mongo.Config
	JwtSecret string
	JwtExp    int
}
