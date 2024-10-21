package controllers

import (
	"mock_cbs1/helpers/utils"
	"mock_cbs1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Account_Statement(c *gin.Context) {
	rawJson := []byte(
		`
		{
			"inquiryAccountStatementRespHeader": {
				"responseCode": 0,
				"responseDesc": "",
				"wsRefId": "98a1ac28-704a-4988-8389-0a49cec7e420"
			},
			"inquiryAccountStatementRespRecs": {
				"inquiryAccountStatementRespRec": [
					{
						"bankCode": "006",
						"bankDesc": "KRUNG THAI BANK PUBLIC COMPANY LTD.",
						"branchCode": 0,
						"branchDesc": "สาขานานาเหนือ",
						"cardNo": "",
						"chequeNo": "",
						"companyName": "นาย ทีคิวเอ ออโตเมท",
						"curCode": "THB",
						"drcrFlag": "C",
						"dueDate": "",
						"ecFlag": "0",
						"ecRevFlag": "0",
						"effDate": "",
						"exchangeRate": 1,
						"fee": 0,
						"interest": 0,
						"interestBaseRate": 0,
						"negotiatedOffsetRate": "",
						"penalty": 0,
						"principal": 50.33,
						"refInfo1": "123456789123456",
						"refInfo2": "",
						"refInfo3": "",
						"revFlag": "0",
						"runningBal": 129349.51,
						"subSequence": "",
						"tellerId": "ITBANK",
						"term": "",
						"transAmt": 50.33,
						"transCode": "PBDDP",
						"transComment": "6047-123456XXXXXX456N",
						"transDate": "20220509",
						"transDesc": "PB PAYMENT DEPOSIT C/A",
						"transLocation": "0006BIZNEXT",
						"transRefId": "105091000001",
						"transSeqNo": 318424,
						"transTime": "00:44:31",
						"witholdingTax": 0
					},
					{
						"bankCode": "006",
						"bankDesc": "KRUNG THAI BANK PUBLIC COMPANY LTD.",
						"branchCode": 0,
						"branchDesc": "สาขานานาเหนือ",
						"cardNo": "",
						"chequeNo": "",
						"companyName": "นาย ทีคิวเอ ออโตเมท",
						"curCode": "THB",
						"drcrFlag": "C",
						"dueDate": "",
						"ecFlag": "0",
						"ecRevFlag": "0",
						"effDate": "",
						"exchangeRate": 1,
						"fee": 0,
						"interest": 0,
						"interestBaseRate": 0,
						"negotiatedOffsetRate": "",
						"penalty": 0,
						"principal": 50.33,
						"refInfo1": "123456789123456",
						"refInfo2": "",
						"refInfo3": "",
						"revFlag": "0",
						"runningBal": 129399.84,
						"subSequence": "",
						"tellerId": "ITBANK",
						"term": "",
						"transAmt": 50.33,
						"transCode": "PBDDP",
						"transComment": "6047-123456XXXXXX456N",
						"transDate": "20220509",
						"transDesc": "PB PAYMENT DEPOSIT C/A",
						"transLocation": "0006BIZNEXT",
						"transRefId": "105091000002",
						"transSeqNo": 318425,
						"transTime": "01:05:25",
						"witholdingTax": 0
					},
					{
						"bankCode": "006",
						"bankDesc": "KRUNG THAI BANK PUBLIC COMPANY LTD.",
						"branchCode": 0,
						"branchDesc": "สาขานานาเหนือ",
						"cardNo": "",
						"chequeNo": "",
						"companyName": "นาย ทีคิวเอ ออโตเมท",
						"curCode": "THB",
						"drcrFlag": "C",
						"dueDate": "",
						"ecFlag": "0",
						"ecRevFlag": "0",
						"effDate": "",
						"exchangeRate": 1,
						"fee": 0,
						"interest": 0,
						"interestBaseRate": 0,
						"negotiatedOffsetRate": "",
						"penalty": 0,
						"principal": 50.33,
						"refInfo1": "123456789123456",
						"refInfo2": "",
						"refInfo3": "",
						"revFlag": "0",
						"runningBal": 129450.17,
						"subSequence": "",
						"tellerId": "ITBANK",
						"term": "",
						"transAmt": 50.33,
						"transCode": "PBDDP",
						"transComment": "6047-123456XXXXXX456N",
						"transDate": "20220509",
						"transDesc": "PB PAYMENT DEPOSIT C/A",
						"transLocation": "0006BIZNEXT",
						"transRefId": "105091000003",
						"transSeqNo": 318426,
						"transTime": "01:25:48",
						"witholdingTax": 0
					},
					{
						"bankCode": "006",
						"bankDesc": "KRUNG THAI BANK PUBLIC COMPANY LTD.",
						"branchCode": 0,
						"branchDesc": "สาขานานาเหนือ",
						"cardNo": "",
						"chequeNo": "",
						"companyName": "นาย ทีคิวเอ ออโตเมท",
						"curCode": "THB",
						"drcrFlag": "C",
						"dueDate": "",
						"ecFlag": "0",
						"ecRevFlag": "0",
						"effDate": "",
						"exchangeRate": 1,
						"fee": 0,
						"interest": 0,
						"interestBaseRate": 0,
						"negotiatedOffsetRate": "",
						"penalty": 0,
						"principal": 50.33,
						"refInfo1": "123456789123456",
						"refInfo2": "",
						"refInfo3": "",
						"revFlag": "0",
						"runningBal": 129500.5,
						"subSequence": "",
						"tellerId": "ITBANK",
						"term": "",
						"transAmt": 50.33,
						"transCode": "PBDDP",
						"transComment": "6047-123456XXXXXX456N",
						"transDate": "20220509",
						"transDesc": "PB PAYMENT DEPOSIT C/A",
						"transLocation": "0006BIZNEXT",
						"transRefId": "105091000019",
						"transSeqNo": 318427,
						"transTime": "04:54:02",
						"witholdingTax": 0
					}
				]
			},
			"lastSeq": 0,
			"moreFlag": "0",
			"totalRec": 4
		}
    `)

	var resObj models.InquiryAccountStatementResponse
	_ = utils.JsonStringToStruct(string(rawJson), &resObj)

	c.IndentedJSON(http.StatusOK, resObj)

}

func Account_payment_record(c *gin.Context) {
	rawJson := []byte(
		`
		{
			"inquiryPaymentRecordRespHeader": {
				"responseCode": 0,
				"responseDesc": "",
				"wsRefId": "8376ca91-c19d-4b96-af13-10a4adf590b7"
			},
			"lastSeq": 102,
			"totalRec": 20,
			"moreFlg": "1",
			"inquiryPaymentRecordRespRecs": {
				"inquiryPaymentRecordRespRec": [
					{
						"acctNo": "100031890682",
						"seq": 121,
						"pmtAmtDue": 7000.00,
						"pmtDueDate": "20220424",
						"totAmtBill": 7000.00,
						"pe1AmtDue": 3677.41,
						"pe2AmtDue": 3322.59,
						"pmtDist": "LNCOLLFE-L-I-P",
						"pe1AmtBill": 3677.41,
						"pe2AmtBill": 3322.59,
						"satisfiedBillDate": "",
						"collfeeAmtDue": 0.00,
						"lchgAmtDue": 0.00,
						"collfeeAmtBill": 0.00,
						"lchgAmtBill": 0.00
					},
					{
						"acctNo": "100031890682",
						"seq": 120,
						"pmtAmtDue": 7000.00,
						"pmtDueDate": "20220324",
						"totAmtBill": 7000.00,
						"pe1AmtDue": 3321.53,
						"pe2AmtDue": 3678.47,
						"pmtDist": "LNCOLLFE-L-I-P",
						"pe1AmtBill": 3321.53,
						"pe2AmtBill": 3678.47,
						"satisfiedBillDate": "",
						"collfeeAmtDue": 0.00,
						"lchgAmtDue": 0.00,
						"collfeeAmtBill": 0.00,
						"lchgAmtBill": 0.00
					},
					{
						"acctNo": "100031890682",
						"seq": 119,
						"pmtAmtDue": 7000.00,
						"pmtDueDate": "20220224",
						"totAmtBill": 7000.00,
						"pe1AmtDue": 3677.42,
						"pe2AmtDue": 3322.58,
						"pmtDist": "LNCOLLFE-L-I-P",
						"pe1AmtBill": 3677.42,
						"pe2AmtBill": 3322.58,
						"satisfiedBillDate": "",
						"collfeeAmtDue": 0.00,
						"lchgAmtDue": 0.00,
						"collfeeAmtBill": 0.00,
						"lchgAmtBill": 0.00
					},
					{
						"acctNo": "100031890682",
						"seq": 118,
						"pmtAmtDue": 7000.00,
						"pmtDueDate": "20220124",
						"totAmtBill": 7000.00,
						"pe1AmtDue": 3677.41,
						"pe2AmtDue": 3322.59,
						"pmtDist": "LNCOLLFE-L-I-P",
						"pe1AmtBill": 3677.41,
						"pe2AmtBill": 3322.59,
						"satisfiedBillDate": "",
						"collfeeAmtDue": 0.00,
						"lchgAmtDue": 0.00,
						"collfeeAmtBill": 0.00,
						"lchgAmtBill": 0.00
					},
					{
						"acctNo": "100031890682",
						"seq": 117,
						"pmtAmtDue": 7000.00,
						"pmtDueDate": "20211224",
						"totAmtBill": 7000.00,
						"pe1AmtDue": 3558.78,
						"pe2AmtDue": 3441.22,
						"pmtDist": "LNCOLLFE-L-I-P",
						"pe1AmtBill": 3558.78,
						"pe2AmtBill": 3441.22,
						"satisfiedBillDate": "",
						"collfeeAmtDue": 0.00,
						"lchgAmtDue": 0.00,
						"collfeeAmtBill": 0.00,
						"lchgAmtBill": 0.00
					},
					{
						"acctNo": "100031890682",
						"seq": 116,
						"pmtAmtDue": 3677.42,
						"pmtDueDate": "20211124",
						"totAmtBill": 3677.42,
						"pe1AmtDue": 3677.42,
						"pe2AmtDue": 0.00,
						"pmtDist": "LNCOLLFE-L-I-P",
						"pe1AmtBill": 3677.42,
						"pe2AmtBill": 0.00,
						"satisfiedBillDate": "",
						"collfeeAmtDue": 0.00,
						"lchgAmtDue": 0.00,
						"collfeeAmtBill": 0.00,
						"lchgAmtBill": 0.00
					},
					{
						"acctNo": "100031890682",
						"seq": 115,
						"pmtAmtDue": 3558.78,
						"pmtDueDate": "20211024",
						"totAmtBill": 3558.78,
						"pe1AmtDue": 3558.78,
						"pe2AmtDue": 0.00,
						"pmtDist": "LNCOLLFE-L-I-P",
						"pe1AmtBill": 3558.78,
						"pe2AmtBill": 0.00,
						"satisfiedBillDate": "",
						"collfeeAmtDue": 0.00,
						"lchgAmtDue": 0.00,
						"collfeeAmtBill": 0.00,
						"lchgAmtBill": 0.00
					},
					{
						"acctNo": "100031890682",
						"seq": 114,
						"pmtAmtDue": 1067.64,
						"pmtDueDate": "20210924",
						"totAmtBill": 1067.64,
						"pe1AmtDue": 1067.64,
						"pe2AmtDue": 0.00,
						"pmtDist": "LNCOLLFE-L-I-P",
						"pe1AmtBill": 1067.64,
						"pe2AmtBill": 0.00,
						"satisfiedBillDate": "",
						"collfeeAmtDue": 0.00,
						"lchgAmtDue": 0.00,
						"collfeeAmtBill": 0.00,
						"lchgAmtBill": 0.00
					},
					{
						"acctNo": "100031890682",
						"seq": 113,
						"pmtAmtDue": 0.00,
						"pmtDueDate": "20210824",
						"totAmtBill": 0.00,
						"pe1AmtDue": 0.00,
						"pe2AmtDue": 0.00,
						"pmtDist": "LNCOLLFE-L-I-P",
						"pe1AmtBill": 0.00,
						"pe2AmtBill": 0.00,
						"satisfiedBillDate": "20210915",
						"collfeeAmtDue": 0.00,
						"lchgAmtDue": 0.00,
						"collfeeAmtBill": 0.00,
						"lchgAmtBill": 0.00
					},
					{
						"acctNo": "100031890682",
						"seq": 112,
						"pmtAmtDue": 0.00,
						"pmtDueDate": "20210724",
						"totAmtBill": 0.00,
						"pe1AmtDue": 0.00,
						"pe2AmtDue": 0.00,
						"pmtDist": "LNCOLLFE-L-I-P",
						"pe1AmtBill": 0.00,
						"pe2AmtBill": 0.00,
						"satisfiedBillDate": "20210915",
						"collfeeAmtDue": 0.00,
						"lchgAmtDue": 0.00,
						"collfeeAmtBill": 0.00,
						"lchgAmtBill": 0.00
					},
					{
						"acctNo": "100031890682",
						"seq": 111,
						"pmtAmtDue": 0.00,
						"pmtDueDate": "20210624",
						"totAmtBill": 0.00,
						"pe1AmtDue": 0.00,
						"pe2AmtDue": 0.00,
						"pmtDist": "LNCOLLFE-L-I-P",
						"pe1AmtBill": 0.00,
						"pe2AmtBill": 0.00,
						"satisfiedBillDate": "20210915",
						"collfeeAmtDue": 0.00,
						"lchgAmtDue": 0.00,
						"collfeeAmtBill": 0.00,
						"lchgAmtBill": 0.00
					},
					{
						"acctNo": "100031890682",
						"seq": 110,
						"pmtAmtDue": 0.00,
						"pmtDueDate": "20210524",
						"totAmtBill": 0.00,
						"pe1AmtDue": 0.00,
						"pe2AmtDue": 0.00,
						"pmtDist": "LNCOLLFE-L-I-P",
						"pe1AmtBill": 0.00,
						"pe2AmtBill": 0.00,
						"satisfiedBillDate": "20210915",
						"collfeeAmtDue": 0.00,
						"lchgAmtDue": 0.00,
						"collfeeAmtBill": 0.00,
						"lchgAmtBill": 0.00
					},
					{
						"acctNo": "100031890682",
						"seq": 109,
						"pmtAmtDue": 0.00,
						"pmtDueDate": "20210424",
						"totAmtBill": 0.00,
						"pe1AmtDue": 0.00,
						"pe2AmtDue": 0.00,
						"pmtDist": "LNCOLLFE-L-I-P",
						"pe1AmtBill": 0.00,
						"pe2AmtBill": 0.00,
						"satisfiedBillDate": "20210915",
						"collfeeAmtDue": 0.00,
						"lchgAmtDue": 0.00,
						"collfeeAmtBill": 0.00,
						"lchgAmtBill": 0.00
					},
					{
						"acctNo": "100031890682",
						"seq": 108,
						"pmtAmtDue": 0.00,
						"pmtDueDate": "20210324",
						"totAmtBill": 0.00,
						"pe1AmtDue": 0.00,
						"pe2AmtDue": 0.00,
						"pmtDist": "LNCOLLFE-L-I-P",
						"pe1AmtBill": 0.00,
						"pe2AmtBill": 0.00,
						"satisfiedBillDate": "20210915",
						"collfeeAmtDue": 0.00,
						"lchgAmtDue": 0.00,
						"collfeeAmtBill": 0.00,
						"lchgAmtBill": 0.00
					},
					{
						"acctNo": "100031890682",
						"seq": 107,
						"pmtAmtDue": 0.00,
						"pmtDueDate": "20210224",
						"totAmtBill": 0.00,
						"pe1AmtDue": 0.00,
						"pe2AmtDue": 0.00,
						"pmtDist": "LNCOLLFE-L-I-P",
						"pe1AmtBill": 0.00,
						"pe2AmtBill": 0.00,
						"satisfiedBillDate": "20210915",
						"collfeeAmtDue": 0.00,
						"lchgAmtDue": 0.00,
						"collfeeAmtBill": 0.00,
						"lchgAmtBill": 0.00
					},
					{
						"acctNo": "100031890682",
						"seq": 106,
						"pmtAmtDue": 0.00,
						"pmtDueDate": "20210124",
						"totAmtBill": 0.00,
						"pe1AmtDue": 0.00,
						"pe2AmtDue": 0.00,
						"pmtDist": "LNCOLLFE-L-I-P",
						"pe1AmtBill": 0.00,
						"pe2AmtBill": 0.00,
						"satisfiedBillDate": "20210915",
						"collfeeAmtDue": 0.00,
						"lchgAmtDue": 0.00,
						"collfeeAmtBill": 0.00,
						"lchgAmtBill": 0.00
					},
					{
						"acctNo": "100031890682",
						"seq": 105,
						"pmtAmtDue": 0.00,
						"pmtDueDate": "20201224",
						"totAmtBill": 0.00,
						"pe1AmtDue": 0.00,
						"pe2AmtDue": 0.00,
						"pmtDist": "LNCOLLFE-L-I-P",
						"pe1AmtBill": 0.00,
						"pe2AmtBill": 0.00,
						"satisfiedBillDate": "20210915",
						"collfeeAmtDue": 0.00,
						"lchgAmtDue": 0.00,
						"collfeeAmtBill": 0.00,
						"lchgAmtBill": 0.00
					},
					{
						"acctNo": "100031890682",
						"seq": 104,
						"pmtAmtDue": 0.00,
						"pmtDueDate": "20201124",
						"totAmtBill": 0.00,
						"pe1AmtDue": 0.00,
						"pe2AmtDue": 0.00,
						"pmtDist": "LNCOLLFE-L-I-P",
						"pe1AmtBill": 0.00,
						"pe2AmtBill": 0.00,
						"satisfiedBillDate": "20210915",
						"collfeeAmtDue": 0.00,
						"lchgAmtDue": 0.00,
						"collfeeAmtBill": 0.00,
						"lchgAmtBill": 0.00
					},
					{
						"acctNo": "100031890682",
						"seq": 103,
						"pmtAmtDue": 0.00,
						"pmtDueDate": "20200924",
						"totAmtBill": 0.00,
						"pe1AmtDue": 0.00,
						"pe2AmtDue": 0.00,
						"pmtDist": "LNCOLLFE-L-I-P",
						"pe1AmtBill": 0.00,
						"pe2AmtBill": 0.00,
						"satisfiedBillDate": "20210915",
						"collfeeAmtDue": 0.00,
						"lchgAmtDue": 0.00,
						"collfeeAmtBill": 0.00,
						"lchgAmtBill": 0.00
					},
					{
						"acctNo": "100031890682",
						"seq": 102,
						"pmtAmtDue": 0.00,
						"pmtDueDate": "20200824",
						"totAmtBill": 0.00,
						"pe1AmtDue": 0.00,
						"pe2AmtDue": 0.00,
						"pmtDist": "LNCOLLFE-L-I-P",
						"pe1AmtBill": 0.00,
						"pe2AmtBill": 0.00,
						"satisfiedBillDate": "20210915",
						"collfeeAmtDue": 0.00,
						"lchgAmtDue": 0.00,
						"collfeeAmtBill": 0.00,
						"lchgAmtBill": 0.00
					}
				]
			}
		}
    `)

	var resObj models.InquiryPaymentRecordResponse
	_ = utils.JsonStringToStruct(string(rawJson), &resObj)
	c.IndentedJSON(http.StatusOK, resObj)

}

func InquiryCBSData(c *gin.Context) {
	rawJson := []byte(
		`
	   	{
			"inquiryCBSDataRespHeader": {
				"responseCode": 0,
				"responseDesc": "",
				"wsRefId": "c79fbbfc-ed2e-45af-b9b8-a892614bf93a"
			},
			"tblKeyValue": "5016101973",
			"tblKeyRespDatas": {
				"tblKeyRespData": [
					{
						"tblKeyName": "LASTPAYAMT",
						"tblKeyValue": ""
					},
					{
						"tblKeyName": "BALANCE",
						"tblKeyValue": "-4113431.58"
					},
					{
						"tblKeyName": "AVLBALANCE",
						"tblKeyValue": "986568.42"
					},
					{
						"tblKeyName": "CHECKHOLDAMT",
						"tblKeyValue": "0"
					},
					{
						"tblKeyName": "PRODUCTGROUP",
						"tblKeyValue": "DDA"
					},
					{
						"tblKeyName": "HOLDAMT",
						"tblKeyValue": "0"
					},
					{
						"tblKeyName": "NEGACRAUTH",
						"tblKeyValue": "1286095.3744"
					},
					{
						"tblKeyName": "NEGACRUNAUTH",
						"tblKeyValue": "3989150.80284"
					},
					{
						"tblKeyName": "AUTHLIMIT",
						"tblKeyValue": "5100000"
					},
					{
						"tblKeyName": "LIMITAMT",
						"tblKeyValue": "5100000"
					},
					{
						"tblKeyName": "ACCOUNTSTATUS",
						"tblKeyValue": "5"
					}
				]
			}
		}	
    `)

	var resObj models.CBSDataRespone
	_ = utils.JsonStringToStruct(string(rawJson), &resObj)

	c.IndentedJSON(http.StatusOK, resObj)

}
