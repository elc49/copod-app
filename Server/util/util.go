package util

import (
	"encoding/json"
	"io"
	"net/http"
)

func WriteHttp(w http.ResponseWriter, v interface{}, code int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	if _, err := w.Write(b); err != nil {
		return err
	}
	return nil
}

func DecodeHttp(in io.ReadCloser, out interface{}) error {
	if err := json.NewDecoder(in).Decode(out); err != nil {
		return err
	}
	return nil
}

func DecodeJson(in []byte, out interface{}) error {
	if err := json.Unmarshal(in, out); err != nil {
		return err
	}

	return nil
}

func EncodeJson(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
