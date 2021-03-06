package greeting_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/goph/emperror"
	"github.com/sagikazarmark/modern-go-application/.gen/openapi/go"
	"github.com/sagikazarmark/modern-go-application/internal/greeting"
	"github.com/sagikazarmark/modern-go-application/internal/greeting/greetingadapter"
	"github.com/sagikazarmark/modern-go-application/internal/greeting/greetingdriver"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testHelloWorld(t *testing.T) {
	helloWorld := greeting.NewHelloWorld(greetingadapter.NewNopLogger())
	controller := greetingdriver.NewGreetingController(helloWorld, nil, emperror.NewNopHandler())

	server := httptest.NewServer(http.HandlerFunc(controller.HelloWorld))

	resp, err := http.Get(server.URL)
	require.NoError(t, err)
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	var hello api.Hello

	err = decoder.Decode(&hello)
	require.NoError(t, err)

	assert.Equal(
		t,
		api.Hello{
			Message: "Hello, World!",
		},
		hello,
	)
}
