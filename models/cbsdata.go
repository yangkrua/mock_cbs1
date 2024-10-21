package models

//CBSDataReqHeader
type CBSDataReqHeader struct {
	ChannelId   string `json:"channelId"`
	RequestDate string `json:"requestDate"`
	RequestTime string `json:"requestTime"`
	RefId       string `json:"refId"`
}

//RespHeader
type RespHeader struct {
	ResponseCode int    `json:"responseCode"`
	ResponseDesc string `json:"responseDesc"`
	WsRefId      string `json:"wsRefId"`
}

//TblFieldData
type TblFieldData struct {
	TblFieldName  string `json:"tblFieldName"`
	TblFieldValue string `json:"tblFieldValue"`
}

//TblFieldDatas
type TblFieldDatas struct {
	TblFieldData []*TblFieldData `json:"tblFieldData"`
}

//TblKeyData
type TblKeyData struct {
	TblKeyName  string `json:"tblKeyName"`
	TblKeyValue string `json:"tblKeyValue"`
}

//TblKeyDatas
type TblKeyDatas struct {
	TblKeyData []*TblKeyData `json:"tblFieldData"`
}

//TblKeyRespDatas
type TblKeyRespDatas struct {
	TblKeyRespData []*TblKeyData `json:"tblKeyRespData"`
}

//CBSDataReq
type CBSDataReq struct {
	CifNo          string            `json:"cifNo"`
	AcctNo         string            `json:"acctNo"`
	Header         *CBSDataReqHeader `json:"inquiryCBSDataReqHeader"`
	InquiryTblFlag string            `json:"inquiryTblFlag"`
	TableName      string            `json:"tableName"`
	TblFieldDatas  *TblFieldDatas    `json:"tblFieldDatas"`
	TblKeyDatas    *TblKeyDatas      `json:"tblKeyDatas"`
	TotFieldCnt    int               `json:"TotFieldCnt"`
}

//CBSDataRespone
type CBSDataRespone struct {
	Header          *RespHeader      `json:"inquiryCBSDataRespHeader"`
	InquiryTblFlag  string           `json:"inquiryTblFlag"`
	TblKeyValue     string           `json:"tblKeyValue"`
	TblKeyRespDatas *TblKeyRespDatas `json:"tblKeyRespDatas"`
}

//CBSData
type CBSData struct {
	ACCOUNT_NO     string `json:"account_no"`
	Dailyacrint    string `json:"dailyacrint"`
	Authlimit      string `json:"authlimit"`
	Balance        string `json:"balance"`
	Payoffamt      string `json:"payoffamt"`
	Paycalmethod   string `json:"paycalmethod"`
	Pipayment      string `json:"pipayment"`
	PRINPASSDUE    string `json:"prinpassdue"`
	INTPASSDUE     string `json:"intpassdue"`
	Miscchgamt     string `json:"miscchgamt"`
	Lchgamt        string `json:"lchgamt"`
	Lchgamtb1jul21 string `json:"lchgamtb1jul21"`
	Totalpayment   string `json:"totalpayment"`
	Avlbalance     string `json:"avlbalance"`
	Limitamt       string `json:"limitamt"`
	Negacrauth     string `json:"negacrauth"`
	Negacrunauth   string `json:"negacrunauth"`
	Holdamt        string `json:"holdamt"`
	Lastpaydate    string `json:"lastpaydate"`
	Lastpayamt     string `json:"lastpayamt"`
	Productgroup   string `json:"productgroup"`
	Nextduedate    string `json:"nextduedate"`
	Checkholdamt   string `json:"checkholdamt"`
	AccuntStatus   string `json:"accuntstatus"`
	PaidAccount    string `json:"paidaccount"`
	LedgerBal      string `json:"ledgerbal"`
}

//CHECK_CUSTOMER_PAID_ACCOUNT_RESPONSE
type CHECK_CUSTOMER_PAID_ACCOUNT_RESPONSE struct {
	RespHeader    *RespHeader `json:"resp_header"`
	CIF_NO        string      `json:"cif_no"`
	FOLLOW_STATUS string      `json:"follow_status"`
	PaidAccount   string      `json:"paid_account"`
	ACCOUNTS      []CBSData   `json:"accounts"`
}
