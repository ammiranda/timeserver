package rest_api

import (
	"encoding/json"
	"fmt"
	"github.com/ammiranda/timeserver/pkg/rest_api/models/response"
	"github.com/ammiranda/timeserver/pkg/timestamp"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTimeStampRoute_NoTime(t *testing.T) {
	s := timestamp.NewService()
	r := NewRouter(s)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/api/", nil)
	r.ServeHTTP(w, req)

	resp := response.TimeResponse{}
	require.Equal(t, 200, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.True(t, resp.Unix > 0)
}

func TestTimeStampRoute_DateTimeStr_Valid(t *testing.T) {
	s := timestamp.NewService()
	r := NewRouter(s)
	w := httptest.NewRecorder()

	getURL := fmt.Sprintf("/api/%s", "2015-12-25")

	req, _ := http.NewRequest("GET", getURL, nil)
	r.ServeHTTP(w, req)

	// resp := response.TimeResponse{}
	require.Equal(t, 200, w.Code)
}

func TestTimeStampRoute_With_Invalid_Date_Bad_Request(t *testing.T) {
	s := timestamp.NewService()
	r := NewRouter(s)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/api/kadkfja", nil)
	r.ServeHTTP(w, req)

	resp := response.TimeResponseError{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Equal(t, 400, w.Code)
	require.Equal(t, errorResponse, resp)
}
