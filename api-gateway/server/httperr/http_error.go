package httperr

import (
	"net/http"

	"github.com/go-chi/render"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func InternalError(err string, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, "Internal server error", w, r, http.StatusInternalServerError)
}

func BadRequest(err string, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, "Bad request", w, r, http.StatusBadRequest)
}

func NotFound(err string, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, "Not found", w, r, http.StatusNotFound)
}

func FromGrpcStatusError(err error, w http.ResponseWriter, r *http.Request) {
	statusErr, ok := status.FromError(err)
	if !ok {
		InternalError(err.Error(), w, r)
		return
	}

	switch statusErr.Code() {
	case codes.InvalidArgument:
		BadRequest(statusErr.Message(), w, r)
	case codes.AlreadyExists:
		BadRequest(statusErr.Message(), w, r)
	case codes.NotFound:
		NotFound(statusErr.Message(), w, r)
	default:
		InternalError(statusErr.Message(), w, r)
	}
}

func httpRespondWithError(err, msg string, w http.ResponseWriter, r *http.Request, status int) {
	resp := ErrorResponse{
		Error:          msg,
		Message:        err,
		httpStatusCode: status,
	}

	if err := render.Render(w, r, resp); err != nil {
		panic(err)
	}
}

type ErrorResponse struct {
	Error          string `json:"error"`
	Message        string `json:"message"`
	httpStatusCode int
}

func (e ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(e.httpStatusCode)
	return nil
}
