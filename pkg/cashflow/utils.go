package cashflow

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

var (
	endPoint   = "query1.finance.yahoo.com"
	corsDomain = "in.finance.yahoo.com"
	period1    = time.Now().AddDate(-10, 0, 0).Unix()
	period2    = time.Now().Unix()
)

func genURL(symbol string) string {

	return "https://" + endPoint + "/ws/fundamentals-timeseries/v1/finance/timeseries/" + symbol + "?padTimeSeries=true&type=annualFreeCashFlow%2CtrailingFreeCashFlow%2CannualCapitalExpenditure%2CtrailingCapitalExpenditure%2CannualOperatingCashFlow%2CtrailingOperatingCashFlow%2CannualEndCashPosition%2CtrailingEndCashPosition%2CannualBeginningCashPosition%2CtrailingBeginningCashPosition%2CannualChangeInCashSupplementalAsReported%2CtrailingChangeInCashSupplementalAsReported%2CannualCashFlowFromContinuingFinancingActivities%2CtrailingCashFlowFromContinuingFinancingActivities%2CannualNetOtherFinancingCharges%2CtrailingNetOtherFinancingCharges%2CannualCashDividendsPaid%2CtrailingCashDividendsPaid%2CannualRepurchaseOfCapitalStock%2CtrailingRepurchaseOfCapitalStock%2CannualCommonStockIssuance%2CtrailingCommonStockIssuance%2CannualRepaymentOfDebt%2CtrailingRepaymentOfDebt%2CannualInvestingCashFlow%2CtrailingInvestingCashFlow%2CannualNetOtherInvestingChanges%2CtrailingNetOtherInvestingChanges%2CannualSaleOfInvestment%2CtrailingSaleOfInvestment%2CannualPurchaseOfInvestment%2CtrailingPurchaseOfInvestment%2CannualPurchaseOfBusiness%2CtrailingPurchaseOfBusiness%2CannualOtherNonCashItems%2CtrailingOtherNonCashItems%2CannualChangeInAccountPayable%2CtrailingChangeInAccountPayable%2CannualChangeInInventory%2CtrailingChangeInInventory%2CannualChangesInAccountReceivables%2CtrailingChangesInAccountReceivables%2CannualChangeInWorkingCapital%2CtrailingChangeInWorkingCapital%2CannualStockBasedCompensation%2CtrailingStockBasedCompensation%2CannualDeferredIncomeTax%2CtrailingDeferredIncomeTax%2CannualDepreciationAndAmortization%2CtrailingDepreciationAndAmortization%2CannualNetIncome%2CtrailingNetIncome&merge=false&period1=" + strconv.FormatInt(period1, 10) + "&period2=" + strconv.FormatInt(period2, 10) + "&corsDomain=" + corsDomain

}

// NewCashflow get incomesheet for the symbol
func NewCashflow(Symbol string) (*Cashflow, error) {

	resp, err := http.Get(genURL(Symbol))
	if err != nil {
		return nil, err
	}

	obj, err := decodeHTTPresponse(resp.Body)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

// decodeHTTPresponse decode http resp.body
func decodeHTTPresponse(rawData io.ReadCloser) (*Cashflow, error) {

	data, err := ioutil.ReadAll(rawData)
	if err != nil {
		return nil, err
	}

	Resp := new(CashflowResponse)
	if err := json.Unmarshal(data, Resp); err != nil {
		return nil, err
	}

	return concatenate(Resp.Timeseries.Result)
}

// ConcatenateIncome merge all incomes
func concatenate(arr []Cashflow) (*Cashflow, error) {

	newObj := new(Cashflow)
	newObjValue := reflect.Indirect(reflect.ValueOf(newObj))

	for _, item := range arr {

		s := reflect.ValueOf(item)
		typeOf := reflect.TypeOf(item)

		for i := 0; i < s.NumField(); i++ {
			if !s.Field(i).IsNil() {
				if newObjValue.FieldByName(typeOf.Field(i).Name).CanSet() {
					newObjValue.FieldByName(typeOf.Field(i).Name).Set(s.Field(i))
				}
			}
		}
	}

	return newObj, nil
}
