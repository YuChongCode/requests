package requests

import (
	"encoding/json"
	"io"
	"net/http"
)

type Response struct {
	*http.Response
}

func (r *Response) JSON(data interface{}) error {
	defer r.Close()
	return json.NewDecoder(r.Body).Decode(data)
}

func (r *Response) Content() (string, error) {
	defer r.Close()
	content, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (r *Response) Close() {
	_ = r.Body.Close()
}
