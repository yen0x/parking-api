// Parking Backend - DAO
//
// User DAO.
//
// 2015

package dao

import (
	. "database/sql"
	"time"

	"bitbucket.org/remeh/parking/db/model"

	"github.com/pborman/uuid"
)

type UserDAO struct {
	db *DB

	insert      *Stmt
	findByEmail *Stmt
}

const (
	USER_FIELDS = `"user"."uid",
				   "user"."email",
				   "user"."firstname",
				   "user"."lastname",
				   "user"."gender",
				   "user"."phone",
				   "user"."address",
				   "user"."creation_time",
				   "user"."last_update"`
)

func NewUserDAO(db *DB) (*UserDAO, error) {
	dao := &UserDAO{
		db: db,
	}
	err := dao.initStmt()
	return dao, err
}

func (d *UserDAO) initStmt() error {
	var err error

	if d.insert, err = d.db.Prepare(`
		INSERT INTO "user"
		(` + insertFields("user", USER_FIELDS) + `)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
	`); err != nil {
		return err
	}

	if d.findByEmail, err = d.db.Prepare(`
		SELECT ` + USER_FIELDS + ` 
		FROM "user"
		WHERE email = $1
	`); err != nil {
		return err
	}

	return nil
}

func (d *UserDAO) Insert(user model.User) (Result, error) {
	if len(user.Uid) == 0 {
		return nil, nil
	}

	return d.insert.Exec(
		user.Uid.String(),
		user.Email,
		user.Firstname,
		user.Lastname,
		user.Gender,
		user.Phone,
		user.Address,
		user.CreationTime,
		user.LastUpdate,
	)
}

func (d *UserDAO) FindByEmail(email string) (model.User, error) {
	found := model.User{}

	if len(email) == 0 {
		return found, nil
	}

	rows, err := d.findByEmail.Query(email)
	if rows == nil || err != nil {
		return found, err
	}

	defer rows.Close()

	if !rows.Next() {
		return found, nil
	}

	return userFromRow(rows)
}

// userFromRow reads an user model from the current row.
func userFromRow(rows *Rows) (model.User, error) {
	var uid,
		email,
		firstname,
		lastname,
		gender,
		phone,
		address string
	var creationTime,
		lastUpdate time.Time

	err := rows.Scan(&uid,
		&email,
		&firstname,
		&lastname,
		&gender,
		&phone,
		&address,
		&creationTime,
		&lastUpdate,
	)

	return model.User{
		Uid:          uuid.Parse(uid),
		Email:        email,
		Firstname:    firstname,
		Lastname:     lastname,
		Gender:       gender,
		Phone:        phone,
		Address:      address,
		CreationTime: creationTime,
		LastUpdate:   lastUpdate,
	}, err
}
