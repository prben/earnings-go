package balancesheet

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

	return "https://" + endPoint + "/ws/fundamentals-timeseries/v1/finance/timeseries/" + symbol + "?padTimeSeries=true&type=annualTotalAssets%2CtrailingTotalAssets%2CannualStockholdersEquity%2CtrailingStockholdersEquity%2CannualGainsLossesNotAffectingRetainedEarnings%2CtrailingGainsLossesNotAffectingRetainedEarnings%2CannualRetainedEarnings%2CtrailingRetainedEarnings%2CannualCapitalStock%2CtrailingCapitalStock%2CannualTotalLiabilitiesNetMinorityInterest%2CtrailingTotalLiabilitiesNetMinorityInterest%2CannualTotalNonCurrentLiabilitiesNetMinorityInterest%2CtrailingTotalNonCurrentLiabilitiesNetMinorityInterest%2CannualOtherNonCurrentLiabilities%2CtrailingOtherNonCurrentLiabilities%2CannualNonCurrentDeferredRevenue%2CtrailingNonCurrentDeferredRevenue%2CannualNonCurrentDeferredTaxesLiabilities%2CtrailingNonCurrentDeferredTaxesLiabilities%2CannualLongTermDebt%2CtrailingLongTermDebt%2CannualCurrentLiabilities%2CtrailingCurrentLiabilities%2CannualOtherCurrentLiabilities%2CtrailingOtherCurrentLiabilities%2CannualCurrentDeferredRevenue%2CtrailingCurrentDeferredRevenue%2CannualCurrentAccruedExpenses%2CtrailingCurrentAccruedExpenses%2CannualIncomeTaxPayable%2CtrailingIncomeTaxPayable%2CannualAccountsPayable%2CtrailingAccountsPayable%2CannualCurrentDebt%2CtrailingCurrentDebt%2CannualTotalNonCurrentAssets%2CtrailingTotalNonCurrentAssets%2CannualOtherNonCurrentAssets%2CtrailingOtherNonCurrentAssets%2CannualOtherIntangibleAssets%2CtrailingOtherIntangibleAssets%2CannualGoodwill%2CtrailingGoodwill%2CannualInvestmentsAndAdvances%2CtrailingInvestmentsAndAdvances%2CannualNetPPE%2CtrailingNetPPE%2CannualAccumulatedDepreciation%2CtrailingAccumulatedDepreciation%2CannualGrossPPE%2CtrailingGrossPPE%2CannualCurrentAssets%2CtrailingCurrentAssets%2CannualOtherCurrentAssets%2CtrailingOtherCurrentAssets%2CannualInventory%2CtrailingInventory%2CannualAccountsReceivable%2CtrailingAccountsReceivable%2CannualCashCashEquivalentsAndShortTermInvestments%2CtrailingCashCashEquivalentsAndShortTermInvestments%2CannualOtherShortTermInvestments%2CtrailingOtherShortTermInvestments%2CannualCashAndCashEquivalents%2CtrailingCashAndCashEquivalents&merge=false&period1=" + strconv.FormatInt(period1, 10) + "&period2=" + strconv.FormatInt(period2, 10) + "&corsDomain=" + corsDomain

}

// NewBalanceSheet get incomesheet for the symbol
func NewBalanceSheet(Symbol string) (*BalanceSheet, error) {

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
func decodeHTTPresponse(rawData io.ReadCloser) (*BalanceSheet, error) {

	data, err := ioutil.ReadAll(rawData)
	if err != nil {
		return nil, err
	}

	Resp := new(BalanceSheetResponse)
	if err := json.Unmarshal(data, Resp); err != nil {
		return nil, err
	}

	return concatenate(Resp.Timeseries.Result)
}

// Concatenate merge all incomes
func concatenate(arr []BalanceSheet) (*BalanceSheet, error) {

	newObj := new(BalanceSheet)
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
