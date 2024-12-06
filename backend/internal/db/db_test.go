package db

import (
	"backend/internal/model"
	"database/sql"
	"testing"

	_ "github.com/lib/pq" // Import the PostgreSQL driver

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/DATA-DOG/go-sqlmock"
)

var (
	mockUser = model.User{
		ID:         1,
		Username:   "testuser",
		Firstname:  "Test",
		Lastname:   "User",
		Email:      "test@example.com",
		Status:     "A",
		Department: "IT",
	}
)

func TestDB(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DB Suite")
}

var _ = Describe("DB", func() {
	var (
		db   *DB
		mock sqlmock.Sqlmock
	)

	BeforeEach(func() {
		// Create SQL mock
		mockDB, sqlMock, err := sqlmock.New()
		Expect(err).NotTo(HaveOccurred())

		// Set the mock DB in your DB struct
		db = &DB{db: mockDB}
		mock = sqlMock
	})

	AfterEach(func() {
		mock.ExpectClose()
		db.db.Close()
	})

	Describe("GetUser", func() {
		It("should retrieve a user by ID", func() {
			// Define expected rows
			rows := sqlmock.NewRows([]string{
				"user_id",
				"user_name",
				"first_name",
				"last_name",
				"email",
				"user_status",
				"department",
				"created_at",
				"updated_at",
			}).
				AddRow(
					mockUser.ID,
					mockUser.Username,
					mockUser.Firstname,
					mockUser.Lastname,
					mockUser.Email,
					mockUser.Status,
					mockUser.Department,
					mockUser.CreatedAt,
					mockUser.UpdatedAt,
				)

			// Expect the query and return mocked rows
			mock.ExpectQuery(`SELECT (.+) FROM users WHERE user_id = \$1`).
				WithArgs(1).
				WillReturnRows(rows)

			user, err := db.GetUser(1)
			Expect(err).To(BeNil())
			Expect(user.ID).To(Equal(1))
			Expect(user.Email).To(Equal("test@example.com"))
		})

		It("should return error when user not found", func() {
			mock.ExpectQuery(`SELECT (.+) FROM users WHERE user_id = \$1`).
				WithArgs(1).
				WillReturnError(sql.ErrNoRows)

			_, err := db.GetUser(1)
			Expect(err).To(HaveOccurred())
		})
	})

	// Add more Describe blocks for other methods
})
