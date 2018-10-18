package demos

import (
	"github.com/gobasis/log"

)

const address = "mrVpnAaPap8BXLnYBc5pNfeMwwhvKcqvMq"
const privkey = "cUSG5zKm1KqRq5YzQsbe2dS86W2VRq1HwwVN5J9itRCYdaduaZED"
const addressFrom1 = "2N8GTyfYg1WKg9c2WEKK6a589gwHsKb9bzo"
const addressFrom1PrivKey = "cSZ8cpqinGkN9i5quXb8ZL4bdayn8jypkJzRruZdEYCVyVtVdvLE"
const addressTo = "mrVpnAaPap8BXLnYBc5pNfeMwwhvKcqvMq"
const address1 = "2MvobEAKN1T8VALjQZdEhr3hfDWCLxhUiLo"
const address2 = "2N7rd1LkJg5UoXDBkdfi4mUuyvTEzJtoBpj"
const address3 = "2N8p8fFpTPDp2rxJAMnYxvgv7W6Peiwk1rG"
const address1pubkey = "03b4e61ee34b37db459316b61f0e8ce540bc3e99b427c7390cb6060fd5ca4084f1"
const address1hex = "0014a906b03a31c458a92f1ce4cb512303b83483d35e"
const address2pubkey = "03c7ffdaf2ef6f6887c05b20cb541ea55b3731cbeb9b3a73e8aa601a702a0fcffa"
const address2hex = "0014c83fcb8543d609e52a04acc355da4839299d6a31"
const address3pubkey = "03b8f68c4a59fde8e2698c63e86801fa5660467860384bf71d69a1c90d8730712d"
const address3hex = "001462ae6c81f7f63b35e780bdb8f61c25b293bc5622"
//2My3izAWtnSDY4Evv7AWw9ci2fM3Tf35fC5	cNj4yz9UkRrravto6GYYgoUPdVJtpshD4BKe5nyR5fof1u5gFaca
func WalletTestInitialize() {
	addr, err := cli.GetAddressInfo(address)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetAddressInfo", "result", addr)
	if addr != nil {
		addrResult := (*addr).(map[string]interface{})
		if !addrResult["ismine"].(bool) {
			log.Info("current address state", "ismine", false)
			err := cli.ImportPrivkey(privkey)
			if err != nil {
				log.Fatal("", "error", err)
			}
			log.Info("ImportPrivkey", "result", "success")
			addr2, err2 := cli.GetAddressInfo(address)
			if err2 != nil {
				log.Fatal("", "error", err2)
			}
			log.Info("GetAddressInfo", "result", addr2)
			log.Info("address new state", "ismine", (*addr2).(map[string]interface{})["ismine"].(bool))
		} else {
			log.Info("current address state", "ismine", true)
		}
	} else {
		log.Fatal("addr is null", "error", err)
	}
}

/*
Description:
A demo test of AbandonTransaction.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func AbandonTransactionTest() {
	err := cli.AbandonTransaction("239f98f9a355fc0c3b8ac8d0556ace4c49251c5e711c22783f37e9e5f185859b")
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("AbandonTransaction", "result", "success")
}

/*
Description:
A demo test of AbortRescan.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func AbortRescanTest() {
	err := cli.AbortRescan()
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("AbortRescan", "result", "success")
}

/*
Description:
A demo test of AddMultiSigAddress.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func AddMultiSigAddressTest() {
	keys := make([]string, 0)
	keys = append(keys, address1)
	keys = append(keys, address2)
	//keys = append(keys, address3)
	result, err := cli.AddMultiSigAddress(2, keys)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("AddMultiSigAddress", "result", result) //{"address":"2N4a69o4gPC9RjQhrcKVPjj2ib3599opZSs",
	// "redeemScript":"522103b4e61ee34b37db459316b61f0e8ce540bc3e99b427c7390cb6060fd5ca4084f12103c7ffdaf2ef6f6887c05b20cb541ea55b3731cbeb9b3a73e8aa601a702a0fcffa52ae"}
}

/*
Description:
A demo test of BackupWallet.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func BackupWalletTest() {
	err := cli.BackupWallet("/root/wallet.bak")
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("BackupWallet", "result", "success")
}

/*
Description:
A demo test of BumpFee.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func BumpFeeTest() {
	option := map[string]interface{}{}
	option["replaceable"] = true
	result, err := cli.BumpFee("c55623b4b9c69595859870be911c548f1e19661ac7111c15a4a11879df1725b0", option)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("BumpFee", "result", result)
}

/*
Description:
A demo test of CreateWallet.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func CreateWalletTest() {
	result, err := cli.CreateWallet("mywallet", false)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("CreateWallet", "result", result)
}

/*
Description:
A demo test of DumpPrivkey.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func DumpPrivkeyTest() {
	//addr, err := cli.GetNewAddress()
	//if err != nil {
	//	log.Fatal("", "error", err)
	//}
	//log.Info("GetNewAddress", "address", addr)
	result, err := cli.DumpPrivkey(addressFrom1)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("DumpPrivkey", "result", result)
	//result, err := cli.DumpPrivkey("mrVpnAaPap8BXLnYBc5pNfeMwwhvKcqvMq")
	//if err != nil {
	//	log.Fatal("", "error", err)
	//}
	//log.Info("DumpPrivkey", "result", result)
}

/*
Description:
A demo test of DumpWallet.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func DumpWalletTest() {
	result, err := cli.DumpWallet("/root/wallet.dump")
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("DumpWallet", "result", result)
}

/*
Description:
A demo test of EncryptWallet.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func EncryptWalletTest() {
	err := cli.EncryptWallet("myphrase")
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("EncryptWallet", "result", "success")
}

/////*
////Description:
////Deprecated
////A demo test of GetAccount.
//// * Author: architect.bian
//// * Date: 2018/10/15 12:14
//// */
////func GetAccountTest() {
////	result, err := cli.GetAccount()
////	if err != nil {
////		log.Fatal("", "error", err)
////	}
////	log.Info("GetAccount", "result", result)
////}
//
/////*
////Description:
////Deprecated
////A demo test of GetAccountAddress.
//// * Author: architect.bian
//// * Date: 2018/10/15 12:14
//// */
////func GetAccountAddressTest() {
////	result, err := cli.GetAccountAddress()
////	if err != nil {
////		log.Fatal("", "error", err)
////	}
////	log.Info("GetAccountAddress", "result", result)
////}

/*
Description:
A demo test of GetAddressesByLabel.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func GetAddressesByLabelTest() {
	result, err := cli.GetAddressesByLabel("")
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetAddressesByLabel", "result", result)
}

/*
Description:
A demo test of GetAddressInfo.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func GetAddressInfoTest() {
	result, err := cli.GetAddressInfo("2N8p8fFpTPDp2rxJAMnYxvgv7W6Peiwk1rG")
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetAddressInfo", "result", result)
}

/*
Description:
A demo test of GetBalance.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func GetBalanceTest() {
	result, err := cli.GetBalance()
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetBalance", "result", result)
}

/*
Description:
A demo test of GetBalanceEntire.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func GetBalanceEntireTest() {
	result, err := cli.GetBalanceEntire("*", 1, true)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetBalance", "result", result)
}

/*
Description:
A demo test of GetNewAddress.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func GetNewAddressTest() {
	result, err := cli.GetNewAddress()
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetNewAddress", "result", result)
}

/*
Description: 
A demo test of GetNewAddressEntireTest.
 * Author: architect.bian
 * Date: 2018/10/15 23:44
 */
func GetNewAddressEntireTest() {
	//result, err := cli.GetNewAddressEntire("", "legacy")
	//if err != nil {
	//	log.Fatal("", "error", err)
	//}
	//log.Info("GetNewAddressEntire", "result-legacy", result) //muzMwJUEkoNoepi6Krx4m4hY56YvgnE2Mf
	result, err := cli.GetNewAddressEntire("addr2018", "legacy")
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetNewAddressEntire", "result-legacy", result) //muzMwJUEkoNoepi6Krx4m4hY56YvgnE2Mf
	result2, err2 := cli.GetNewAddressEntire("addr2018", "p2sh-segwit")
	if err2 != nil {
		log.Fatal("", "error", err2)
	}
	log.Info("GetNewAddressEntire", "result-p2sh-segwit", result2) //2Mx18zD3m3MSG1wkex416pegoKKaDZJmogx
	result3, err3 := cli.GetNewAddressEntire("addr2018", "bech32")
	if err3 != nil {
		log.Fatal("", "error", err3)
	}
	log.Info("GetNewAddressEntire", "result-bech32", result3) //bcrt1q9grsd8w802zjnmj935r4q074rdhq786danjvg0
}

/*
Description:
A demo test of GetRawChangeAddress.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func GetRawChangeAddressTest() {
	result, err := cli.GetRawChangeAddressEntire("legacy")
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetRawChangeAddress", "result", result)
}

/////*
////Description:
////Deprecated
////A demo test of GetReceivedByAccount.
//// * Author: architect.bian
//// * Date: 2018/10/15 12:14
//// */
////func GetReceivedByAccountTest() {
////	result, err := cli.GetReceivedByAccount()
////	if err != nil {
////		log.Fatal("", "error", err)
////	}
////	log.Info("GetReceivedByAccount", "result", result)
////}

/*
Description:
A demo test of GetReceivedByAddress.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func GetReceivedByAddressTest() {
	result, err := cli.GetReceivedByAddress(addressTo, 0)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetReceivedByAddress", "result", result)
}

/*
Description:
A demo test of GetTransaction.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func GetTransactionTest() {
	result, err := cli.GetTransaction("fbb4b899ae57f05154d3fdf913510ce7d91a25ca7691ddedfbbecb7bd5a1ea28", true)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetTransaction", "result", result)
	//result, err := cli.GetTransaction("239f98f9a355fc0c3b8ac8d0556ace4c49251c5e711c22783f37e9e5f185859b", true)
	//if err != nil {
	//	log.Fatal("", "error", err)
	//}
	//log.Info("GetTransaction", "result", result)
	//result2, err2 := cli.GetTransaction("65bcaab247a598ae175eb5edf34531616ec7d8c4147d8ca6ad2b9b3c69543dc1", true)
	//if err2 != nil {
	//	log.Fatal("", "error", err2)
	//}
	//log.Info("GetTransaction", "result", result2)
}

/*
Description:
A demo test of GetUnconfirmedBalance.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func GetUnconfirmedBalanceTest() {
	result, err := cli.GetUnconfirmedBalance()
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetUnconfirmedBalance", "result", result)
}

/*
Description:
A demo test of getwalletinfo.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func GetWalletInfoTest() {
	result, err := cli.GetWalletInfo()
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("getwalletinfo", "result", result)
}

/*
Description:
A demo test of Importaddress.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func ImportaddressTest() {
	result, err := cli.Importaddress("76a9147872b1b6a3593899ce9eb25160ad3edd92d4138c88ac")
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("Importaddress", "result", result)
}

/*
Description:
A demo test of ImportMulti.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func ImportMultiTest() {
	requests := make([]interface{}, 0)
	scriptPubKey := make(map[string]interface{})
	addr := make(map[string]interface{})
	addr["address"] = "mrVpnAaPap8BXLnYBc5pNfeMwwhvKcqvMq"
	scriptPubKey["scriptPubKey"] = addr
	scriptPubKey["timestamp"] = 1455191480
	requests = append(requests[:], scriptPubKey)
	option := make(map[string]interface{})
	option["rescan"] = false
	result, err := cli.ImportMulti(requests, option)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("ImportMulti", "result", result)
}

/*
Description:
A demo test of ImportPrivkey.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func ImportPrivkeyTest() {
	err := cli.ImportPrivkey(privkey)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("ImportPrivkey", "result", "success")
}

///*
//Description:
//A demo test of ImportPrunedFunds.
// * Author: architect.bian
// * Date: 2018/10/15 12:14
// */
//func ImportPrunedFundsTest() {
//	result, err := cli.ImportPrunedFunds()
//	if err != nil {
//		log.Fatal("", "error", err)
//	}
//	log.Info("ImportPrunedFunds", "result", result)
//}

/*
Description:
A demo test of ImportPubkey.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func ImportPubkeyTest() {
	err := cli.ImportPubkey("76a9147872b1b6a3593899ce9eb25160ad3edd92d4138c88ac")
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("ImportPubkey", "result", "success")
}

/*
Description:
A demo test of ImportWallet.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func ImportWalletTest() {
	err := cli.ImportWallet("/root/wallet.dump")
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("ImportWallet", "result", "success")
}

/*
Description:
A demo test of KeypoolRefill.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func KeypoolRefillTest() {
	err := cli.KeypoolRefill(100)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("KeypoolRefill", "result", "success")
}

/////*
////Description:
////Deprecated
////A demo test of ListAccounts.
//// * Author: architect.bian
//// * Date: 2018/10/15 12:14
//// */
////func ListAccountsTest() {
////	result, err := cli.ListAccounts()
////	if err != nil {
////		log.Fatal("", "error", err)
////	}
////	log.Info("ListAccounts", "result", result)
////}
//
/*
Description:
A demo test of ListAddressGroupings.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func ListAddressGroupingsTest() {
	result, err := cli.ListAddressGroupings()
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("ListAddressGroupings", "result", result)
}

/*
Description:
A demo test of ListLabels.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func ListLabelsTest() {
	result, err := cli.ListLabels("")
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("ListLabels", "result", result)
	result2, err2 := cli.ListLabels("receive")
	if err2 != nil {
		log.Fatal("", "error", err2)
	}
	log.Info("ListLabels", "result", result2)
	result3, err3 := cli.ListLabels("send")
	if err3 != nil {
		log.Fatal("", "error", err3)
	}
	log.Info("ListLabels", "result", result3) //empty?
}

/*
Description:
A demo test of ListLockUnspent.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func ListLockUnspentTest() {
	result, err := cli.ListLockUnspent()
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("ListLockUnspent", "result", result)
}

/////*
////Description:
////Deprecated
////A demo test of ListReceivedByAccount.
//// * Author: architect.bian
//// * Date: 2018/10/15 12:14
//// */
////func ListReceivedByAccountTest() {
////	result, err := cli.ListReceivedByAccount()
////	if err != nil {
////		log.Fatal("", "error", err)
////	}
////	log.Info("ListReceivedByAccount", "result", result)
////}

/*
Description:
A demo test of ListReceivedByAddress.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func ListReceivedByAddressTest() {
	result, err := cli.ListReceivedByAddress(0, true, true, "")
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("ListReceivedByAddress", "result", result)
}

/*
Description:
A demo test of ListsInceBlock.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func ListSinceBlockTest() {
	result, err := cli.ListSinceBlock("69bf27f38a4d66f1abd4a83b69a7c2345971b8efbf121ed740dc7aaf5c4ce868", 1, false, true)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("ListsInceBlock", "result", result)
}

/*
Description:
A demo test of ListTransactions.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func ListTransactionsTest() {
	result, err := cli.ListTransactions()
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("ListTransactions", "result", result)
}

/*
Description:
A demo test of ListUnspent.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func ListUnspentTest() {
	result, err := cli.ListUnspent()
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("ListUnspent", "result", result)
}

/*
Description:
A demo test of ListUnspentEntire.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func ListUnspentEntireTest() {
	result, err := cli.ListUnspentEntire(0, 999999, nil, true, nil)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("ListUnspentEntire", "result", result)
}

/*
Description:
A demo test of ListWallets.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func ListWalletsTest() {
	result, err := cli.ListWallets()
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("ListWallets", "result", result)
}

/*
Description:
A demo test of LoadWallet.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func LoadWalletTest() {
	//result, err := cli.LoadWallet("/root/.bitcoin/regtest/wallets")
	result, err := cli.LoadWallet("/root/.bitcoin/regtest/wallets/mywallet")
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("LoadWallet", "result", result)
}

/*
Description:
A demo test of LockUnspent.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func LockUnspentTest() {
	result, err := cli.ListUnspent()
	if err != nil {
		log.Fatal("", "error", err)
	}
	utxos := (*result).([]interface{})
	log.Info("ListUnspent", "result", utxos)
	if len(utxos) > 0 {
		utxo := utxos[0].(map[string]interface{})
		txs := make([]map[string]interface{}, 1)
		tx := make(map[string]interface{})
		tx["txid"] = utxo["txid"]
		tx["vout"] = utxo["vout"]
		txs[0] = tx
		result1, err1 := cli.LockUnspent(false, txs)
		if err1 != nil {
			log.Fatal("", "error", err1)
		}
		log.Info("LockUnspent", "result", result1)
		ListLockUnspentTest()
		result2, err2 := cli.LockUnspent(true, txs)
		if err2 != nil {
			log.Fatal("", "error", err2)
		}
		log.Info("LockUnspent", "result", result2)
		ListLockUnspentTest()
	}

}

/////*
////Description:
////Deprecated
////A demo test of Move.
//// * Author: architect.bian
//// * Date: 2018/10/15 12:14
//// */
////func MoveTest() {
////	result, err := cli.Move()
////	if err != nil {
////		log.Fatal("", "error", err)
////	}
////	log.Info("Move", "result", result)
////}

/*
Description:
A demo test of RemovePrunedFunds.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func RemovePrunedFundsTest() {
	err := cli.RemovePrunedFunds("76a261989d3ff432f0f1846f23c6653deec70a548be76c08560bca3b1e888e77")
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("RemovePrunedFunds", "result", "success")
}

/*
Description:
A demo test of RescanBlockChain.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func RescanBlockChainTest() {
	result, err := cli.RescanBlockChain()
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("RescanBlockChain", "result", result)
}

/////*
////Description:
////Deprecated
////A demo test of SendFrom.
//// * Author: architect.bian
//// * Date: 2018/10/15 12:14
//// */
////func SendFromTest() {
////	result, err := cli.SendFrom()
////	if err != nil {
////		log.Fatal("", "error", err)
////	}
////	log.Info("SendFrom", "result", result)
////}

/*
Description:
A demo test of SendMany.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func SendManyTest() {
	to := make(map[string]interface{})
	to[addressTo] = 10
	result, err := cli.SendMany(to)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("SendMany", "result", result)
}

/*
Description:
A demo test of SendToAddress.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func SendToAddressTest() {
	result, err := cli.SendToAddress(addressTo, 150)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("SendToAddress", "result", result)
}

/*
Description:
A demo test of SendToAddressEntire.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func SendToAddressEntireTest() {
	result, err := cli.SendToAddressEntire(addressTo, 6, "", "",
		false, true, 1, "UNSET")
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("SendToAddress", "result", result)
}

/////*
////Description:
////Deprecated
////A demo test of SetAccount.
//// * Author: architect.bian
//// * Date: 2018/10/15 12:14
//// */
////func SetAccountTest() {
////	result, err := cli.SetAccount()
////	if err != nil {
////		log.Fatal("", "error", err)
////	}
////	log.Info("SetAccount", "result", result)
////}

/*
Description:
A demo test of SetHDSeed.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func SetHDSeedTest() {
	err := cli.SetHDSeed()
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("SetHDSeed", "result", "success")
}

/*
Description:
A demo test of SetTXFee.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func SetTXFeeTest() {
	result, err := cli.SetTXFee(0.00001)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("SetTXFee", "result", result)
}

/*
Description:
A demo test of SignMessage.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func SignMessageTest() {
	result, err := cli.SignMessage(address, "hi world!")
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("SignMessage", "result", result)
}

///*
//Description:
//Deprecated
//A demo test of SignRawtransactionWithWallet.
//* Author: architect.bian
//* Date: 2018/10/15 12:14
//*/
//func SignRawtransactionWithWalletTest() {
//	result, err := cli.SignRawtransactionWithWallet()
//	if err != nil {
//		log.Fatal("", "error", err)
//	}
//	log.Info("SignRawtransactionWithWallet", "result", result)
//}

/*
Description:
A demo test of UnloadWallet.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func UnloadWalletTest() {
	err := cli.UnloadWallet("/root/.bitcoin/regtest/wallets/mywallet")
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("UnloadWallet", "result", "success")
}

///*
//Description:
//A demo test of WalletCreateFundedPSBT.
// * Author: architect.bian
// * Date: 2018/10/15 12:14
// */
//func WalletCreateFundedPSBTTest() {
//	result, err := cli.WalletCreateFundedPSBT()
//	if err != nil {
//		log.Fatal("", "error", err)
//	}
//	log.Info("WalletCreateFundedPSBT", "result", result)
//}

/*
Description:
A demo test of WalletLock.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func WalletLockTest() {
	err := cli.WalletLock()
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("WalletLock", "result", "success")
}

/*
Description:
A demo test of WalletPassphrase.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func WalletPassphraseTest() {
	err := cli.WalletPassphrase("myphrase", 60000)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("WalletPassphrase", "result", "success")
	SignMessageTest()
}

/*
Description:
A demo test of WalletPassphraseChange.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func WalletPassphraseChangeTest() {
	err := cli.WalletPassphraseChange("myphrase", "rpcpwd")
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("WalletPassphraseChange", "result", "success")
	err2 := cli.WalletPassphraseChange("rpcpwd", "myphrase")
	if err2 != nil {
		log.Fatal("", "error", err2)
	}
	log.Info("WalletPassphraseChange", "result", "success")
}

///*
//Description:
//A demo test of WalletProcessPSBT.
// * Author: architect.bian
// * Date: 2018/10/15 12:14
// */
//func WalletProcessPSBTTest() {
//	result, err := cli.WalletProcessPSBT()
//	if err != nil {
//		log.Fatal("", "error", err)
//	}
//	log.Info("WalletProcessPSBT", "result", result)
//}