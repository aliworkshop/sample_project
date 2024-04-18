package tests

import (
	"encoding/json"
	"github.com/aliworkshop/echoserver"
	"github.com/aliworkshop/gateway/v2"
	"github.com/aliworkshop/logger"
	"github.com/aliworkshop/logger/writers"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewHiHandler(t *testing.T) {
	app := testApp()
	respond := echoserver.NewResponder(i18n.NewBundle(language.English))
	c := gateway.NewController(respond, logger.NewSimpleLogger(writers.WarnLevel, logger.JsonEncoding))
	s := echoserver.NewTestServer(c)
	rg := s.NewRouterGroup("/")
	rg.READ("hi", app.HiModule.Hi)

	req, _ := http.NewRequest(http.MethodGet, "/hi", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Language", "fa")
	w := httptest.NewRecorder()

	rg.ServeHttp(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var res gateway.Response
	err := json.NewDecoder(w.Body).Decode(&res)
	assert.Nil(t, err)

	assert.Equal(t, res.Page, 1)
	assert.Equal(t, res.PerPage, 10)
	assert.Equal(t, res.Total, uint64(1))
	assert.Len(t, res.Items, 1)
	assert.Equal(t, res.Items.(map[string]any)["message"], "hello world")
}
