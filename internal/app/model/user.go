package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id       int
	Email    string
	Password string
	EncPass  string
}

// BeforeCreate validates user creation
func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encPass(u.Password)
		if err != nil {
			return err
		}

		u.EncPass = enc
	}
	return nil
}

func encPass(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
