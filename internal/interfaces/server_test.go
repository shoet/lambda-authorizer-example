package interfaces_test

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/shoet/lambda-authorizer-example/internal/interfaces"
)

func Test_Server_HealthHandler(t *testing.T) {
	srv, err := interfaces.NewServer()
	if err != nil {
		t.Fatalf("failed to create server: %v", err)
	}

	request := httptest.NewRequest("GET", "/health", nil)
	request.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	healthHandler := interfaces.NewHealthHandler()
	textCtx := srv.NewContext(request, w)

	if err := healthHandler.Handler(textCtx); err != nil {
		t.Fatalf("failed to call handler: %v", err)
	}

	type ResponseBody struct {
		Message string `json:"message"`
	}
	var got ResponseBody
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	want := ResponseBody{
		Message: "OK",
	}
	if cmp.Diff(got, want) != "" {
		t.Fatalf("response body is not expected: %v", cmp.Diff(got, want))
	}
}
