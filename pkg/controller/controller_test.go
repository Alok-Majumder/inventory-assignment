package controller_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Alok-Majumder/inventory-assignment/pkg/controller"
)

func MockController() *controller.Controller {
	return &controller.Controller{}
}

func MockHttpRequest() (*httptest.ResponseRecorder, *http.Request) {
	req, err := http.NewRequest("GET", "/products/wrongURI", nil)
	if err != nil {
		log.Fatal(err)
		return nil, nil
	}
	rr := httptest.NewRecorder()

	return rr, req

}

func TestControllerServeHTTPDoNotAcceptOtherURI(t *testing.T) {

	c := MockController()
	rr, req := MockHttpRequest()

	c.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {

		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)

	}

}
