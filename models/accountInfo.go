package models

//AccountInfoResponse
type AccountInfoResponse struct {
	RespCode               string    `json:"respCode"`
	RespDesc               string    `json:"respDesc"`
	Loan_limit             string    `json:"loan_limit"`
	Outstanding_balance    string    `json:"outstanding_balance"`
	Pay_off_amount         string    `json:"pay_off_amount"`
	P_i_payment_amt        string    `json:"p_i_payment_amt"`
	Principle_amt          string    `json:"principle_amt"`
	Interest_amt           string    `json:"interest_amt"`
	Misc_charge_amt        string    `json:"misc_charge_amt"`
	Late_charge_amt        string    `json:"late_charge_amt"`
	Total_due_amount       string    `json:"total_due_amount"`
	Daily_accrual_interest string    `json:"daily_accrual_interest"`
	Last_payment_date      string    `json:"last_payment_date"`
	Last_payment_amount    string    `json:"last_payment_amount"`
	Payment                []Payment `json:"paymentreccord"`
}

//Payment
type Payment struct {
	No                    string `json:"no"`
	Due_date              string `json:"due_date"`
	Total                 string `json:"total"`
	Unpaid                string `json:"unpaid"`
	Unpaid_collection_fee string `json:"unpaid_collection_fee"`
	Unpaid_late_charge    string `json:"unpaid_late_charge"`
	Unpaid_interest       string `json:"unpaid_interest"`
	Unpaid_principal      string `json:"unpaid_principal"`
}
