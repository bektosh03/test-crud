package ports

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bektosh03/test-crud/api-gateway/server/httperr"
	"github.com/bektosh03/test-crud/genprotos/datapb"
	"github.com/bektosh03/test-crud/genprotos/postpb"
	"github.com/go-chi/chi/v5"
)

type HTTPServer struct {
	dataServiceClient datapb.DataServiceClient
	postServiceClient postpb.PostServiceClient
}

func NewHTTPServer(dataServiceClient datapb.DataServiceClient) HTTPServer {
	return HTTPServer{
		dataServiceClient: dataServiceClient,
	}
}

// DownloadPosts handler
// @Router /posts/download [post]
// @Summary start downloading posts
// @Description API for starting the download of posts
// @Tags posts
// @Produce json
// @Success 200 {object} ResponseMessage
// @Failure 500 {object} httperr.ErrorResponse
func (s HTTPServer) DownloadPosts(w http.ResponseWriter, r *http.Request) {
	_, err := s.dataServiceClient.DownloadPosts(context.Background(), &datapb.DownloadPostsRequest{})
	if err != nil {
		httperr.FromGrpcStatusError(err, w, r)
		return
	}

	s.respondJSON(w, r, ResponseMessage{
		Message: "download of posts started",
	}, http.StatusOK)
}

// GetDownloadStatus handler
// @Router /posts/download/status [get]
// @Summary report back status of downloading posts
// @Description API for checking the status of posts download
// @Tags posts
// @Produce json
// @Success 200 {object} ResponseMessage
// @Failure 500 {object} httperr.ErrorResponse
func (s HTTPServer) GetDownloadStatus(w http.ResponseWriter, r *http.Request) {
	response, err := s.dataServiceClient.GetDownloadStatus(r.Context(), &datapb.GetDownloadStatusRequest{})
	if err != nil {
		httperr.FromGrpcStatusError(err, w, r)
		return
	}

	s.respondJSON(w, r, response, http.StatusOK)
}

// GetPost handler
// @Router /posts/{post-id} [get]
// @Summary get individual post
// @Description API for fetching an individual post
// @Tags posts
// @Produce json
// @Success 200 {object} Post
// @Failure 404 {object} httperr.ErrorResponse
// @Failure 500 {object} httperr.ErrorResponse
func (s HTTPServer) GetPost(w http.ResponseWriter, r *http.Request) {
	postIDParam := chi.URLParam(r, "post-id")
	postID, err := strconv.ParseInt(postIDParam, 10, 64)
	if err != nil {
		httperr.BadRequest(fmt.Sprintf("bad value for param 'post-id': %s", postIDParam), w, r)
		return
	}

	post, err := s.postServiceClient.GetPost(r.Context(), &postpb.GetPostRequest{
		PostId: postID,
	})
	if err != nil {
		httperr.FromGrpcStatusError(err, w, r)
		return
	}

	s.respondJSON(w, r, post, http.StatusOK)
}

// ListPost handler
// @Router /posts [get]
// @Summary get list of posts
// @Description API for fetching a list of posts
// @Tags posts
// @Produce json
// @Success 200 {object} ListPostsResponse
// @Failure 500 {object} httperr.ErrorResponse
func (s HTTPServer) ListPosts(w http.ResponseWriter, r *http.Request) {
	page, limit := s.parsePageAndLimitQueryParams(r)
	posts, err := s.postServiceClient.ListPosts(r.Context(), &postpb.ListPostsRequest{
		Page:  page,
		Limit: limit,
	})
	if err != nil {
		httperr.FromGrpcStatusError(err, w, r)
		return
	}

	s.respondJSON(w, r, posts, http.StatusOK)
}

// UpdatePost handler
// @Router /posts [put]
// @Summary update post info
// @Description API for updating post's info
// @Tags posts
// @Produce json
// @Success 200 {object} ResponseMessage
// @Failure 404 {object} httperr.ErrorResponse
// @Failure 500 {object} httperr.ErrorResponse
func (s HTTPServer) UpdatePost(w http.ResponseWriter, r *http.Request) {
	var request UpdatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		httperr.BadRequest(err.Error(), w, r)
		return
	}

	_, err := s.postServiceClient.UpdatePost(r.Context(), &postpb.Post{
		Id:    int64(request.PostID),
		Title: request.Title,
		Body:  request.Body,
	})
	if err != nil {
		httperr.FromGrpcStatusError(err, w, r)
		return
	}

	s.respondJSON(w, r, ResponseMessage{
		Message: "ok",
	}, http.StatusOK)
}

// DeletePost handler
// @Router /posts [delete]
// @Summary delete a post
// @Description API for deleting a post
// @Tags posts
// @Produce json
// @Success 200 {object} ResponseMessage
// @Failure 404 {object} httperr.ErrorResponse
// @Failure 500 {object} httperr.ErrorResponse
func (s HTTPServer) DeletePost(w http.ResponseWriter, r *http.Request) {
	var request DeletePostRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		httperr.BadRequest(err.Error(), w, r)
		return
	}

	_, err := s.postServiceClient.DeletePost(r.Context(), &postpb.DeletePostRequest{
		PostId: int64(request.PostID),
	})
	if err != nil {
		httperr.FromGrpcStatusError(err, w, r)
		return
	}

	s.respondJSON(w, r, ResponseMessage{
		Message: "ok",
	}, http.StatusOK)
}

func (s HTTPServer) respondJSON(w http.ResponseWriter, r *http.Request, body interface{}, status int) {
	bytesBody, err := json.MarshalIndent(body, " ", "  ")
	if err != nil {
		httperr.InternalError(err.Error(), w, r)
		return
	}

	w.WriteHeader(status)
	w.Write(bytesBody)
}

func (s HTTPServer) parsePageAndLimitQueryParams(r *http.Request) (int64, int64) {
	queries := r.URL.Query()
	pageVal, limitVal := queries.Get("page"), queries.Get("limit")

	page, _ := strconv.ParseInt(pageVal, 10, 64)
	limit, _ := strconv.ParseInt(limitVal, 10, 64)

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	return page, limit
}
