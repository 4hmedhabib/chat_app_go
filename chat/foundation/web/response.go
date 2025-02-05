package web

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

type NoResponse struct{}

func NewNoResponse() NoResponse {
	return NoResponse{}
}

func (NoResponse) Encode() ([]byte, string, error) {
	return nil, "", nil
}

type httpStatus interface {
	HTTPStatus() int
}

func Respond(ctx context.Context, w http.ResponseWriter, dataModel Encoder) error {
	if _, ok := dataModel.(NoResponse); ok {
		return nil
	}

	if err := ctx.Err(); err != nil {
		if errors.Is(err, context.Canceled) {
			return errors.New("client disconnected, do not send response")
		}
	}

	var statusCode = http.StatusOK

	switch v := dataModel.(type) {
	case httpStatus:
		statusCode = v.HTTPStatus()
	case error:
		statusCode = http.StatusInternalServerError
	default:
		if dataModel == nil {
			statusCode = http.StatusNoContent
		}
	}

	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
		return nil
	}

	data, contentType, err := dataModel.Encode()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return fmt.Errorf("respond: encode: %w", err)
	}

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)

	if _, err := w.Write(data); err != nil {
		return fmt.Errorf("respond: write: %w", err)
	}

	return nil
}
