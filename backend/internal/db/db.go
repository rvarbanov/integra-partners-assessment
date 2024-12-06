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
		Where(sq.Eq{"id": ID})

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
		err := rows.Scan(&users)
		if err != nil {
			return users, err
		}
	}
	// mock data 3 users
	users = append(users, model.User{ID: 1, Username: "test", Firstname: "test", Lastname: "test", Email: "test", Status: "test", Department: "test"})
	users = append(users, model.User{ID: 2, Username: "test2", Firstname: "test2", Lastname: "test2", Email: "test2", Status: "test2", Department: "test2"})
	users = append(users, model.User{ID: 3, Username: "test3", Firstname: "test3", Lastname: "test3", Email: "test3", Status: "test3", Department: "test3"})

	return users, err
}

func (d *DB) CreateUser(user model.User) (int, error) {
	sqCreateUser := sq.Insert("users").
		Columns("user_name", "first_name", "last_name", "email", "user_status", "department").
		Values(user.Username, user.Firstname, user.Lastname, user.Email, user.Status, user.Department).
		Suffix("RETURNING \"id\"")

	query, args, err := sqCreateUser.ToSql()
	if err != nil {
		return 0, err
	}

	rows, err := d.db.Query(query, args...)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Firstname,
			&user.Lastname,
			&user.Email,
			&user.Status,
			&user.Department,
		)
		if err != nil {
			return 0, err
		}
	}

	return user.ID, err
}

func (d *DB) UpdateUser(ID int, user model.User) (model.User, error) {
	sqUpdatedUser := sq.Update("users").
		Set("user_name", user.Username).
		Set("first_name", user.Firstname).
		Set("last_name", user.Lastname).
		Set("email", user.Email).
		Set("user_status", user.Status).
		Set("department", user.Department).
		Where(sq.Eq{"id": ID}).
		Suffix("RETURNING \"*\"")

	query, args, err := sqUpdatedUser.ToSql()
	if err != nil {
		return user, err
	}

	rows, err := d.db.Query(query, args...)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Firstname,
			&user.Lastname,
			&user.Email,
			&user.Status,
			&user.Department,
		)
		if err != nil {
			return user, err
		}
	}

	return user, err
}

func (d *DB) DeleteUser(ID int) error {
	sqDeleteUser := sq.Delete("users").Where(sq.Eq{"id": ID})

	query, args, err := sqDeleteUser.ToSql()
	if err != nil {
		return err
	}

	_, err = d.db.Exec(query, args...)
	return err
}

func (d *DB) GetUserByUsername(username string) (model.User, error) {
	var user model.User

	sqGetUser := sq.Select("*").
		From("users").
		Where(sq.Eq{"user_name": username})

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
	)
	if err != nil {
		return user, err
	}

	return user, err
}
