package app

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"

	"github.com/bektosh03/test-crud/data-service/app/urlpaginator"
	"github.com/bektosh03/test-crud/data-service/app/worker"
	"github.com/bektosh03/test-crud/data-service/domain/post"
	"github.com/sirupsen/logrus"
)

const (
	NumPages = 50
	postsURL = "https://gorest.co.in/public/v1/posts"
)

type App struct {
	wp worker.Pool
}

func New() App {
	return App{
		wp: worker.NewPool(runtime.NumCPU()),
	}
}

func (a App) DownloadPosts(ctx context.Context) error {
	jobs := make([]worker.Job, 0, NumPages)
	var execFn worker.ExecFn = func(ctx context.Context, arg interface{}) (interface{}, error) {
		url, ok := arg.(string)
		if !ok {
			return nil, fmt.Errorf("wrong argument type: %T", arg)
		}

		request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, &bytes.Buffer{})
		if err != nil {
			return nil, err
		}
		httpResponse, err := http.DefaultClient.Do(request)
		if err != nil {
			return nil, err
		}
		if !(httpResponse.StatusCode >= 200 && httpResponse.StatusCode < 300) {
			logrus.Infof("STATUS CODE: %d", httpResponse.StatusCode)
			body, err := io.ReadAll(httpResponse.Body)
			if err != nil {
				logrus.Errorf("failed to read from response body")
				return nil, err
			}
			logrus.Infof("unsuccessfull response from server: %v", string(body))
			return nil, fmt.Errorf("unsuccessfull response from server: %v", string(body))
		}
		logrus.Infof("URL: %s, HTTP RESPONSE: %v, BODY: %v", url, httpResponse, httpResponse.Body)

		var response GetPostsResponse
		if err = json.NewDecoder(httpResponse.Body).Decode(&response); err != nil {
			return nil, err
		}

		return response.Data, nil
	}

	urlPaginator := urlpaginator.New(postsURL, NumPages)
	logrus.Info("BEFORE FOR LOOP")
	for urlPaginator.NextPage() {
		jobs = append(jobs, worker.Job{
			Descriptor: worker.JobDescriptor{
				ID: urlPaginator.CurrentPage(),
				Metadata: map[string]interface{}{
					"url": urlPaginator.URL(),
				},
			},
			Exec: execFn,
			Arg:  urlPaginator.URL(),
		})
	}

	logrus.Info("AFTER FOR LOOP")

	go a.wp.GenerateJobs(jobs)
	go a.wp.Run(ctx)

	for {
		select {
		case result, ok := <-a.wp.Results():
			if !ok {
				continue
			}

			logrus.Infof("VALUE: %v", result.Value)
			val, ok := result.Value.([]post.Post)
			if !ok {
				return fmt.Errorf("wrong value type: %T", result.Value)
			}

			logrus.Infof("fetched: %v", val)

		case <-a.wp.Done():
			return nil
			// default:
			// 	logrus.Info("DEFAULT")
		}
	}
}

type GetPostsResponse struct {
	Data []post.Post `json:"data"`
}
