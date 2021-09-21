package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/mux"

	"github.com/d-kuznetsov/calendar-backend/dto"
)

type FakeService struct{}

var userId string
var user dto.User
var eventId string

func (svc *FakeService) Register(userData dto.User) (string, error) {
	return userId, nil
}
func (svc *FakeService) CreateToken(id string) (string, error) {
	return "token" + id, nil
}
func (svc *FakeService) Login(userData dto.User) (dto.User, error) {
	return user, nil
}
func (svc *FakeService) ParseToken(token string) (string, error) {
	return userId, nil
}
func (svc *FakeService) CreateEvent(eventData dto.Event) (string, error) {
	return eventId, nil
}
func (svc *FakeService) UpdateEvent(eventData dto.Event) error {
	return nil
}
func (svc *FakeService) DeleteEvent(id string) error {
	return nil
}
func (svc *FakeService) GetEvents(params dto.PeriodParams) ([]dto.Event, error) {
	return make([]dto.Event, 0), nil
}

var router *mux.Router
var hdlr handler

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	router = mux.NewRouter()
	hdlr = CreateHandler(&FakeService{})
}

func TestRegister(t *testing.T) {
	router.HandleFunc("/register", hdlr.Register).Methods("POST")
	writer := httptest.NewRecorder()

	userId = "007"
	name := "James Bond"
	token := "token" + userId
	jsonStr := fmt.Sprintf(`{"name":"%v"}`, name)
	reqBody := strings.NewReader(jsonStr)

	request, _ := http.NewRequest("POST", "/register", reqBody)
	router.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var resBody struct {
		Token string `json:"token"`
		Name  string `json:"name"`
	}
	json.Unmarshal(writer.Body.Bytes(), &resBody)
	if resBody.Token != token {
		t.Errorf("Wrong token, was expecting %v but got %v", token, resBody.Token)
	}
	if resBody.Name != name {
		t.Errorf("Wrong name, was expecting %v but got %v", name, resBody.Name)
	}
}

func TestLogin(t *testing.T) {
	router.HandleFunc("/login", hdlr.Login).Methods("POST")
	writer := httptest.NewRecorder()

	userId = "007"
	name := "James Bond"
	user = dto.User{Id: userId, Name: name}
	token := "token" + userId

	request, _ := http.NewRequest("POST", "/login", strings.NewReader("{}"))
	router.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var resBody struct {
		Token string `json:"token"`
		Name  string `json:"name"`
	}
	json.Unmarshal(writer.Body.Bytes(), &resBody)
	if resBody.Token != token {
		t.Errorf("Wrong token, was expecting %v but got %v", token, resBody.Token)
	}
	if resBody.Name != name {
		t.Errorf("Wrong name, was expecting %v but got %v", name, resBody.Name)
	}

}
