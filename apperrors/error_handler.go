package apperrors

import (
	"encoding/json"
	"errors"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	var appErr *MyAppError
	if errors.As(err, &appErr) {
		appErr = &MyAppError{
			ErrCode: Unknown,
			Message: "internal server error",
			Err:     err,
		}
	}

	var statusCode int
	switch appErr.ErrCode {
	case NAData:
		statusCode = http.StatusNotFound
	case NoTargetData, BadParam, ReqBodyDecodeFailed:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)
}