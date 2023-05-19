package metadata

import (
	"github.com/spoonboy-io/koan"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

var testCache = ".testCache"
var testTTL = 1 * time.Minute

func createTestServer(data string, status int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseBody := data
		responseStatusCode := status
		w.WriteHeader(responseStatusCode)
		w.Write([]byte(responseBody))
	}))
}

func TestGetMetadata(t *testing.T) {
	logger := &koan.Logger{}

	testCases := []struct {
		name           string
		serverResponse string
		serverStatus   int
		wantData       []byte
		wantErr        error
	}{
		{
			"read from server and cache the data",
			"test data for cache",
			200,
			[]byte("test data for cache"),
			nil,
		},
		{
			"read from cache",
			"test data on server is now this, so we should be using cache",
			200,
			[]byte("test data for cache"),
			nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// create test server
			testServer := createTestServer(tc.serverResponse, tc.serverStatus)

			gotData, gotErr := GetMetadata(testServer.URL, testCache, testTTL, logger)
			if gotErr != tc.wantErr {
				t.Errorf("wanted '%v' got '%v'", tc.wantErr, gotErr)
			}

			if string(gotData) != string(tc.wantData) {
				t.Errorf("wanted '%v' got '%v'", string(tc.wantData), string(gotData))
			}

			testServer.Close()
		})
	}

	if err := os.Remove(testCache); err != nil {
		t.Fatalf("something went wrong: %v", err)
	}
}

func TestGetMetadataBadResponse(t *testing.T) {
	logger := &koan.Logger{}

	testCases := []struct {
		name         string
		serverStatus int
	}{
		{
			"bad response",
			500,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// create test server
			testServer := createTestServer("", tc.serverStatus)
			_, gotErr := GetMetadata(testServer.URL, testCache, testTTL, logger)
			if gotErr == nil {
				t.Errorf("wanted error but got nil")
			}

			testServer.Close()
		})
	}
}

func TestHaveCached(t *testing.T) {
	testCases := []struct {
		name        string
		createCache bool
		delay       time.Duration
		TTL         time.Duration
		want        bool
	}{
		{
			"no cache file, should be false",
			false,
			0,
			1 * time.Second,
			false,
		},
		{
			"has cache file, fresher than TTL, should be true",
			true,
			0,
			1 * time.Second,
			true,
		},
		{
			"has cache file, older than TTL, should be false",
			true,
			2 * time.Second,
			1 * time.Second,
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			if tc.createCache {
				// create cache file
				if err := os.WriteFile(testCache, []byte("test data"), 0700); err != nil {
					t.Fatalf("something went wrong: %v", err)
				}
				defer func() {
					if err := os.Remove(testCache); err != nil {
						t.Fatalf("something went wrong: %v", err)
					}
				}()
			}

			time.Sleep(tc.delay) // allow the cache file to hit TTL

			got := haveCached(testCache, tc.TTL)
			if got != tc.want {
				t.Errorf("wanted '%v' got '%v'", tc.want, got)
			}

		})
	}
}
