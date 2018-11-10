package responses

import(
	"encoding/json"
	"net/http"
)

type Request struct {
	Request				*http.Request
	ResponseWriter		http.ResponseWriter
}

// Get JSON body
func (r *Request)GetJSONBody() (model interface{}) {
	decoder := json.NewDecoder(r.Request.Body)
	return decoder.Decode(&model)
}