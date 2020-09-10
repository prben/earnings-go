package earningsgo

import (
	"github.com/prben/earnings-go/pkg/balancesheet"
	"github.com/prben/earnings-go/pkg/cashflow"
	"github.com/prben/earnings-go/pkg/income"
	"github.com/prben/earnings-go/pkg/recommended"
)

// Earnings var
type Earnings struct {
	Cashflow     *cashflow.Cashflow
	BalanceSheet *balancesheet.BalanceSheet
	Income       *income.Income
	Recommended  []recommended.Recommended
}

// NewEarnings fetch
func NewEarnings(symbol string) (*Earnings, error) {

	var err error
	e := new(Earnings)

	e.BalanceSheet, err = balancesheet.NewBalanceSheet(symbol)
	if err != nil {
		return nil, err
	}

	e.Cashflow, err = cashflow.NewCashflow(symbol)
	if err != nil {
		return nil, err
	}

	e.Income, err = income.NewIncomeSheet(symbol)
	if err != nil {
		return nil, err
	}

	e.Recommended, err = recommended.NewRecommended(symbol)
	if err != nil {
		return nil, err
	}

	return e, nil
}
