package response

import (
	"encoding/json"
	"net/http"
	"fmt"
)

const ContentTypeHeader = "Content-Type"

func View(viewName string, data interface{}) {

}

func Json(w http.ResponseWriter, data interface{}){
	jsonResponse, ok := json.Marshal(data)
	if ok == nil {
		w.Header().Set(ContentTypeHeader, "application/json")
		fmt.Fprintf(w, string(jsonResponse))
	}
}
