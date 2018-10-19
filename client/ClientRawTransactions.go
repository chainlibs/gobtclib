package client

import (
	"github.com/chainlibs/gobtclib/futures"
)

/*
Description:
CombinePSBT returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See CombinePSBT for more details.
 * Author: architect.bian
 * Date: 2018/10/15 20:09
 */
func (c *Client) CombinePSBTAsync(txHexs []string) futures.FutureResult {
	cmd := NewCommand("combinepsbt", txHexs)
	return c.sendCmd(cmd)
}

/*
Description: 
Combine multiple partially signed Bitcoin transactions into one transaction.
Implements the Combiner role.
 * Author: architect.bian
 * Date: 2018/10/15 20:18
 */
func (c *Client) CombinePSBT(txHexs []string) (*interface{}, error) {
	return c.CombinePSBTAsync(txHexs).Receive()
}

/*
Description:
CombineRawTransactionAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See CombineRawTransaction for more details.
 * Author: architect.bian
 * Date: 2018/10/15 20:09
 */
func (c *Client) CombineRawTransactionAsync(txs []string) futures.FutureResult {
	cmd := NewCommand("combinerawtransaction", txs)
	return c.sendCmd(cmd)
}

/*
Description:
Combine multiple partially signed transactions into one transaction.
The combined transaction may be another partially signed transaction or a
fully signed transaction.
 * Author: architect.bian
 * Date: 2018/10/15 20:18
 */
func (c *Client) CombineRawTransaction(txs []string) (*interface{}, error) {
	return c.CombineRawTransactionAsync(txs).Receive()
}

/*
Description:
ConvertToPSBTAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See ConvertToPSBT for more details.
 * Author: architect.bian
 * Date: 2018/10/15 20:09
 */
func (c *Client) ConvertToPSBTAsync(txHex string, permitsigdata bool, iswitness bool) futures.FutureString {
	cmd := NewCommand("converttopsbt", txHex, permitsigdata, iswitness)
	return c.sendCmd(cmd)
}

/*
Description:
Converts a network serialized transaction to a PSBT. This should be used only with createrawtransaction and fundrawtransaction
createpsbt and walletcreatefundedpsbt should be used for new applications.
 * Author: architect.bian
 * Date: 2018/10/15 20:22
 */
func (c *Client) ConvertToPSBT(txHex string) (*string, error) {
	cmd := NewCommand("converttopsbt", txHex)
	return futures.FutureString(c.sendCmd(cmd)).Receive()
}

/*
Description: 
Converts a network serialized transaction to a PSBT. This should be used only with createrawtransaction and fundrawtransaction
createpsbt and walletcreatefundedpsbt should be used for new applications.
 * Author: architect.bian
 * Date: 2018/10/15 20:23
 */
func (c *Client) ConvertToPSBTEntire(txHex string, permitsigdata bool, iswitness bool) (*string, error) {
	return c.ConvertToPSBTAsync(txHex, permitsigdata, iswitness).Receive()
}

/*
Description:
CreatepSBTAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See CreatepSBT for more details.
 * Author: architect.bian
 * Date: 2018/10/15 20:09
 */
func (c *Client) CreatepSBTAsync(inputs []map[string]interface{}, outputs []map[string]interface{}, locktime int32, replaceable bool) futures.FutureString {
	cmd := NewCommand("createpsbt", inputs, outputs, locktime, replaceable)
	return c.sendCmd(cmd)
}

/*
Description:
Creates a transaction in the Partially Signed Transaction format.
Implements the Creator role.
 * Author: architect.bian
 * Date: 2018/10/15 20:27
 */
func (c *Client) CreatepSBT(inputs []map[string]interface{}, outputs []map[string]interface{}, locktime int32, replaceable bool) (*string, error) {
	return c.CreatepSBTAsync(inputs, outputs, locktime, replaceable).Receive()
}

/*
Description:
CreateRawTransactionAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See CreateRawTransaction for more details.
 * Author: architect.bian
 * Date: 2018/10/15 20:09
 */
func (c *Client) CreateRawTransactionAsync(inputs []map[string]interface{}, outputs []map[string]interface{}, locktime int32, replaceable bool) futures.FutureString {
	cmd := NewCommand("createrawtransaction", inputs, outputs, locktime, replaceable)
	return c.sendCmd(cmd)
}

/*
Description:
Create a transaction spending the given inputs and creating new outputs.
Outputs can be addresses or data.
Returns hex-encoded raw transaction.
Note that the transaction's inputs are not signed, and
it is not stored in the wallet or transmitted to the network.
 * Author: architect.bian
 * Date: 2018/10/15 20:31
 */
func (c *Client) CreateRawTransaction(inputs []map[string]interface{}, outputs []map[string]interface{},
	locktime int32, replaceable bool) (*string, error) {
		return c.CreateRawTransactionAsync(inputs, outputs, locktime, replaceable).Receive()
}

/*
Description:
DecodePSBTAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See DecodePSBT for more details.
 * Author: architect.bian
 * Date: 2018/10/15 20:09
 */
func (c *Client) DecodePSBTAsync(psbtBase64 string) futures.FutureResult {
	cmd := NewCommand("decodepsbt", psbtBase64)
	return c.sendCmd(cmd)
}

/*
Description: 
Return a JSON object representing the serialized, base64-encoded partially signed Bitcoin transaction.
 * Author: architect.bian
 * Date: 2018/10/15 20:32
 */
func (c *Client) DecodePSBT(psbtBase64 string) (*interface{}, error) {
	return c.DecodePSBTAsync(psbtBase64).Receive()
}

/*
Description:
DecodeRawTransactionAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See DecodeRawTransaction for more details.
 * Author: architect.bian
 * Date: 2018/10/15 20:09
 */
func (c *Client) DecodeRawTransactionAsync(txHex string, isWitness bool) futures.FutureResult {
	cmd := NewCommand("decoderawtransaction", txHex, isWitness)
	return c.sendCmd(cmd)
}

/*
Description: 
Return a JSON object representing the serialized, hex-encoded transaction.
 * Author: architect.bian
 * Date: 2018/10/15 20:34
 */
func (c *Client) DecodeRawTransaction(txHex string, isWitness bool) (*interface{}, error) {
	return c.DecodeRawTransactionAsync(txHex, isWitness).Receive()
}

/*
Description:
DecodeScriptAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See DecodeScript for more details.
 * Author: architect.bian
 * Date: 2018/10/15 20:09
 */
func (c *Client) DecodeScriptAsync(scriptHex string) futures.FutureResult {
	cmd := NewCommand("decodescript", scriptHex)
	return c.sendCmd(cmd)
}

/*
Description: 
Decode a hex-encoded script.
 * Author: architect.bian
 * Date: 2018/10/15 20:36
 */
func (c *Client) DecodeScript(scriptHex string) (*interface{}, error) {
	return c.DecodeScriptAsync(scriptHex).Receive()
}

/*
Description:
FinalizePSBTAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See FinalizePSBT for more details.
 * Author: architect.bian
 * Date: 2018/10/15 20:09
 */
func (c *Client) FinalizePSBTAsync(psbt string, extract bool) futures.FutureResult {
	cmd := NewCommand("finalizepsbt", psbt, extract)
	return c.sendCmd(cmd)
}

/*
Description:
Finalize the inputs of a PSBT. If the transaction is fully signed, it will produce a
network serialized transaction which can be broadcast with sendrawtransaction. Otherwise a PSBT will be
created which has the final_scriptSig and final_scriptWitness fields filled for inputs that are complete.
Implements the Finalizer and Extractor roles.
 * Author: architect.bian
 * Date: 2018/10/15 20:38
 */
func (c *Client) FinalizePSBT(psbt string, extract bool) (*interface{}, error) {
	return c.FinalizePSBTAsync(psbt, extract).Receive()
}

/*
Description:
FundRawTransactionAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See FundRawTransaction for more details.
 * Author: architect.bian
 * Date: 2018/10/15 20:09
 */
func (c *Client) FundRawTransactionAsync(txHex string, options map[string]interface{}, iswitness bool) futures.FutureResult {
	cmd := NewCommand("fundrawtransaction", txHex, options, iswitness)
	return c.sendCmd(cmd)
}

/*
Description:
Add inputs to a transaction until it has enough in value to meet its out value.
This will not modify existing inputs, and will add at most one change output to the outputs.
No existing outputs will be modified unless "subtractFeeFromOutputs" is specified.
Note that inputs which were signed may need to be resigned after completion since in/outputs have been added.
The inputs added will not be signed, use signrawtransaction for that.
Note that all existing inputs must have their previous output transaction be in the wallet.
Note that all inputs selected must be of standard form and P2SH scripts must be
in the wallet using importaddress or addmultisigaddress (to calculate fees).
You can see whether this is the case by checking the "solvable" field in the listunspent output.
Only pay-to-pubkey, multisig, and P2SH versions thereof are currently supported for watch-only
 * Author: architect.bian
 * Date: 2018/10/15 20:36
 */
func (c *Client) FundRawTransaction(txHex string, options map[string]interface{}, iswitness bool) (*interface{}, error) {
	return c.FundRawTransactionAsync(txHex, options, iswitness).Receive()
}

/*
Description:
GetRawTransactionAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See GetRawTransaction for more details.
 * Author: architect.bian
 * Date: 2018/10/15 20:09
 */
func (c *Client) GetRawTransactionAsync(txid string, verbose bool, blockhash string) futures.FutureResult {
	cmd := NewCommand("getrawtransaction", txid, verbose)
	if blockhash != "" {
		cmd.AddArgs(blockhash)
	}
	return c.sendCmd(cmd)
}

/*
Description:
NOTE: By default this function only works for mempool transactions. If the -txindex option is
enabled, it also works for blockchain transactions. If the block which contains the transaction
is known, its hash can be provided even for nodes without -txindex. Note that if a blockhash is
provided, only that block will be searched and if the transaction is in the mempool or other
blocks, or if this node does not have the given block available, the transaction will not be found.
DEPRECATED: for now, it also works for transactions with unspent outputs.

Return the raw transaction data.
 * Author: architect.bian
 * Date: 2018/10/15 20:41
 */
func (c *Client) GetRawTransaction(txid string, blockhash string) (*string, error) {
	result, err := futures.FutureString(c.GetRawTransactionAsync(txid, false, blockhash)).Receive()
	if err != nil {
		return nil, err
	}
	return result, nil
}

/*
Description:
Return the verbose raw transaction data.
 * Author: architect.bian
 * Date: 2018/10/15 20:42
 */
func (c *Client) GetRawTransactionVerbose(txid string, blockhash string) (*interface{}, error) {
	return c.GetRawTransactionAsync(txid, true, blockhash).Receive()
}

/*
Description:
SendRawTransactionAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See SendRawTransaction for more details.
 * Author: architect.bian
 * Date: 2018/10/15 20:09
 */
func (c *Client) SendRawTransactionAsync(txHex string, allowhighfees bool) futures.FutureResult {
	cmd := NewCommand("sendrawtransaction", txHex, allowhighfees)
	return c.sendCmd(cmd)
}

/*
Description: 
Submits raw transaction (serialized, hex-encoded) to local node and network.

Also see createrawtransaction and signrawtransaction calls.
 * Author: architect.bian
 * Date: 2018/10/15 20:54
 */
func (c *Client) SendRawTransaction(txHex string, allowhighfees bool) (*interface{}, error) {
	return c.SendRawTransactionAsync(txHex, allowhighfees).Receive()
}

/*
Description:
SignRawTransactionWithKeyAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See SignRawTransactionWithKey for more details.
 * Author: architect.bian
 * Date: 2018/10/15 20:09
 */
func (c *Client) SignRawTransactionWithKeyAsync(txHex string, privkeys []string, prevtxs interface{}, sighashtype string) futures.FutureResult {
	cmd := NewCommand("signrawtransactionwithkey", txHex, privkeys, prevtxs, sighashtype)
	return c.sendCmd(cmd)
}

/*
Description:
Sign inputs for raw transaction (serialized, hex-encoded).
The second argument is an array of base58-encoded private
keys that will be the only keys used to sign the transaction.
The third optional argument (may be null) is an array of previous transaction outputs that
this transaction depends on but may not yet be in the block chain.
 * Author: architect.bian
 * Date: 2018/10/15 21:11
 */
func (c *Client) SignRawTransactionWithKey(txHex string, privkeys []string) (*interface{}, error) {
	return c.SignRawTransactionWithKeyEntire(txHex, privkeys, nil, "ALL")
}

/*
Description:
Sign inputs for raw transaction (serialized, hex-encoded).
The second argument is an array of base58-encoded private
keys that will be the only keys used to sign the transaction.
The third optional argument (may be null) is an array of previous transaction outputs that
this transaction depends on but may not yet be in the block chain.
 * Author: architect.bian
 * Date: 2018/10/15 21:11
 */
func (c *Client) SignRawTransactionWithKeyEntire(txHex string, privkeys []string, prevtxs interface{}, sighashtype string) (*interface{}, error) {
	return c.SignRawTransactionWithKeyAsync(txHex, privkeys, prevtxs, sighashtype).Receive()
}

/*
Description:
TestMempoolAcceptAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See TestMempoolAccept for more details.
 * Author: architect.bian
 * Date: 2018/10/15 20:09
 */
func (c *Client) TestMempoolAcceptAsync(rawTXs []string, allowHighFees bool) futures.FutureResult {
	cmd := NewCommand("testmempoolaccept", rawTXs, allowHighFees)
	return c.sendCmd(cmd)
}

/*
Description:
Returns if raw transaction (serialized, hex-encoded) would be accepted by mempool.
 * Author: architect.bian
 * Date: 2018/10/15 23:13
 */
func (c *Client) TestMempoolAccept(rawTXs []string, allowHighFees bool) (*interface{}, error) {
	return c.TestMempoolAcceptAsync(rawTXs, allowHighFees).Receive()
}
