package repositories

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/lathief/learn-fiber-go/pkg/models"
)

type userRepository struct {
	DB *sqlx.DB
}

type UserRepository interface {
	Create(ctx context.Context, user models.User) (int64, error)
	GetAll(ctx context.Context) ([]models.User, error)
	GetById(ctx context.Context, id int64) (models.User, error)
	GetByEmail(ctx context.Context, email string) (models.User, error)
	GetByUsername(ctx context.Context, username string) (models.User, error)
	Update(ctx context.Context, user models.User) error
	Delete(ctx context.Context, id int64) error
}

func NewUserRepository(DB *sqlx.DB) UserRepository {
	return &userRepository{
		DB: DB,
	}
}

func (u *userRepository) Create(ctx context.Context, user models.User) (int64, error) {
	res, err := u.DB.NamedExecContext(ctx, "INSERT INTO users (username, first_name, last_name, email, password, phone_number, address, role_id) VALUES (:username, :first_name, :last_name, :email, :password, :phone_number, :address, :role_id)", user)
	if err != nil {
		return 0, err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return 0, errors.New("Rows not affected")
	}
	userId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func (u *userRepository) GetAll(ctx context.Context) ([]models.User, error) {
	var users []models.User
	err := u.DB.SelectContext(ctx, &users, `SELECT * FROM users`)
	return users, err
}

func (u *userRepository) GetByUsername(ctx context.Context, username string) (models.User, error) {
	var user models.User
	err := u.DB.GetContext(ctx, &user,
		`SELECT * FROM users WHERE username = ?`, username)
	if err != nil {
		return models.User{}, err
	}
	return user, err
}

func (u *userRepository) GetByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	err := u.DB.GetContext(ctx, &user,
		`SELECT * FROM users WHERE email = ?`, email)
	if err != nil {
		return models.User{}, err
	}
	return user, err
}

func (u *userRepository) GetById(ctx context.Context, id int64) (models.User, error) {
	var user models.User
	err := u.DB.GetContext(ctx, &user,
		`SELECT * FROM users WHERE id = ?`, id)
	if err != nil {
		return models.User{}, err
	}
	return user, err
}

func (u *userRepository) Update(ctx context.Context, user models.User) error {
	res, err := u.DB.NamedExecContext(ctx,
		`UPDATE users SET first_name=:first_name, last_name=:last_name, email=:email, username=:username, 
                 password=:password, phone_number=:phone_number, address=:address, role_id=:role_id WHERE id = :id`,
		user)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (u *userRepository) Delete(ctx context.Context, id int64) error {
	res, err := u.DB.ExecContext(ctx, "DELETE FROM users WHERE id=?", id)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
