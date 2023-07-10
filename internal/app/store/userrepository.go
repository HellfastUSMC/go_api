package store

import "github.com/HellfastUSMC/goapi/internal/app/model"

// UserRepository user repository def
type UserRepository struct {
	store *Store
}

// Create creates new UserRepository item
func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}
	if err := r.store.db.QueryRow("INSERT INTO users (email, enc_pass) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncPass,
	).Scan(&u.Id); err != nil {
		return nil, err
	}
	return u, nil
}

// FindByEmail search by email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow("SELECT id, email, enc_pass FROM users WHERE email = $1",
		email,
	).Scan(
		&u.Id,
		&u.Email,
		&u.EncPass,
	); err != nil {
		return nil, err
	}
	return u, nil
}
