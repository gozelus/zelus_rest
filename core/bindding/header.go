package binding

import "net/http"

type headerBinding struct{}

func (headerBinding) Name() string {
	return "header"
}

func (headerBinding) Bind(req *http.Request, obj interface{}) error {
	values := req.Header
	if err := mapForm(obj, values, "header"); err != nil {
		return err
	}
	return nil
}
