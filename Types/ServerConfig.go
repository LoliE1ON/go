package Types

import "github.com/LoliE1ON/go/Net/Db/MongoDb"

type ServerConfig struct {
	Port  int
	Mongo MongoDb.Config
}
