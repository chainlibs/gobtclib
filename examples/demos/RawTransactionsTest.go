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

/*
Description:
A demo test of CombineRawTransaction.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func CombineRawTransactionTest() {
	result, err := cli.CombineRawTransaction([]string{"0200000001731c3007e201b1e49ac79b99994453df6df152d8da082b4025ee79ab57cc26d30000000000fdffffff02605af405000000001976a9147872b1b6a3593899ce9eb25160ad3edd92d4138c88ac00a3e1110000000017a914a4c5fafd4ca4013ce0a79a6e7f3517c3e18432728700000000"})
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("CombineRawTransaction", "result", result)
}

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

/*
Description:
A demo test of CreateRawTransaction.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func CreateRawTransactionTest() {
	inputs := make([]map[string]interface{}, 0)
	vin := make(map[string]interface{})
	vin["txid"] = "d326cc57ab79ee25402b08dad852f16ddf534499999bc79ae4b101e207301c73"
	vin["vout"] = 0
	//vin["sequence"] = 4294967295
	inputs = append(inputs, vin)
	outputs := make([]map[string]interface{}, 0)
	out := make(map[string]interface{})
	out[addressTo] = 0.9999
	outputs = append(outputs, out)
	out = make(map[string]interface{})
	out[addressFrom1] = 3
	outputs = append(outputs, out)
	result, err := cli.CreateRawTransaction(inputs, outputs, 0, true)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("CreateRawTransaction", "result", result)
	//result: 020000000001f0b9f505000000001976a9147872b1b6a3593899ce9eb25160ad3edd92d4138c88ac00000000
}

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

/*
Description:
A demo test of DecodeRawTransaction.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func DecodeRawTransactionTest() {
	//hexStr, err := cli.GetRawTransaction("d326cc57ab79ee25402b08dad852f16ddf534499999bc79ae4b101e207301c73", "")
	//if err != nil {
	//	log.Fatal("", "error", err)
	//}
	//log.Info("GetRawTransaction", "result", *hexStr)
	//result, err := cli.DecodeRawTransaction(*hexStr, false)
	//if err != nil {
	//	log.Fatal("", "error", err)
	//}
	//log.Info("DecodeRawTransaction", "result", result)
	//result, err := cli.DecodeRawTransaction("0200000001731c3007e201b1e49ac79b99994453df6df152d8da082b4025ee79ab57cc26d30000000000fdffffff02f0b9f505000000001976a9147872b1b6a3593899ce9eb25160ad3edd92d4138c88ac00a3e1110000000017a914a4c5fafd4ca4013ce0a79a6e7f3517c3e18432728700000000", false)
	//if err != nil {
	//	log.Fatal("", "error", err)
	//}
	//log.Info("DecodeRawTransaction", "result", result)
	//result, err := cli.DecodeRawTransaction("02000000000101731c3007e201b1e49ac79b99994453df6df152d8da082b4025ee79ab57cc26d30000000017160014fad971efdd5409cb3fba6d0af021e2513fa5b915fdffffff02f0b9f505000000001976a9147872b1b6a3593899ce9eb25160ad3edd92d4138c88ac00a3e1110000000017a914a4c5fafd4ca4013ce0a79a6e7f3517c3e1843272870247304402200c61419242a85f245a81f27029d524e0c9f18a5124a51bc3f411185160da9b9202201989f8ce56e1a00586e426000412f3f630f71d7344e300284545c3189bcd1bdd012102d774b2106db7564f9c84b4ea587af57e7ee82c8b4f1353489889c901409ef6de00000000", false)
	//if err != nil {
	//	log.Fatal("", "error", err) //TX decode failed
	//}
	//log.Info("DecodeRawTransaction", "result", result)
	result, err := cli.DecodeRawTransaction("0200000001731c3007e201b1e49ac79b99994453df6df152d8da082b4025ee79ab57cc26d30000000000fdffffff03605af405000000001976a9147872b1b6a3593899ce9eb25160ad3edd92d4138c88ac00a3e1110000000017a914a4c5fafd4ca4013ce0a79a6e7f3517c3e184327287256101000000000017a914f70a0b73b39e9836aef68b586fa36f3a5c58ae178700000000", false)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("DecodeRawTransaction", "result", result)
}

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

/*
Description:
A demo test of FundrawTransaction.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func FundRawTransactionTest() {
	result, err := cli.FundRawTransaction("0200000001731c3007e201b1e49ac79b99994453df6df152d8da082b4025ee79ab57cc26d30000000000fdffffff02605af405000000001976a9147872b1b6a3593899ce9eb25160ad3edd92d4138c88ac00a3e1110000000017a914a4c5fafd4ca4013ce0a79a6e7f3517c3e18432728700000000",
		nil, false)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("FundrawTransaction", "result", result)
}

/*
Description:
A demo test of GetRawTransaction.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func GetRawTransactionTest() {
	result, err := cli.GetRawTransaction("d326cc57ab79ee25402b08dad852f16ddf534499999bc79ae4b101e207301c73", "")
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetRawTransaction", "result", result)
}

/*
Description:
A demo test of GetRawTransactionVerbose.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func GetRawTransactionVerboseTest() {
	//result, err := cli.GetRawTransactionVerbose("d326cc57ab79ee25402b08dad852f16ddf534499999bc79ae4b101e207301c73", "")
	//if err != nil {
	//	log.Fatal("", "error", err)
	//}
	//log.Info("GetRawTransactionVerbose", "result", result)
	result, err := cli.GetRawTransactionVerbose("fbb4b899ae57f05154d3fdf913510ce7d91a25ca7691ddedfbbecb7bd5a1ea28", "")
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetRawTransactionVerbose", "result", result)
}

/*
Description:
A demo test of SendRawTransaction.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func SendRawTransactionTest() {
	result, err := cli.SendRawTransaction("02000000000101731c3007e201b1e49ac79b99994453df6df152d8da082b4025ee79ab57cc26d30000000017160014fad971efdd5409cb3fba6d0af021e2513fa5b915fdffffff02f0b9f505000000001976a9147872b1b6a3593899ce9eb25160ad3edd92d4138c88ac00a3e1110000000017a914a4c5fafd4ca4013ce0a79a6e7f3517c3e1843272870247304402200c61419242a85f245a81f27029d524e0c9f18a5124a51bc3f411185160da9b9202201989f8ce56e1a00586e426000412f3f630f71d7344e300284545c3189bcd1bdd012102d774b2106db7564f9c84b4ea587af57e7ee82c8b4f1353489889c901409ef6de00000000",
		false)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("SendRawTransaction", "result", result)
}

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

/*
Description:
A demo test of SignRawTransactionWithKey.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func SignRawTransactionWithKeyTest() {
	result, err := cli.SignRawTransactionWithKey("0200000001731c3007e201b1e49ac79b99994453df6df152d8da082b4025ee79ab57cc26d30000000000fdffffff02f0b9f505000000001976a9147872b1b6a3593899ce9eb25160ad3edd92d4138c88ac00a3e1110000000017a914a4c5fafd4ca4013ce0a79a6e7f3517c3e18432728700000000",
		[]string{addressFrom1PrivKey})
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("SignRawTransactionWithKey", "result", result)
	//Result: {"complete":true,"hex":"02000000000101731c3007e201b1e49ac79b99994453df6df152d8da082b4025ee79ab57cc26d30000000017160014fad971efdd5409cb3fba6d0af021e2513fa5b915fdffffff02f0b9f505000000001976a9147872b1b6a3593899ce9eb25160ad3edd92d4138c88ac00a3e1110000000017a914a4c5fafd4ca4013ce0a79a6e7f3517c3e1843272870247304402200c61419242a85f245a81f27029d524e0c9f18a5124a51bc3f411185160da9b9202201989f8ce56e1a00586e426000412f3f630f71d7344e300284545c3189bcd1bdd012102d774b2106db7564f9c84b4ea587af57e7ee82c8b4f1353489889c901409ef6de00000000"}
}

/*
Description:
A demo test of TestMempoolAccept.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func TestMempoolAcceptTest() {
	result, err := cli.TestMempoolAccept([]string{"02000000000101731c3007e201b1e49ac79b99994453df6df152d8da082b4025ee79ab57cc26d30000000017160014fad971efdd5409cb3fba6d0af021e2513fa5b915fdffffff02605af405000000001976a9147872b1b6a3593899ce9eb25160ad3edd92d4138c88ac00a3e1110000000017a914a4c5fafd4ca4013ce0a79a6e7f3517c3e1843272870247304402207fc2b478c3445a95016fdc96b38d952f7849b2d9514e7dc7042289774277cfe302204a249216f05e6e4c9a3513a639fe862ca0c08449acdeb91345771bba14149f9d012102d774b2106db7564f9c84b4ea587af57e7ee82c8b4f1353489889c901409ef6de00000000"},
		false)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("TestMempoolAccept", "result", result)
}