package balancesheet

import (
	"github.com/prben/earnings-go/pkg/entry"
)

type BalanceSheetResponse struct {
	Timeseries struct {
		Result []BalanceSheet `json:"result"`
	} `json:"timeseries"`
}

type BalanceSheet struct {
	AnnualTotalNonCurrentLiabilitiesNetMinorityInterest []entry.Entry `json:"annualTotalNonCurrentLiabilitiesNetMinorityInterest,omitempty"`
	AnnualCapitalStock                                  []entry.Entry `json:"annualCapitalStock,omitempty"`
	AnnualOtherNonCurrentLiabilities                    []entry.Entry `json:"annualOtherNonCurrentLiabilities,omitempty"`
	AnnualAccountsReceivable                            []entry.Entry `json:"annualAccountsReceivable,omitempty"`
	AnnualCurrentAssets                                 []entry.Entry `json:"annualCurrentAssets,omitempty"`
	AnnualCurrentLiabilities                            []entry.Entry `json:"annualCurrentLiabilities,omitempty"`
	AnnualTotalNonCurrentAssets                         []entry.Entry `json:"annualTotalNonCurrentAssets,omitempty"`
	AnnualOtherNonCurrentAssets                         []entry.Entry `json:"annualOtherNonCurrentAssets,omitempty"`
	AnnualOtherCurrentLiabilities                       []entry.Entry `json:"annualOtherCurrentLiabilities,omitempty"`
	AnnualNetPPE                                        []entry.Entry `json:"annualNetPPE,omitempty"`
	AnnualGrossPPE                                      []entry.Entry `json:"annualGrossPPE,omitempty"`
	AnnualRetainedEarnings                              []entry.Entry `json:"annualRetainedEarnings,omitempty"`
	AnnualAccountsPayable                               []entry.Entry `json:"annualAccountsPayable,omitempty"`
	AnnualCashAndCashEquivalents                        []entry.Entry `json:"annualCashAndCashEquivalents,omitempty"`
	AnnualInventory                                     []entry.Entry `json:"annualInventory,omitempty"`
	AnnualAccumulatedDepreciation                       []entry.Entry `json:"annualAccumulatedDepreciation,omitempty"`
	AnnualTotalAssets                                   []entry.Entry `json:"annualTotalAssets,omitempty"`
	AnnualNonCurrentDeferredTaxesLiabilities            []entry.Entry `json:"annualNonCurrentDeferredTaxesLiabilities,omitempty"`
	AnnualOtherShortTermInvestments                     []entry.Entry `json:"annualOtherShortTermInvestments,omitempty"`
	AnnualCashCashEquivalentsAndShortTermInvestments    []entry.Entry `json:"annualCashCashEquivalentsAndShortTermInvestments,omitempty"`
	AnnualOtherIntangibleAssets                         []entry.Entry `json:"annualOtherIntangibleAssets,omitempty"`
	AnnualTotalLiabilitiesNetMinorityInterest           []entry.Entry `json:"annualTotalLiabilitiesNetMinorityInterest,omitempty"`
	AnnualLongTermDebt                                  []entry.Entry `json:"annualLongTermDebt,omitempty"`
	AnnualCurrentDebt                                   []entry.Entry `json:"annualCurrentDebt,omitempty"`
	AnnualStockholdersEquity                            []entry.Entry `json:"annualStockholdersEquity,omitempty"`
}
