package reposytory

import (
	"errors"
	"github.com/doug-martin/goqu/v9"
	"my-bank-service/internal/config"
	"my-bank-service/internal/data"
	"my-bank-service/pkg/logging"
	"my-bank-service/pkg/session"
)

type AuthRepository interface {
	FetchAuth(authD *data.AuthDetails) (*data.Auth, error)
	DeleteAuth(authD *data.AuthDetails) error
	CreateAuth(authD *data.AuthDetails) (*data.Auth, error)
}

// authRepository has the implementation of the db methods.
type authRepository struct {
	session *session.Session
	logger  logging.Logger
}

// NewAuthRepository returns a new userRepository instance
func NewAuthRepository(s *session.Session, l logging.Logger) AuthRepository {
	return &authRepository{s, l}
}

func (a *authRepository) FetchAuth(authD *data.AuthDetails) (*data.Auth, error) {
	dialect := goqu.Dialect(config.Driver)
	ds := dialect.From(config.AuthTable).Where(
		goqu.Ex{config.UserId: authD.UserId, config.AuthUUID: authD.AuthUuid})
	sqlStr, _, err := ds.ToSQL()
	if err != nil {
		a.logger.Error("An error occurred while generating the SQL: ", err)
		return nil, err
	}
	res, err := a.session.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = res.Close()
		if err != nil {
			a.logger.Error(err)
		}
	}()
	au := &data.Auth{}
	if res.Next() {
		err = res.Scan(&au.ID, &au.UserID, &au.AuthUUID)
		if err != nil {
			return nil, err
		}
		return au, nil
	} else {
		return nil, errors.New("invalid token")
	}
}

func (a *authRepository) DeleteAuth(authD *data.AuthDetails) error {
	dialect := goqu.Dialect(config.Driver)
	ds := dialect.Delete(config.AuthTable).Where(
		goqu.Ex{config.UserId: authD.UserId, config.AuthUUID: authD.AuthUuid})
	sqlStr, _, err := ds.ToSQL()
	if err != nil {
		a.logger.Error("An error occurred while generating the SQL: ", err)
		return err
	}
	_, err = a.session.Exec(sqlStr)
	if err != nil {
		a.logger.Error(err)
		return err
	}
	return nil
}

func (a *authRepository) CreateAuth(authD *data.AuthDetails) (*data.Auth, error) {
	auth, err := a.FetchAuth(authD)
	if auth != nil {
		return auth, nil
	}
	au := &data.Auth{}
	au.AuthUUID = authD.AuthUuid
	au.UserID = authD.UserId
	dialect := goqu.Dialect(config.Driver)
	ds := dialect.Insert(config.AuthTable).Rows(
		goqu.Record{config.UserId: au.UserID, config.AuthUUID: au.AuthUUID})
	sqlStr, _, err := ds.ToSQL()
	if err != nil {
		a.logger.Error("An error occurred while generating the SQL: ", err)
		return nil, err
	}
	_, err = a.session.Exec(sqlStr)
	if err != nil {
		a.logger.Error(err)
		return nil, err
	}
	return au, nil
}
