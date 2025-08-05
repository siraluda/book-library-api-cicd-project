package main

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// )

// func TestGetBookHandler(t *testing.T)  {
// 	req, err := http.NewRequest(http.MethodGet,"/books/1", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	resr := httptest.NewRecorder()

// 	handler := http.HandlerFunc()
// 	handler.ServeHTTP(resr, req)

// 	if status := resr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// 	if resr.Body.String() == "" {
// 		t.Errorf("handler returned unexpected body: got %v want %v",
// 			resr.Body.String(), "A book object")
// 	}

// }