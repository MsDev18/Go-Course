package mysqluser

import (
	"E-01/entity"
	"E-01/pkg/errmsg"
	"E-01/pkg/richerror"
	"E-01/repository/mysql"
	"database/sql"
	"fmt"
	"time"
)

func (d *DB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	const op = "mysql.IsPhoneNumberUnique"

	row := d.conn.Conn().QueryRow(`select * from users where phone_number = ?`, phoneNumber)
	_, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}

		return false, richerror.New(op).
			WithErr(err).
			WithMessage(errmsg.ErrMsgCantScanQueryResult).
			WithKind(richerror.KindUnexpected)
	}
	return false, nil
}

func (d *DB) Register(u entity.User) (entity.User, error) {
	res, err := d.conn.Conn().Exec(`insert into users(name, phone_number, password, role) values(? , ?, ?, ?);`, u.Name, u.PhoneNumber, u.Password, u.Role.String())
	if err != nil {
		return entity.User{}, fmt.Errorf("can't execute command: %w", err)
	}
	// err is always nil
	id, _ := res.LastInsertId()
	u.ID = uint(id)
	return u, nil
}

func (d *DB) GetUserByPhoneNumber(phoneNumber string) (entity.User, error) {
	const op = "mysql.GetUserByPhoneNumber"

	row := d.conn.Conn().QueryRow(`select * from users where phone_number = ?`, phoneNumber)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, richerror.New(op).
				WithErr(err).WithMessage(errmsg.ErrMsgNotFound).
				WithKind(richerror.KindNotFound)
		}

	}
	return user, nil
}

func (d *DB) GetUserByID(userID uint) (entity.User, error) {
	const op = "mysql.GetUserByID"
	row := d.conn.Conn().QueryRow(`select * from users where id = ?`, userID)

	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, richerror.New(op).
				WithErr(err).
				WithMessage(errmsg.ErrMsgNotFound).
				WithKind(richerror.KindNotFound)
		}
		return entity.User{}, richerror.New(op).
			WithErr(err).
			WithMessage(errmsg.ErrMsgCantScanQueryResult).
			WithKind(richerror.KindUnexpected)
	}
	return user, nil
}

func scanUser(scanner mysql.Scanner) (entity.User, error) {
	var user entity.User
	var createdAt time.Time
	var roleStr string

	err := scanner.Scan(&user.ID, &user.Name, &user.PhoneNumber, &createdAt, &user.Password, &roleStr)

	user.Role = entity.MapToRoleEntity(roleStr)
	return user, err
}
