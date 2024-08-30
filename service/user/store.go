package user

import (
    "database/sql"
    "fmt"

    "github.com/hd2yao/ecom/types"
)

type Store struct {
    db *sql.DB
}

func NewStore(db *sql.DB) *Store {
    return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.Users, error) {
    rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
    if err != nil {
        return nil, err
    }

    u := new(types.Users)
    for rows.Next() {
        u, err = scanRowsIntoUser(rows)
        if err != nil {
            return nil, err
        }
    }

    if u.ID == 0 {
        return nil, fmt.Errorf("user not found")
    }

    return u, nil
}

func scanRowsIntoUser(rows *sql.Rows) (*types.Users, error) {
    user := new(types.Users)

    err := rows.Scan(
        &user.ID,
        &user.FirstName,
        &user.LastName,
        &user.Email,
        &user.Password,
        &user.CreatedAt,
    )
    if err != nil {
        return nil, err
    }
    return user, nil
}
