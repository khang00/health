package integration_test

import (
	"github.com/khang00/health/internal/route"
	"github.com/khang00/health/internal/store"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"net/url"
	"strconv"
	"testing"
	"time"
)

func TestUserHandler_GetUserMeal(t *testing.T) {
	// Setup
	e := echo.New()

	// Set log level
	e.Logger.SetLevel(log.INFO)

	// Set up storage
	storeClient := store.NewStoreClient()
	defer storeClient.Closed()
	// Setup Route
	route.SetUpRoutes(e, storeClient)
	// Start server
	initChan := make(chan int, 0)
	go func() {
		if err := e.Start(":3000"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}

		initChan <- 1
	}()
	_ = <-initChan

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	from := strconv.FormatInt(time.Now().Add(-time.Hour*24).Unix(), 10)
	to := strconv.FormatInt(time.Now().Add(time.Hour*24).Unix(), 10)
	q := make(url.Values)
	q.Set("user_id", "1")
	q.Set("from", from)
	q.Set("to", to)

	req, err := http.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/user/meal/?"+q.Encode(), nil)
	if err != nil {
		t.Fail()
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fail()
	}

	if resp.StatusCode != http.StatusOK {
		t.Fail()
	}
}
