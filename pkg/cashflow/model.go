package cashflow

import "github.com/prben/earnings-go/pkg/entry"

type CashflowResponse struct {
	Timeseries struct {
		Result []Cashflow `json:"result"`
	} `json:"timeseries"`
}

type Cashflow struct {
	AnnualOperatingCashFlow                  []entry.Entry `json:"annualOperatingCashFlow,omitempty"`
	AnnualPurchaseOfInvestment               []entry.Entry `json:"annualPurchaseOfInvestment,omitempty"`
	AnnualBeginningCashPosition              []entry.Entry `json:"annualBeginningCashPosition,omitempty"`
	AnnualInvestingCashFlow                  []entry.Entry `json:"annualInvestingCashFlow,omitempty"`
	AnnualCapitalExpenditure                 []entry.Entry `json:"annualCapitalExpenditure,omitempty"`
	AnnualChangeInInventory                  []entry.Entry `json:"annualChangeInInventory,omitempty"`
	AnnualEndCashPosition                    []entry.Entry `json:"annualEndCashPosition,omitempty"`
	TrailingNetIncome                        []entry.Entry `json:"trailingNetIncome,omitempty"`
	AnnualOtherNonCashItems                  []entry.Entry `json:"annualOtherNonCashItems,omitempty"`
	AnnualDepreciationAndAmortization        []entry.Entry `json:"annualDepreciationAndAmortization,omitempty"`
	AnnualRepaymentOfDebt                    []entry.Entry `json:"annualRepaymentOfDebt,omitempty"`
	AnnualChangeInWorkingCapital             []entry.Entry `json:"annualChangeInWorkingCapital,omitempty"`
	AnnualChangeInCashSupplementalAsReported []entry.Entry `json:"annualChangeInCashSupplementalAsReported,omitempty"`
	AnnualSaleOfInvestment                   []entry.Entry `json:"annualSaleOfInvestment,omitempty"`
	AnnualNetIncome                          []entry.Entry `json:"annualNetIncome,omitempty"`
	AnnualCommonStockIssuance                []entry.Entry `json:"annualCommonStockIssuance,omitempty"`
	AnnualFreeCashFlow                       []entry.Entry `json:"annualFreeCashFlow,omitempty"`
}
