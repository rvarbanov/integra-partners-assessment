package db

import (
	"database/sql"
	"fmt"
	"log"

	"backend/internal/config"
	"backend/internal/model"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
)

type DBInterface interface {
	GetUser(ID int) (model.User, error)
	GetUsers() ([]model.User, error)
	CreateUser(user model.User) (int, error)
	UpdateUser(ID int, user model.User) (model.User, error)
	DeleteUser(ID int) error
	GetUserByUsername(username string) (model.User, error)
}

type DB struct {
	db *sql.DB
}

func NewDB(config config.Database) DBInterface {
	dataSourceName := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Pass,
		config.Name,
	)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return &DB{db: db}
}

func (d *DB) GetUser(ID int) (model.User, error) {
	var user model.User

	sqGetUser := sq.Select("*").
		From("users").
		Where(sq.Eq{"user_id": ID}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := sqGetUser.ToSql()
	if err != nil {
		return user, err
	}

	row := d.db.QueryRow(query, args...)
	err = row.Scan(
		&user.ID,
		&user.Username,
		&user.Firstname,
		&user.Lastname,
		&user.Email,
		&user.Status,
		&user.Department,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return user, err
	}

	return user, err
}

func (d *DB) GetUsers() ([]model.User, error) {
	var users []model.User

	sqGetUsers := sq.Select("*").From("users")

	query, args, err := sqGetUsers.ToSql()
	if err != nil {
		return users, err
	}

	rows, err := d.db.Query(query, args...)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Firstname,
			&user.Lastname,
			&user.Email,
			&user.Status,
			&user.Department,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, err
}

func (d *DB) CreateUser(user model.User) (int, error) {
	sqCreateUser := sq.Insert("users").
		Columns(
			"user_name",
			"first_name",
			"last_name",
			"email",
			"user_status",
			"department",
		).
		Values(
			user.Username,
			user.Firstname,
			user.Lastname,
			user.Email,
			user.Status,
			user.Department,
		).
		Suffix("RETURNING user_id").
		PlaceholderFormat(sq.Dollar)

	query, args, err := sqCreateUser.ToSql()
	if err != nil {
		return 0, fmt.Errorf("failed to build query: %w", err)
	}

	var userID int
	err = d.db.QueryRow(query, args...).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("user creation failed: no ID returned")
		}
		return 0, fmt.Errorf("failed to create user: %w", err)
	}

	return userID, err
}

func (d *DB) UpdateUser(ID int, user model.User) (model.User, error) {
	sqUpdatedUser := sq.Update("users").
		Set("user_name", user.Username).
		Set("first_name", user.Firstname).
		Set("last_name", user.Lastname).
		Set("email", user.Email).
		Set("user_status", user.Status).
		Set("department", user.Department).
		Where(sq.Eq{"user_id": ID}).
		Suffix("RETURNING user_id, user_name, first_name, last_name, email, user_status, department, created_at, updated_at").
		PlaceholderFormat(sq.Dollar)

	query, args, err := sqUpdatedUser.ToSql()
	if err != nil {
		return user, err
	}

	row := d.db.QueryRow(query, args...)
	err = row.Scan(
		&user.ID,
		&user.Username,
		&user.Firstname,
		&user.Lastname,
		&user.Email,
		&user.Status,
		&user.Department,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return user, err
	}

	return user, err
}

func (d *DB) DeleteUser(ID int) error {
	sqDeleteUser := sq.Delete("users").
		Where(sq.Eq{"user_id": ID}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := sqDeleteUser.ToSql()
	if err != nil {
		return err
	}

	result, err := d.db.Exec(query, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with ID %d not found", ID)
	}

	return err
}

func (d *DB) GetUserByUsername(username string) (model.User, error) {
	var user model.User

	sqGetUser := sq.Select("*").
		From("users").
		Where(sq.Eq{"user_name": username}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := sqGetUser.ToSql()
	if err != nil {
		return user, fmt.Errorf("failed to build query: %w", err)
	}

	row := d.db.QueryRow(query, args...)
	err = row.Scan(
		&user.ID,
		&user.Username,
		&user.Firstname,
		&user.Lastname,
		&user.Email,
		&user.Status,
		&user.Department,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return user, nil
	}

	if err != nil {
		return user, fmt.Errorf("failed to scan user: %w", err)
	}

	return user, err
}
