package IndexController

import (
	"net/http"

	"github.com/LoliE1ON/go/Helpers/HttpHelper"

	"github.com/LoliE1ON/go/Types"
)

func Action(w http.ResponseWriter, r *http.Request) {

	var response Types.ResponseData
	response.Data = "Index action"

	HttpHelper.ResponseWriter(response, w)
}
