package models

//InquiryPaymentRecordRequest
type InquiryPaymentRecordRequest struct {
	Header      *CBSDataReqHeader `json:"inquiryPaymentRecordReqHeader"`
	CifNo       string            `json:"cifNo"`
	AcctNo      string            `json:"acctNo"`
	LastSeq     int64             `json:"lastSeq"`
	OrderStatus string            `json:"orderStatus"`
}

//InquiryPaymentRecordRespRec
type InquiryPaymentRecordRespRec struct {
	AcctNo            string  `json:"acctNo"`
	Seq               int64   `json:"seq"`
	PmtAmtDue         float64 `json:"pmtAmtDue"`
	PmtDueDate        string  `json:"pmtDueDate"`
	TotAmtBill        float64 `json:"totAmtBill"`
	Pe1AmtDue         float64 `json:"pe1AmtDue"`
	Pe2AmtDue         float64 `json:"pe2AmtDue"`
	PmtDist           string  `json:"pmtDist"`
	Pe1AmtBill        float64 `json:"pe1AmtBill"`
	Pe2AmtBill        float64 `json:"pe2AmtBill"`
	SatisfiedBillDate string  `json:"satisfiedBillDate"`
	CollfeeAmtDue     float64 `json:"collfeeAmtDue"`
	LchgAmtDue        float64 `json:"lchgAmtDue"`
	CollfeeAmtBill    float64 `json:"collfeeAmtBill"`
	LchgAmtBill       float64 `json:"lchgAmtBill"`
}

//InquiryPaymentRecordRespRecs
type InquiryPaymentRecordRespRecs struct {
	PaymentRecord []*InquiryPaymentRecordRespRec `json:"inquiryPaymentRecordRespRec"`
}

//InquiryPaymentRecordResponse
type InquiryPaymentRecordResponse struct {
	Header         *RespHeader                   `json:"inquiryPaymentRecordRespHeader"`
	LastSeq        int64                         `json:"lastSeq"`
	MoreFlg        string                        `json:"moreFlg"`
	TotalRec       int                           `json:"totalRec"`
	PaymentRecords *InquiryPaymentRecordRespRecs `json:"inquiryPaymentRecordRespRecs"`
}
