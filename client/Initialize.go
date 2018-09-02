package client

/*
Description:
register all the rpc command. Reference: https://bitcoin.org/en/developer-reference
 * Author: architect.bian
 * Date: 2018/08/27 19:24
 */
func init() {
	// No special flags for commands in this file.
	flags := UsageFlag(0)

	//MustRegisterCmd("addnode", (*AddNodeCmd)(nil), flags)
	//MustRegisterCmd("createrawtransaction", (*CreateRawTransactionCmd)(nil), flags)
	//MustRegisterCmd("decoderawtransaction", (*DecodeRawTransactionCmd)(nil), flags)
	//MustRegisterCmd("decodescript", (*DecodeScriptCmd)(nil), flags)
	//MustRegisterCmd("getaddednodeinfo", (*GetAddedNodeInfoCmd)(nil), flags)
	MustRegisterCmd("getbestblockhash", (*GetBestBlockHashCmd)(nil), flags)
	MustRegisterCmd("getblock", (*GetBlockCmd)(nil), flags)
	MustRegisterCmd("getblockchaininfo", (*GetBlockChainInfoCmd)(nil), flags)
	MustRegisterCmd("getblockcount", (*GetBlockCountCmd)(nil), flags)
	MustRegisterCmd("getblockhash", (*GetBlockHashCmd)(nil), flags)
	MustRegisterCmd("getblockheader", (*GetBlockHeaderCmd)(nil), flags)
	//MustRegisterCmd("getblocktemplate", (*GetBlockTemplateCmd)(nil), flags)
	MustRegisterCmd("getcfilter", (*GetCFilterCmd)(nil), flags)
	MustRegisterCmd("getcfilterheader", (*GetCFilterHeaderCmd)(nil), flags)
	//MustRegisterCmd("getchaintips", (*GetChainTipsCmd)(nil), flags)
	//MustRegisterCmd("getconnectioncount", (*GetConnectionCountCmd)(nil), flags)
	MustRegisterCmd("getdifficulty", (*GetDifficultyCmd)(nil), flags)
	//MustRegisterCmd("getgenerate", (*GetGenerateCmd)(nil), flags)
	//MustRegisterCmd("gethashespersec", (*GetHashesPerSecCmd)(nil), flags)
	//MustRegisterCmd("getinfo", (*GetInfoCmd)(nil), flags)
	MustRegisterCmd("getmempoolentry", (*GetMempoolEntryCmd)(nil), flags)
	//MustRegisterCmd("getmempoolinfo", (*GetMempoolInfoCmd)(nil), flags)
	//MustRegisterCmd("getmininginfo", (*GetMiningInfoCmd)(nil), flags)
	//MustRegisterCmd("getnetworkinfo", (*GetNetworkInfoCmd)(nil), flags)
	//MustRegisterCmd("getnettotals", (*GetNetTotalsCmd)(nil), flags)
	//MustRegisterCmd("getnetworkhashps", (*GetNetworkHashPSCmd)(nil), flags)
	//MustRegisterCmd("getpeerinfo", (*GetPeerInfoCmd)(nil), flags)
	MustRegisterCmd("getrawmempool", (*GetRawMempoolCmd)(nil), flags)
	//MustRegisterCmd("getrawtransaction", (*GetRawTransactionCmd)(nil), flags)
	MustRegisterCmd("gettxout", (*GetTxOutCmd)(nil), flags)
	//MustRegisterCmd("gettxoutproof", (*GetTxOutProofCmd)(nil), flags)
	//MustRegisterCmd("gettxoutsetinfo", (*GetTxOutSetInfoCmd)(nil), flags)
	//MustRegisterCmd("getwork", (*GetWorkCmd)(nil), flags)
	//MustRegisterCmd("help", (*HelpCmd)(nil), flags)
	MustRegisterCmd("invalidateblock", (*InvalidateBlockCmd)(nil), flags)
	//MustRegisterCmd("ping", (*PingCmd)(nil), flags)
	//MustRegisterCmd("preciousblock", (*PreciousBlockCmd)(nil), flags)
	//MustRegisterCmd("reconsiderblock", (*ReconsiderBlockCmd)(nil), flags)
	//MustRegisterCmd("searchrawtransactions", (*SearchRawTransactionsCmd)(nil), flags)
	//MustRegisterCmd("sendrawtransaction", (*SendRawTransactionCmd)(nil), flags)
	//MustRegisterCmd("setgenerate", (*SetGenerateCmd)(nil), flags)
	//MustRegisterCmd("stop", (*StopCmd)(nil), flags)
	//MustRegisterCmd("submitblock", (*SubmitBlockCmd)(nil), flags)
	//MustRegisterCmd("uptime", (*UptimeCmd)(nil), flags)
	//MustRegisterCmd("validateaddress", (*ValidateAddressCmd)(nil), flags)
	MustRegisterCmd("verifychain", (*VerifyChainCmd)(nil), flags)
	//MustRegisterCmd("verifymessage", (*VerifyMessageCmd)(nil), flags)
	//MustRegisterCmd("verifytxoutproof", (*VerifyTxOutProofCmd)(nil), flags)
}
