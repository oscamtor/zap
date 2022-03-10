package main


import (
	"testing"

	"go.uber.org/zap/zaptest"
)


func TestHello(t *testing.T) {


	got := Hello()
	want := "Hello, worldd"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}



	logger := zaptest.NewLogger(t)
	
	//logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	
	const url = "http://example.com"
	const attempt = 3
	
	sugar.Infow("failed to fetch URL",
	  // Structured context as loosely typed key-value pairs.
	  "url", url,
	  "attempt", attempt,
	  //"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}