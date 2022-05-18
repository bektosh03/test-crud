package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/bektosh03/test-crud/data-service/domain/post"
	"github.com/bektosh03/test-crud/data-service/internal/urlpaginator"
	"github.com/bektosh03/test-crud/data-service/internal/worker"
	"github.com/sirupsen/logrus"
)

const (
	numPages = 50
	postsURL = "https://gorest.co.in/public/v1/posts"
)

type App struct {
	postRepo post.Repository
}

func New(postRepo post.Repository) App {
	return App{
		postRepo: postRepo,
	}
}

func (a App) DownloadPosts(ctx context.Context) error {
	wp := worker.NewPool(4)
	jobs := a.createJobs(postsURL, numPages)

	startTime := time.Now()

	go wp.GenerateJobs(jobs)
	go wp.Run(ctx)

	postsBatch := make(chan []post.Post)
	defer close(postsBatch)

	for {
		select {
		case result, ok := <-wp.Results():
			if !ok {
				continue
			}
			if result.Err != nil {
				logrus.WithFields(logrus.Fields{
					"job_id":   result.Descriptor.JobID,
					"metadata": result.Descriptor.Metadata,
					"error":    result.Err.Error(),
				}).Error("Failed job")
				return result.Err
			}

			posts, ok := result.Value.([]post.Post)
			if !ok {
				return fmt.Errorf("wrong value type: %T", result.Value)
			}

			postsBatch <- posts

		case <-wp.Done():
			logrus.Infof("elapsed in: %d ms", time.Since(startTime).Milliseconds())
			return nil
		}
	}
}

func (a App) downloadPosts(ctx context.Context, arg interface{}) (interface{}, error) {
	url, ok := arg.(string)
	if !ok {
		return nil, fmt.Errorf("wrong argument type: %T", arg)
	}

	httpResponse, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if !(httpResponse.StatusCode >= 200 && httpResponse.StatusCode < 300) {
		body, err := io.ReadAll(httpResponse.Body)
		if err != nil {
			logrus.Errorf("failed to read from response body: %v", err)
			return nil, err
		}

		logrus.Errorf("unsuccessfull response from server: %v", string(body))
		return nil, fmt.Errorf("server responded with: %d", httpResponse.StatusCode)
	}

	var response GetPostsResponse
	if err = json.NewDecoder(httpResponse.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (a App) createJobs(url string, pagesCount int) []worker.Job {
	jobs := make([]worker.Job, 0, pagesCount)
	urlPaginator := urlpaginator.New(url, pagesCount)
	for urlPaginator.NextPage() {
		jobs = append(jobs, worker.Job{
			Descriptor: worker.JobDescriptor{
				JobID: urlPaginator.CurrentPage(),
				Metadata: map[string]interface{}{
					"url": urlPaginator.URL(),
				},
			},
			Exec: a.downloadPosts,
			Arg:  urlPaginator.URL(),
		})
	}

	return jobs
}

type GetPostsResponse struct {
	Data []post.Post `json:"data"`
}
