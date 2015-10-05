// Parking Backend - DAO
//
// User DAO.
// Note that an won't have its password loaded.
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

	insert             *Stmt
	create             *Stmt
	getCryptedPassword *Stmt
	findByEmail        *Stmt
	findByParking      *Stmt
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

	if d.create, err = d.db.Prepare(`
		INSERT INTO "user"
		(
			"uid",
			"email",
			"firstname",
			"password",
			"creation_time",
			"last_update"
		)
		VALUES ($1, $2, $3, $4, $5, $6);
	`); err != nil {
		return err
	}

	if d.getCryptedPassword, err = d.db.Prepare(`
		SELECT "user"."password"
		FROM "user"
		WHERE "user"."uid" = $1
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

	if d.findByParking, err = d.db.Prepare(`
		SELECT ` + USER_FIELDS + ` 
		FROM "user" join "parking" on "parking"."user_id" = "user"."uid"
		WHERE "parking"."uid" = $1
	`); err != nil {
		return err
	}

	return nil
}

func (d *UserDAO) Create(uid, email, firstname, password string, creationTime time.Time) (Result, error) {
	return d.create.Exec(
		uid,
		email,
		firstname,
		password,
		creationTime,
		creationTime,
	)
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

func (d *UserDAO) GetCryptedPassword(user model.User) (string, error) {
	rows, err := d.getCryptedPassword.Query(user.Uid.String())

	if err != nil || rows == nil {
		return "", err
	}

	defer rows.Close()

	var pwd string
	if rows.Next() {
		err = rows.Scan(&pwd)
	}
	return pwd, err
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

func (d *UserDAO) FindByParking(parkingId uuid.UUID) (model.User, error) {
	found := model.User{}

	rows, err := d.findByParking.Query(parkingId.String())
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
