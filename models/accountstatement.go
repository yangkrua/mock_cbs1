package models

//AccountStatementRequest
type AccountStatementRequest struct {
	Header        *CBSDataReqHeader `json:"inquiryAccountStatementReqHeader"`
	CifNo         string            `json:"cifNo"`
	AccountNo     string            `json:"accountNo"`
	StartDate     string            `json:"startDate"`
	EndDate       string            `json:"endDate"`
	StartTransAmt float64           `json:"startTransAmt"`
	EndTransAmt   float64           `json:"endTransAmt"`
	LastSeq       int               `json:"lastSeq"`
}

//AccountStatementItem
type AccountStatementItem struct {
	TransCode            string  `json:"transCode"`
	TransDesc            string  `json:"transDesc"`
	TransLocation        string  `json:"transLocation"`
	BankCode             string  `json:"bankCode"`
	BankDesc             string  `json:"bankDesc"`
	BranchCode           int64   `json:"branchCode"`
	BranchDesc           string  `json:"branchDesc"`
	TellerId             string  `json:"tellerId"`
	TransDate            string  `json:"transDate"`
	TransTime            string  `json:"transTime"`
	EffDate              string  `json:"effDate"`
	CurCode              string  `json:"curCode"`
	TransAmt             float64 `json:"transAmt"`
	RunningBal           float64 `json:"runningBal"`
	TransSeqNo           int64   `json:"transSeqNo"`
	ChequeNo             string  `json:"chequeNo"`
	TransComment         string  `json:"transComment"`
	CompanyName          string  `json:"companyName"`
	RefInfo1             string  `json:"refInfo1"`
	RefInfo2             string  `json:"refInfo2"`
	RefInfo3             string  `json:"refInfo3"`
	DrcrFlag             string  `json:"drcrFlag"`
	EcFlag               string  `json:"ecFlag"`
	RevFlag              string  `json:"revFlag"`
	EcRevFlag            string  `json:"ecRevFlag"`
	Term                 string  `json:"term"`
	NegotiatedOffsetRate string  `json:"negotiatedOffsetRate"`
	InterestBaseRate     float64 `json:"interestBaseRate"`
	SubSequence          string  `json:"subSequence"`
	WitholdingTax        float64 `json:"witholdingTax"`
	DueDate              string  `json:"dueDate"`
	Fee                  float64 `json:"fee"`
	Penalty              float64 `json:"penalty"`
	Interest             float64 `json:"interest"`
	Principal            float64 `json:"principal"`
	ExchangeRate         float64 `json:"exchangeRate"`
	TransRefId           string  `json:"transRefId"`
	CardNo               string  `json:"cardNo"`
}

//AccountStatementRespRecs
type AccountStatementRespRecs struct {
	AccountStatementItem []*AccountStatementItem `json:"inquiryAccountStatementRespRec"`
}

//InquiryAccountStatementResponse
type InquiryAccountStatementResponse struct {
	Header                  *RespHeader               `json:"inquiryAccountStatementRespHeader"`
	AccountStatementRecords *AccountStatementRespRecs `json:"inquiryAccountStatementRespRecs"`
	LastSeq                 int                       `json:"lastSeq"`
	MoreFlag                string                    `json:"moreFlag"`
	TotalRec                int                       `json:"totalRec"`
}

//AccountStatement
type AccountStatement struct {
	ACCOUNT_NO    string  `json:"ACCOUNT_NO"`
	TRANSDATE     string  `json:"TRANSDATE"`
	TRANSTIME     string  `json:"TRANSTIME"`
	TRANSSEQNO    string  `json:"TRANSSEQNO"`
	TRANSLOCATION string  `json:"TRANSLOCATION"`
	TRANSAMT      float64 `json:"TRANSAMT"`
	RUNNINGBAL    float64 `json:"RUNNINGBAL"`
	BRANCHCODE    string  `json:"BRANCHCODE"`
	TELLERID      string  `json:"TELLERID"`
}
