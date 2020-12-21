package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gozelus/zelus_rest/logger"
)

const (
	ApplicationJson = "application/json"
	ContentEncoding = "Content-Encoding"
	ContentType     = "Content-Type"
)

func Error(w http.ResponseWriter, code int, err error) {
	WriteJson(w, code, err.Error())
}
func OkJson(w http.ResponseWriter, v interface{}) {
	WriteJson(w, http.StatusOK, v)
}

func WriteJson(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set(ContentType, ApplicationJson)
	w.WriteHeader(code)

	if bs, err := json.Marshal(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if n, err := w.Write(bs); err != nil {
		// http.ErrHandlerTimeout has been handled by http.TimeoutHandler,
		// so it's ignored here.
		if err != http.ErrHandlerTimeout {
			logger.Errorf("write response failed, error: %s", err)
		}
	} else if n < len(bs) {
		logger.Errorf("actual bytes: %d, written bytes: %d", len(bs), n)
	}
}

const (
	maxBodyLen = 8 << 20
)

func JsonBodyFromRequest(req *http.Request, jsonStruct interface{}) error {
	bytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	defer req.Body.Close()
	return json.Unmarshal(bytes, jsonStruct)
}
