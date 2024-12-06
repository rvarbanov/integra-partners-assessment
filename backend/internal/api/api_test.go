package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	mock_controller "backend/internal/controller/mock"
	"backend/internal/model"
)

func TestAPI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "API Suite")
}

var _ = Describe("API", func() {
	var (
		testCtrl *gomock.Controller
		mockCtrl *mock_controller.MockControllerInterface

		api *API
	)

	BeforeEach(func() {
		testCtrl = gomock.NewController(GinkgoT())
		mockCtrl = mock_controller.NewMockControllerInterface(testCtrl)

		api = NewAPI(mockCtrl)
	})

	AfterEach(func() {
		testCtrl.Finish()
	})

	Describe("GetStatus", func() {
		It("should return a status", func() {
			req, _ := http.NewRequest("GET", "/status", nil)
			res := httptest.NewRecorder()
			c := echo.New().NewContext(req, res)

			err := api.getStatus(c)
			Expect(err).To(BeNil())
			Expect(res.Code).To(Equal(http.StatusOK))
		})
	})

	Describe("createUser", func() {
		It("should create a user", func() {
			reqBody := strings.NewReader(`{
				"username": "testuser",
				"email": "testuser@example.com"
			}`)

			req, _ := http.NewRequest("POST", "/user", reqBody)
			req.Header.Set("Content-Type", "application/json")
			res := httptest.NewRecorder()
			c := echo.New().NewContext(req, res)

			expectedUser := model.User{
				Email: "testuser@example.com",
			}
			mockCtrl.EXPECT().CreateUser(expectedUser).Return(1, nil)

			err := api.createUser(c)
			Expect(err).To(BeNil())
			Expect(res.Code).To(Equal(http.StatusOK))
		})

		It("should return an internal server error if the db returns an error", func() {
			reqBody := strings.NewReader(`{
				"email": "testuser@example.com"
			}`)

			req, _ := http.NewRequest("POST", "/user", reqBody)
			req.Header.Set("Content-Type", "application/json")
			res := httptest.NewRecorder()
			c := echo.New().NewContext(req, res)

			mockCtrl.EXPECT().CreateUser(gomock.Any()).Return(0, fmt.Errorf("db error"))

			err := api.createUser(c)
			Expect(err).To(BeNil())
			Expect(res.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Describe("getUsers", func() {
		It("should get users", func() {
			users := []model.User{
				{ID: 1, Email: "testuser1@example.com"},
				{ID: 2, Email: "testuser2@example.com"},
			}

			req, _ := http.NewRequest("GET", "/users", nil)
			res := httptest.NewRecorder()
			c := echo.New().NewContext(req, res)

			mockCtrl.EXPECT().GetUsers().Return(users, nil)

			err := api.getUsers(c)
			Expect(err).To(BeNil())
			Expect(res.Code).To(Equal(http.StatusOK))

			resp := model.Response{}
			err = json.Unmarshal(res.Body.Bytes(), &resp)
			Expect(err).To(BeNil())

			jsonData, err := json.Marshal(resp.Data)
			Expect(err).To(BeNil())

			var respUsers []model.User
			err = json.Unmarshal(jsonData, &respUsers)
			Expect(err).To(BeNil())

			Expect(respUsers).To(Equal(users))
		})

		It("should return an internal server error if the db returns an error", func() {
			req, _ := http.NewRequest("GET", "/users", nil)
			res := httptest.NewRecorder()
			c := echo.New().NewContext(req, res)

			mockCtrl.EXPECT().GetUsers().Return(nil, fmt.Errorf("db error"))

			err := api.getUsers(c)
			Expect(err).To(BeNil())
			Expect(res.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	// TODO: add tests for other endpoints
})
