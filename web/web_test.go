package web

import (
	"encoding/json"
	"github.com/Sfeir/handsongo/dao"
	"github.com/Sfeir/handsongo/model"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSpiritHandlerGet(t *testing.T) {

	// get mock dao
	daoMock, _ := dao.GetSpiritDAO("", dao.DAOMock)
	controller := NewSpiritController(daoMock)

	// build a request
	req, err := http.NewRequest(http.MethodGet, "localhost/v1/spirits", nil)
	if err != nil {
		t.Fatal(err)
	}

	// build the recorder
	res := httptest.NewRecorder()

	// execute the query
	controller.GetAll(res, req)

	var spiritOut []model.Spirit
	json.NewDecoder(res.Body).Decode(&spiritOut)

	if err != nil {
		t.Errorf("Unable to get JSON content %v", err)
	}

	if res.Code != http.StatusOK {
		t.Error("Wrong response code")
	}

	if len(spiritOut) < 1 {
		t.Fatal("Wrong response body size")
	}

	if !dao.MockedSpirit.Equal(spiritOut[0]) {
		t.Errorf("Expected different from %v output %v", dao.MockedSpirit, spiritOut[0])
	}
}

func TestSpiritHandlerGetServer(t *testing.T) {

	srv, err := BuildWebServer("", dao.DAOMock, 250*time.Millisecond)

	if err != nil {
		t.Error(err)
	}

	ts := httptest.NewServer(srv)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/v1/spirits")

	if err != nil {
		t.Error(err)
	}

	var resSpirit []model.Spirit
	err = json.NewDecoder(res.Body).Decode(&resSpirit)

	if err != nil {
		t.Errorf("Unable to get JSON content %v", err)
	}

	res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Error("Wrong response code")
	}

	if len(resSpirit) < 1 {
		t.Fatal("Wrong response body size")
	}

	if !dao.MockedSpirit.Equal(resSpirit[0]) {
		t.Error("Wrong response body content")
	}
}

func BenchmarkSpiritHandlerGet(t *testing.B) {

	// tooling purpose
	_ = new([45000]int)

	// get mock dao
	daoMock, _ := dao.GetSpiritDAO("", dao.DAOMock)
	controller := NewSpiritController(daoMock)

	// build a request
	req, err := http.NewRequest("GET", "localhost/v1/spirits", nil)
	if err != nil {
		t.Fatal(err)
	}

	// build the recorder
	res := httptest.NewRecorder()

	// execute the query
	controller.GetAll(res, req)

	var spiritOut []model.Spirit
	err = json.NewDecoder(res.Body).Decode(&spiritOut)

	if err != nil {
		t.Errorf("Unable to get JSON content %v", err)
	}

	if !dao.MockedSpirit.Equal(spiritOut[0]) {
		t.Errorf("Expected different from %v output %v", dao.MockedSpirit, spiritOut)
	}
}
