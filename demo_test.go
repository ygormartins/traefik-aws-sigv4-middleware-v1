package sigv4middleware_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	plugin "github.com/ygormartins/traefikAwsSigv4Middlewarev1"

	"testing"
)

func TestHandler(t *testing.T) {
	c := plugin.CreateConfig()
	c.AccessKey = "Q3AM3UQ867SPQQA43P2F"
	c.SecretKey = "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	c.Service = "s3"
	c.Endpoint = "play.min.io"
	c.Region = "us-east-1"

	ctx := context.Background()

	// Make Bucket
	bucketName := "treafikmiddlewares3v4sig"
	objectName := "index.html"

	next := http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {})

	handler, err := plugin.New(ctx, next, c, "foo")

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	reqURL := fmt.Sprintf("http://%s/%s/%s", c.Endpoint, bucketName, objectName)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	res := recorder.Result()
	defer func() {
		if err := res.Body.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	body, err := io.ReadAll(recorder.Result().Body)

	if err != nil {
		t.Fatal(err)
	}

	if recorder.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %v: %v", recorder.Result().StatusCode, string(body))
	}
}
