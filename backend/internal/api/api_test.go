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
			req, _ := http.NewRequest("GET", "/api/v1/status", nil)
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

			req, _ := http.NewRequest("POST", "/api/v1/user", reqBody)
			req.Header.Set("Content-Type", "application/json")
			res := httptest.NewRecorder()
			c := echo.New().NewContext(req, res)

			expectedUser := model.User{
				Email: "testuser@example.com",
			}
			mockCtrl.EXPECT().CreateUser(expectedUser).Return(1, nil)

			err := api.createUser(c)
			Expect(err).To(BeNil())
			Expect(res.Code).To(Equal(http.StatusCreated))
		})

		It("should return an internal server error if the db returns an error", func() {
			reqBody := strings.NewReader(`{
				"email": "testuser@example.com"
			}`)

			req, _ := http.NewRequest("POST", "/api/v1/user", reqBody)
			req.Header.Set("Content-Type", "application/json")
			res := httptest.NewRecorder()
			c := echo.New().NewContext(req, res)

			mockCtrl.EXPECT().CreateUser(gomock.Any()).Return(0, fmt.Errorf("db error"))

			err := api.createUser(c)
			Expect(err).To(BeNil())
			Expect(res.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Describe("getUser", func() {
		It("should get a user", func() {
			req, _ := http.NewRequest("GET", "/api/v1/user/1", nil)
			res := httptest.NewRecorder()
			c := echo.New().NewContext(req, res)
			c.SetParamNames("id")
			c.SetParamValues("1")

			mockCtrl.EXPECT().GetUser(1).Return(model.User{ID: 1, Email: "testuser@example.com"}, nil)

			err := api.getUser(c)
			Expect(err).To(BeNil())
			Expect(res.Code).To(Equal(http.StatusOK))
		})
	})

	Describe("updateUser", func() {
		It("should update a user", func() {
			reqBody := strings.NewReader(`{
				"email": "testuser@example.com"
			}`)

			req, _ := http.NewRequest("PUT", "/api/v1/user/1", reqBody)
			req.Header.Set("Content-Type", "application/json")
			res := httptest.NewRecorder()
			c := echo.New().NewContext(req, res)
			c.SetParamNames("id")
			c.SetParamValues("1")

			mockCtrl.EXPECT().UpdateUser(1, gomock.Any()).Return(model.User{ID: 1, Email: "testuser@example.com"}, nil)

			err := api.updateUser(c)
			Expect(err).To(BeNil())
			Expect(res.Code).To(Equal(http.StatusOK))
		})
	})

	Describe("deleteUser", func() {
		It("should delete a user", func() {
			req, _ := http.NewRequest("DELETE", "/api/v1/user/1", nil)
			res := httptest.NewRecorder()
			c := echo.New().NewContext(req, res)
			c.SetParamNames("id")
			c.SetParamValues("1")

			mockCtrl.EXPECT().DeleteUser(1).Return(nil)

			err := api.deleteUser(c)
			Expect(err).To(BeNil())
			Expect(res.Code).To(Equal(http.StatusNoContent))
		})
	})

	Describe("getUsers", func() {
		It("should get users", func() {
			users := []model.User{
				{ID: 1, Email: "testuser1@example.com"},
				{ID: 2, Email: "testuser2@example.com"},
			}

			req, _ := http.NewRequest("GET", "/api/v1/users", nil)
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
			req, _ := http.NewRequest("GET", "/api/v1/users", nil)
			res := httptest.NewRecorder()
			c := echo.New().NewContext(req, res)

			mockCtrl.EXPECT().GetUsers().Return(nil, fmt.Errorf("db error"))

			err := api.getUsers(c)
			Expect(err).To(BeNil())
			Expect(res.Code).To(Equal(http.StatusInternalServerError))
		})
	})
})
