package controller

import (
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	mock_dao "backend/internal/dao/mock"
	"backend/internal/model"
)

var (
	mockUser = model.User{
		ID:         1,
		Username:   "testuser",
		Firstname:  "Test",
		Lastname:   "User",
		Email:      "testuser@example.com",
		Status:     "A",
		Department: "IT",
	}
)

func TestController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller Suite")
}

var _ = Describe("Controller", func() {
	var (
		testCtrl *gomock.Controller
		mockDao  *mock_dao.MockDaoInterface

		ctrl *Controller
	)

	BeforeEach(func() {
		testCtrl = gomock.NewController(GinkgoT())
		mockDao = mock_dao.NewMockDaoInterface(testCtrl)

		ctrl = NewController(mockDao)
	})

	AfterEach(func() {
		testCtrl.Finish()
	})

	Describe("GetUser", func() {
		It("should return a list of users", func() {
			mockDao.EXPECT().GetUser(mockUser.ID).Return(mockUser, nil)

			user, err := ctrl.GetUser(mockUser.ID)
			Expect(err).To(BeNil())
			Expect(user).To(Equal(mockUser))
		})
	})

	Describe("CreateUser", func() {
		It("should create a user", func() {
			mockDao.EXPECT().GetUserByUsername(mockUser.Username).Return(model.User{}, nil)
			mockDao.EXPECT().CreateUser(mockUser).Return(1, nil)

			id, err := ctrl.CreateUser(mockUser)
			Expect(err).To(BeNil())
			Expect(id).To(Equal(1))
		})
	})

	// Add tests for update and delete
})
