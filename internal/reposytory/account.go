package reposytory

import (
	"errors"
	"github.com/doug-martin/goqu/v9"
	"github.com/uniplaces/carbon"
	"math"
	"my-bank-service/internal/config"
	"my-bank-service/internal/data"
	"my-bank-service/internal/utils"
	"my-bank-service/pkg/logging"
	"my-bank-service/pkg/session"
)

type AccountRepository interface {
	Withdraw(userId string, sum float64) (float64, error)
}

// accountRepository has the implementation of the db methods.
type accountRepository struct {
	session *session.Session
	logger  logging.Logger
}

// NewAccountRepository returns a new userRepository instance
func NewAccountRepository(s *session.Session, l logging.Logger) AccountRepository {
	return &accountRepository{s, l}
}

func (a *accountRepository) Withdraw(userId string, sum float64) (float64, error) {
	err := a.session.Begin()
	if err != nil {
		return 0, err
	}
	defer func() {
		err = a.session.Rollback()
		if err != nil {
			a.logger.Error(err)
		}
	}()

	balance, err := a.fetchBalance(userId)
	if err != nil {
		return 0, err
	}
	floatBalance := balance.IntegerPart + balance.FractionPart
	if floatBalance < sum {
		return 0, errors.New("there are not enough funds in your accountRepository")
	}
	paymentHistory := &data.PaymentHistory{
		BalanceID:         balance.ID,
		InitialBalance:    floatBalance,
		DifferenceBalance: sum,
	}
	res := math.Dim(floatBalance, sum)
	balance.IntegerPart, balance.FractionPart = math.Modf(res)
	balance.FractionPart = utils.Round(balance.FractionPart, .5, 2)
	err = a.updateBalance(balance)
	if err != nil {
		return 0, err
	} else {
		paymentHistory.FinalBalance = balance.FractionPart + balance.IntegerPart
		err = a.createPaymentHistory(paymentHistory)
		if err != nil {
			return 0, err
		}
	}
	err = a.session.Commit()
	if err != nil {
		return 0, err
	}
	return paymentHistory.FinalBalance, nil
}

func (a *accountRepository) fetchBalance(userId string) (*data.Balance, error) {
	dialect := goqu.Dialect(config.Driver)
	ds := dialect.From(config.BalanceTable).Where(goqu.Ex{config.UserId: userId})
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
	balance := &data.Balance{}
	if res.Next() {
		err = res.Scan(&balance.ID, &balance.UserID, &balance.IntegerPart,
			&balance.FractionPart, &balance.Currency)
		if err != nil {
			return nil, err
		}
	}
	return balance, nil
}

func (a *accountRepository) updateBalance(balance *data.Balance) error {
	dialect := goqu.Dialect(config.Driver)
	ds := dialect.Update(config.BalanceTable).Set(
		goqu.Record{config.IntegerPart: balance.IntegerPart, config.FractionPart: balance.FractionPart,
			config.Currency: balance.Currency}).Where(
		goqu.Ex{config.UserId: balance.UserID})
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

func (a *accountRepository) createPaymentHistory(history *data.PaymentHistory) error {
	history.CreatedAt = carbon.Now().String()

	dialect := goqu.Dialect(config.Driver)
	ds := dialect.Insert(config.PaymentHistoryTable).Rows(
		goqu.Record{config.BalanceId: history.BalanceID,
			config.CreatedAt: history.CreatedAt, config.InitialBalance: history.InitialBalance,
			config.FinalBalance: history.FinalBalance, config.DifferenceBalance: history.DifferenceBalance})
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
