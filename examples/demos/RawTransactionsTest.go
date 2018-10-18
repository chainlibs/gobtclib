package demos

import (
	"github.com/gobasis/log"

)

///*
//Description:
//A demo test of CombinePSBT.
//* Author: architect.bian
//* Date: 2018/10/15 12:14
//*/
//func CombinePSBTTest() {
//	result, err := cli.CombinePSBT()
//	if err != nil {
//		log.Fatal("", "error", err)
//	}
//	log.Info("CombinePSBT", "result", result)
//}
//
///*
//Description:
//A demo test of CombineRawTransaction.
//* Author: architect.bian
//* Date: 2018/10/15 12:14
//*/
//func CombineRawTransactionTest() {
//	result, err := cli.CombineRawTransaction()
//	if err != nil {
//		log.Fatal("", "error", err)
//	}
//	log.Info("CombineRawTransaction", "result", result)
//}
//
///*
//Description:
//A demo test of ConvertTopSBT.
//* Author: architect.bian
//* Date: 2018/10/15 12:14
//*/
//func ConvertTopSBTTest() {
//	result, err := cli.ConvertTopSBT()
//	if err != nil {
//		log.Fatal("", "error", err)
//	}
//	log.Info("ConvertTopSBT", "result", result)
//}
//
///*
//Description:
//A demo test of CreatepSBT.
//* Author: architect.bian
//* Date: 2018/10/15 12:14
//*/
//func CreatepSBTTest() {
//	result, err := cli.CreatepSBT()
//	if err != nil {
//		log.Fatal("", "error", err)
//	}
//	log.Info("CreatepSBT", "result", result)
//}
//
///*
//Description:
//A demo test of CreateRawTransaction.
//* Author: architect.bian
//* Date: 2018/10/15 12:14
//*/
//func CreateRawTransactionTest() {
//	result, err := cli.CreateRawTransaction()
//	if err != nil {
//		log.Fatal("", "error", err)
//	}
//	log.Info("CreateRawTransaction", "result", result)
//}
//
///*
//Description:
//A demo test of DecodepSBT.
//* Author: architect.bian
//* Date: 2018/10/15 12:14
//*/
//func DecodepSBTTest() {
//	result, err := cli.DecodepSBT()
//	if err != nil {
//		log.Fatal("", "error", err)
//	}
//	log.Info("DecodepSBT", "result", result)
//}
//
///*
//Description:
//A demo test of DecodeRawTransaction.
//* Author: architect.bian
//* Date: 2018/10/15 12:14
//*/
//func DecodeRawTransactionTest() {
//	result, err := cli.DecodeRawTransaction()
//	if err != nil {
//		log.Fatal("", "error", err)
//	}
//	log.Info("DecodeRawTransaction", "result", result)
//}

/*
Description:
A demo test of DecodeScript.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func DecodeScriptTest() {
	result, err := cli.DecodeScript("21035be9cb700c1b2309b8c5e345efb7dacb6b183817c1b5b3f4408948a1779d6543ac")
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("DecodeScript", "result", result)
	log.Info("DecodeScript", "result.addresses", ((*result).(map[string]interface{}))["addresses"])
}

///*
//Description:
//A demo test of FinalizepSBT.
//* Author: architect.bian
//* Date: 2018/10/15 12:14
//*/
//func FinalizepSBTTest() {
//	result, err := cli.FinalizepSBT()
//	if err != nil {
//		log.Fatal("", "error", err)
//	}
//	log.Info("FinalizepSBT", "result", result)
//}
//
///*
//Description:
//A demo test of FundrawTransaction.
//* Author: architect.bian
//* Date: 2018/10/15 12:14
//*/
//func FundrawTransactionTest() {
//	result, err := cli.FundrawTransaction()
//	if err != nil {
//		log.Fatal("", "error", err)
//	}
//	log.Info("FundrawTransaction", "result", result)
//}
//
///*
//Description:
//A demo test of GetRawTransaction.
//* Author: architect.bian
//* Date: 2018/10/15 12:14
//*/
//func GetRawTransactionTest() {
//	result, err := cli.GetRawTransaction()
//	if err != nil {
//		log.Fatal("", "error", err)
//	}
//	log.Info("GetRawTransaction", "result", result)
//}
//
///*
//Description:
//A demo test of SendRawTransaction.
//* Author: architect.bian
//* Date: 2018/10/15 12:14
//*/
//func SendRawTransactionTest() {
//	result, err := cli.SendRawTransaction()
//	if err != nil {
//		log.Fatal("", "error", err)
//	}
//	log.Info("SendRawTransaction", "result", result)
//}
//
///*
//Description:
//A demo test of SignRawTransaction.
//* Author: architect.bian
//* Date: 2018/10/15 12:14
//*/
//func SignRawTransactionTest() {
//	result, err := cli.SignRawTransaction()
//	if err != nil {
//		log.Fatal("", "error", err)
//	}
//	log.Info("SignRawTransaction", "result", result)
//}
//
///*
//Description:
//A demo test of SignRawTransactionWithKey.
//* Author: architect.bian
//* Date: 2018/10/15 12:14
//*/
//func SignRawTransactionWithKeyTest() {
//	result, err := cli.SignRawTransactionWithKey()
//	if err != nil {
//		log.Fatal("", "error", err)
//	}
//	log.Info("SignRawTransactionWithKey", "result", result)
//}
//
///*
//Description:
//A demo test of TestMempoolAccept.
//* Author: architect.bian
//* Date: 2018/10/15 12:14
//*/
//func TestMempoolAcceptTest() {
//	result, err := cli.TestMempoolAccept()
//	if err != nil {
//		log.Fatal("", "error", err)
//	}
//	log.Info("TestMempoolAccept", "result", result)
//}