package repositories

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/lathief/learn-fiber-go/pkg/models"
)

type roleRepository struct {
	DB *sqlx.DB
}

type RoleRepository interface {
	Create(ctx context.Context, role models.Role) error
	GetAll(ctx context.Context) ([]models.Role, error)
	GetById(ctx context.Context, id int64) (models.Role, error)
	GetByName(ctx context.Context, name string) (models.Role, error)
	Update(ctx context.Context, role models.Role) error
	Delete(ctx context.Context, id int64) error
}

func NewRoleRepository(DB *sqlx.DB) RoleRepository {
	return &roleRepository{
		DB: DB,
	}
}

func (r *roleRepository) Create(ctx context.Context, role models.Role) error {
	res, err := r.DB.NamedExecContext(ctx, "INSERT INTO roles (role_name, description) VALUES (:role_name, :description)", role)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return err
	}
	return nil
}

func (r *roleRepository) GetAll(ctx context.Context) ([]models.Role, error) {
	var roles []models.Role
	err := r.DB.SelectContext(ctx, &roles, `SELECT * FROM roles`)
	return roles, err
}

func (r *roleRepository) GetById(ctx context.Context, id int64) (models.Role, error) {
	var role models.Role
	err := r.DB.GetContext(ctx, &role,
		`SELECT * FROM roles WHERE id = ?`, id)
	if err != nil {
		return models.Role{}, err
	}
	return role, err
}
func (r *roleRepository) GetByName(ctx context.Context, name string) (models.Role, error) {
	var role models.Role
	err := r.DB.GetContext(ctx, &role,
		`SELECT * FROM roles WHERE role_name = ?`, name)
	if err != nil {
		return models.Role{}, err
	}
	return role, err
}
func (r *roleRepository) Update(ctx context.Context, role models.Role) error {
	res, err := r.DB.NamedExecContext(ctx,
		`UPDATE category SET name=:name, 
				description= CASE WHEN :description IS NOT NULL AND LENGTH(:description) > 0 THEN :description ELSE description END
				WHERE id = :id`,
		role)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *roleRepository) Delete(ctx context.Context, id int64) error {
	res, err := r.DB.ExecContext(ctx, "DELETE FROM roles WHERE id=?", id)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
