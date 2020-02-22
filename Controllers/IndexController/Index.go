package IndexController

import (
	"net/http"

	"github.com/LoliE1ON/go/Models/UserModel"

	"github.com/LoliE1ON/go/Helpers/HttpHelper"

	"github.com/LoliE1ON/go/Types"
)

func Action(w http.ResponseWriter, r *http.Request) {

	var response Types.ResponseData

	data, err := UserModel.GetAll()
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

	response.Data = data

	HttpHelper.ResponseWriter(response, w)
}
