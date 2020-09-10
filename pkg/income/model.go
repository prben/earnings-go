package income

import "github.com/prben/earnings-go/pkg/entry"

// Income struct
// AnnualGrossProfit = AnnualTotalRevenue - AnnualCostOfRevenue
// AnnualOperatingIncome = AnnualGrossProfit - AnnualOperatingExpense
// AnnualNetIncome = AnnualOperatingIncome - TAXES
type IncomeResponse struct {
	Timeseries struct {
		Result []Income `json:"result"`
	} `json:"timeseries"`
}

type Income struct {
	AnnualTotalRevenue                    []entry.Entry `json:"annualTotalRevenue,omitempty"`
	TrailingNetIncomeCommonStockholders   []entry.Entry `json:"trailingNetIncomeCommonStockholders,omitempty"`
	TrailingOperatingIncome               []entry.Entry `json:"trailingOperatingIncome,omitempty"`
	AnnualCostOfRevenue                   []entry.Entry `json:"annualCostOfRevenue,omitempty"`
	AnnualNetIncomeContinuousOperations   []entry.Entry `json:"annualNetIncomeContinuousOperations,omitempty"`
	TrailingNetIncomeContinuousOperations []entry.Entry `json:"trailingNetIncomeContinuousOperations,omitempty"`
	AnnualNetIncomeCommonStockholders     []entry.Entry `json:"annualNetIncomeCommonStockholders,omitempty"`
	AnnualEbitda                          []entry.Entry `json:"annualEbitda,omitempty"`
	TrailingGrossProfit                   []entry.Entry `json:"trailingGrossProfit,omitempty"`
	AnnualOperatingExpense                []entry.Entry `json:"annualOperatingExpense,omitempty"`
	AnnualSellingGeneralAndAdministration []entry.Entry `json:"annualSellingGeneralAndAdministration,omitempty"`
	AnnualGrossProfit                     []entry.Entry `json:"annualGrossProfit,omitempty"`
	AnnualNetIncome                       []entry.Entry `json:"annualNetIncome,omitempty"`
	TrailingTaxProvision                  []entry.Entry `json:"trailingTaxProvision,omitempty"`
	AnnualBasicEPS                        []entry.Entry `json:"annualBasicEPS,omitempty"`
	TrailingInterestExpense               []entry.Entry `json:"trailingInterestExpense,omitempty"`
	AnnualDilutedEPS                      []entry.Entry `json:"annualDilutedEPS,omitempty"`
	TrailingPretaxIncome                  []entry.Entry `json:"trailingPretaxIncome,omitempty"`
	AnnualTaxProvision                    []entry.Entry `json:"annualTaxProvision,omitempty"`
	AnnualDilutedAverageShares            []entry.Entry `json:"annualDilutedAverageShares,omitempty"`
	AnnualInterestExpense                 []entry.Entry `json:"annualInterestExpense,omitempty"`
	TrailingOperatingExpense              []entry.Entry `json:"trailingOperatingExpense,omitempty"`
	AnnualOperatingIncome                 []entry.Entry `json:"annualOperatingIncome,omitempty"`
	AnnualPretaxIncome                    []entry.Entry `json:"annualPretaxIncome,omitempty"`
	TrailingTotalRevenue                  []entry.Entry `json:"trailingTotalRevenue,omitempty"`
	TrailingCostOfRevenue                 []entry.Entry `json:"trailingCostOfRevenue,omitempty"`
	TrailingNetIncome                     []entry.Entry `json:"trailingNetIncome,omitempty"`
	AnnualBasicAverageShares              []entry.Entry `json:"annualBasicAverageShares,omitempty"`
}
