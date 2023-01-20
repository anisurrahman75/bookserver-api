package api

import (
	"apiDesign/middleware"
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Welcome(t *testing.T) {
	s := CreateNewServer()
	s.MountHandlers()

	req, _ := http.NewRequest("GET", "/", nil)
	response := executeRequest(req, s)
	checkResponseCode(t, http.StatusOK, response.Code)
}
func Test_logIn(t *testing.T) {
	s := CreateNewServer()
	s.MountHandlers()
	type Test struct {
		method             string
		url                string
		body               io.Reader
		expectedStatusCode int
		token              string
	}
	test := []Test{
		{"POST",
			"/logIn",
			nil,
			http.StatusOK,
			middleware.BasicToken("sunny2741", "123")},
		{"POST",
			"/logIn",
			nil,
			http.StatusUnauthorized,
			middleware.BasicToken("snny241",
				"123")},
		{"POST",
			"/lgIn",
			nil,
			http.StatusNotFound,
			middleware.BasicToken("mridul12", "123")},
	}
	for _, i := range test {
		req, _ := http.NewRequest(i.method, i.url, i.body)
		req.Header.Set("Authorization", i.token)
		response := executeRequest(req, s)
		checkResponseCode(t, i.expectedStatusCode, response.Code)
	}
}
func Test_Register(t *testing.T) {
	s := CreateNewServer()
	s.MountHandlers()
	type Test struct {
		method             string
		url                string
		body               io.Reader
		expectedStatusCode int
	}
	test := []Test{
		{
			"POST",
			"/registerUser",
			bytes.NewReader([]byte(`{"FirstName": "imtiaj", "LastName": "Halder", "UserName": "imtiaj12", "Password": "123"}`)),
			http.StatusOK,
		},
		{
			"POST",
			"/registerUser",
			bytes.NewReader([]byte(`{"FirstName": "imtiaj", "LastName": "Halder", "UserName": "imtiaj12, "Password": "123"}`)),
			http.StatusBadRequest,
		},
	}
	for _, i := range test {
		req, _ := http.NewRequest(i.method, i.url, i.body)
		response := executeRequest(req, s)
		checkResponseCode(t, i.expectedStatusCode, response.Code)
	}
}
func Test_AllBookList(t *testing.T) {
	s := CreateNewServer()
	s.MountHandlers()
	type Test struct {
		method             string
		url                string
		body               io.Reader
		token              string
		expectedStatusCode int
	}
	test := []Test{
		{
			"GET",
			"/books",
			nil,
			middleware.BearerToken("testing"),
			200,
		},
		{
			"GET",
			"/books",
			nil,
			middleware.BearerToken("testing"),
			200,
		},
	}
	for _, i := range test {
		req, _ := http.NewRequest(i.method, i.url, i.body)
		req.Header.Set("Authorization", i.token)
		response := executeRequest(req, s)
		checkResponseCode(t, i.expectedStatusCode, response.Code)
	}
}
func Test_AddBook(t *testing.T) {
	s := CreateNewServer()
	s.MountHandlers()
	type Test struct {
		method             string
		url                string
		body               io.Reader
		token              string
		expectedStatusCode int
	}
	test := []Test{
		{"POST",
			"/books",
			bytes.NewReader([]byte(`{"UUID": 10, "Name": "learn-api", "Author": "Anisur", "PublishDate": "01-02-2022", "ISBN": "0999-0555-5954"}`)),
			middleware.BearerToken("sunny"),
			http.StatusOK,
		},
		{"POST",
			"/books",
			bytes.NewReader([]byte(`{"UUID: 10, "Name": "learn-api", "Author": "Anisur", "PublishDate": "01-02-2022", "ISBN": "0999-0555-5954"}`)),
			middleware.BearerToken("sunny"),
			http.StatusBadRequest,
		},
	}
	for _, i := range test {
		middleware.VerifyJWT(s.Router)
		req, _ := http.NewRequest(i.method, i.url, i.body)
		req.Header.Set("Authorization", i.token)
		response := executeRequest(req, s)
		checkResponseCode(t, i.expectedStatusCode, response.Code)
	}
}
func Test_FindBook(t *testing.T) {
	s := CreateNewServer()
	s.MountHandlers()
	type Test struct {
		method             string
		url                string
		body               io.Reader
		token              string
		expectedStatusCode int
	}
	test := []Test{
		{"GET",
			"/books/1",
			nil,
			middleware.BearerToken("sunny"),
			http.StatusOK,
		},
		{"GET",
			"/books/10",
			nil,
			middleware.BearerToken("sunny"),
			http.StatusBadRequest,
		},
	}
	for _, i := range test {
		req, _ := http.NewRequest(i.method, i.url, i.body)
		req.Header.Set("Authorization", i.token)
		response := executeRequest(req, s)
		checkResponseCode(t, i.expectedStatusCode, response.Code)
	}
}
func Test_UpdateBook(t *testing.T) {
	s := CreateNewServer()
	s.MountHandlers()
	type Test struct {
		method             string
		url                string
		body               io.Reader
		token              string
		expectedStatusCode int
	}
	test := []Test{
		{"PUT",
			"/books/4",
			bytes.NewReader([]byte(`{"UUID": 4, "Name": "learn-api", "Author": "RaselVai", "PublishDate": "01-02-2022", "ISBN": "0999-0555-5954"},`)),
			middleware.BearerToken("sunny"),
			http.StatusOK,
		},
		{"PUT",
			"/books/100",
			bytes.NewReader([]byte(`{"UUID": 4, "Name": "learn-api", "Author": "RaselVai", "PublishDate": "01-02-2022", "ISBN": "0999-0555-5954"},`)),
			middleware.BearerToken("sunny"),
			http.StatusBadRequest,
		},
	}
	for _, i := range test {
		req, _ := http.NewRequest(i.method, i.url, i.body)
		req.Header.Set("Authorization", i.token)
		response := executeRequest(req, s)
		checkResponseCode(t, i.expectedStatusCode, response.Code)
	}
}
func Test_DeleteBook(t *testing.T) {
	s := CreateNewServer()
	s.MountHandlers()
	type Test struct {
		method             string
		url                string
		body               io.Reader
		token              string
		expectedStatusCode int
	}
	test := []Test{
		{"DELETE",
			"/books/4",
			nil,
			middleware.BearerToken("sunny"),
			http.StatusOK,
		},
		{"DELETE",
			"/books/100",
			nil,
			middleware.BearerToken("sunny"),
			http.StatusBadRequest,
		},
	}
	for _, i := range test {
		req, _ := http.NewRequest(i.method, i.url, i.body)
		req.Header.Set("Authorization", i.token)
		response := executeRequest(req, s)
		checkResponseCode(t, i.expectedStatusCode, response.Code)
	}
}
func executeRequest(req *http.Request, s *Server) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)
	return rr
}
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}