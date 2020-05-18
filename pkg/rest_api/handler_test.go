package rest_api

import (
	"encoding/json"
	"github.com/ammiranda/timeserver/pkg/rest_api/models/response"
	"github.com/ammiranda/timeserver/pkg/timestamp"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTimeStampRoute_NoTime(t *testing.T) {
	s := timestamp.NewService()
	r := Handler(s)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/api/timestamp", nil)
	r.ServeHTTP(w, req)

	resp := response.TimeResponse{}
	require.Equal(t, 200, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.True(t, resp.Unix > 0)
}
