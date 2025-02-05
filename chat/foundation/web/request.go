package web

import (
	"fmt"
	"io"
	"net/http"
)

func Param(r *http.Request, key string) string {
	return r.PathValue(key)
}

type Decoder interface {
	Decode(data []byte) error
}

type validator interface {
	Validator() error
}

func Decode(r *http.Request, v Decoder) error {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("request: unable to read payload: %w", err)
	}

	if err := v.Decode(data); err != nil {
		return fmt.Errorf("request: decode: %w", err)
	}

	if v, ok := v.(validator); ok {
		if err := v.Validator(); err != nil {
			return err
		}
	}

	return nil
}
