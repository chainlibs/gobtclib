package client

//// FutureRescanBlocksResult is a future promise to deliver the result of a
//// RescanBlocksAsync RPC invocation (or an applicable error).
////
//// NOTE: This is a btcsuite extension ported from
//// github.com/decred/dcrrpcclient.
//type FutureRescanBlocksResult chan *response
//
//// Receive waits for the response promised by the future and returns the
//// discovered rescanblocks data.
////
//// NOTE: This is a btcsuite extension ported from
//// github.com/decred/dcrrpcclient.
//func (r FutureRescanBlocksResult) Receive() ([]RescannedBlock, error) {
//	res, err := receiveFuture(r)
//	if err != nil {
//		return nil, err
//	}
//
//	var rescanBlocksResult []RescannedBlock
//	err = json.Unmarshal(res, &rescanBlocksResult)
//	if err != nil {
//		return nil, err
//	}
//
//	return rescanBlocksResult, nil
//}
//
//// RescanBlocksCmd defines the rescan JSON-RPC command.
////
//// NOTE: This is a btcd extension ported from github.com/decred/dcrd/dcrjson
//// and requires a websocket connection.
//type RescanBlocksCmd struct {
//	// Block hashes as a string array.
//	BlockHashes []string
//}
//
//// RescanBlocksAsync returns an instance of a type that can be used to get the
//// result of the RPC at some future time by invoking the Receive function on the
//// returned instance.
////
//// See RescanBlocks for the blocking version and more details.
////
//// NOTE: This is a btcsuite extension ported from
//// github.com/decred/dcrrpcclient.
//func (c *Client) RescanBlocksAsync(blockHashes []Hash) futures.FutureRescanBlocksResult {
//	strBlockHashes := make([]string, len(blockHashes))
//	for i := range blockHashes {
//		strBlockHashes[i] = blockHashes[i].String()
//	}
//
//	cmd := &RescanBlocksCmd{BlockHashes: strBlockHashes}
//	return c.sendCmd(cmd)
//}
//
//// RescanBlocks rescans the blocks identified by blockHashes, in order, using
//// the client's loaded transaction filter.  The blocks do not need to be on the
//// main chain, but they do need to be adjacent to each other.
////
//// NOTE: This is a btcsuite extension ported from
//// github.com/decred/dcrrpcclient.
//func (c *Client) RescanBlocks(blockHashes []Hash) ([]RescannedBlock, error) {
//	return c.RescanBlocksAsync(blockHashes).Receive()
//}

//// FutureGetCFilterResult is a future promise to deliver the result of a
//// GetCFilterAsync RPC invocation (or an applicable error).
//type FutureGetCFilterResult chan *response
//
//// Receive waits for the response promised by the future and returns the raw
//// filter requested from the server given its block hash.
//func (r FutureGetCFilterResult) Receive() (*results.MsgCFilter, error) {
//	res, err := receiveFuture(r)
//	if err != nil {
//		return nil, err
//	}
//
//	// Unmarshal result as a string.
//	var filterHex string
//	err = json.Unmarshal(res, &filterHex)
//	if err != nil {
//		return nil, err
//	}
//
//	// Decode the serialized cf hex to raw bytes.
//	serializedFilter, err := hex.DecodeString(filterHex)
//	if err != nil {
//		return nil, err
//	}
//
//	// Assign the filter bytes to the correct field of the wire message.
//	// We aren't going to set the block hash or extended flag, since we
//	// don't actually get that back in the RPC response.
//	var msgCFilter MsgCFilter
//	msgCFilter.Data = serializedFilter
//	return &msgCFilter, nil
//}
//
//// FilterType is used to represent a filter type.
//type FilterType uint8
//
//// GetCFilterCmd defines the getcfilter JSON-RPC command.
//type GetCFilterCmd struct {
//	Hash       string
//	FilterType FilterType
//}
//
//// GetCFilterAsync returns an instance of a type that can be used to get the
//// result of the RPC at some future time by invoking the Receive function on the
//// returned instance.
////
//// See GetCFilter for the blocking version and more details.
//func (c *Client) GetCFilterAsync(blockHash *results.Hash, filterType FilterType) futures.FutureGetCFilterResult {
//	hash := ""
//	if blockHash != nil {
//		hash = blockHash.String()
//	}
//
//	cmd := &GetCFilterCmd{
//		Hash:       hash,
//		FilterType: filterType,
//	}
//	return c.sendCmd(cmd)
//}
//
////GetCFilter returns a raw filter from the server given its block hash.
//func (c *Client) GetCFilter(blockHash *results.Hash, filterType FilterType) (*results.MsgCFilter, error) {
//	return c.GetCFilterAsync(blockHash, filterType).Receive()
//}

////FutureGetCFilterHeaderResult is a future promise to deliver the result of a
////GetCFilterHeaderAsync RPC invocation (or an applicable error).
//type FutureGetCFilterHeaderResult chan *response
//
//// Receive waits for the response promised by the future and returns the raw
//// filter header requested from the server given its block hash.
//func (r FutureGetCFilterHeaderResult) Receive() (*results.MsgCFHeaders, error) {
//	res, err := receiveFuture(r)
//	if err != nil {
//		return nil, err
//	}
//
//	// Unmarshal result as a string.
//	var headerHex string
//	err = json.Unmarshal(res, &headerHex)
//	if err != nil {
//		return nil, err
//	}
//
//	// Assign the decoded header into a hash
//	headerHash, err := NewHashFromStr(headerHex)
//	if err != nil {
//		return nil, err
//	}
//
//	// Assign the hash to a headers message and return it.
//	msgCFHeaders := MsgCFHeaders{PrevFilterHeader: *headerHash}
//	return &msgCFHeaders, nil
//
//}
//
//// GetCFilterHeaderCmd defines the getcfilterheader JSON-RPC command.
//type GetCFilterHeaderCmd struct {
//	Hash       string
//	FilterType FilterType
//}
//
//// GetCFilterHeaderAsync returns an instance of a type that can be used to get
//// the result of the RPC at some future time by invoking the Receive function
//// on the returned instance.
////
//// See GetCFilterHeader for the blocking version and more details.
//func (c *Client) GetCFilterHeaderAsync(blockHash *results.Hash,
//	filterType FilterType) futures.FutureGetCFilterHeaderResult {
//	hash := ""
//	if blockHash != nil {
//		hash = blockHash.String()
//	}
//
//	cmd := &GetCFilterHeaderCmd{
//		Hash:       hash,
//		FilterType: filterType,
//	}
//	return c.sendCmd(cmd)
//}
//
//// GetCFilterHeader returns a raw filter header from the server given its block
//// hash.
//func (c *Client) GetCFilterHeader(blockHash *results.Hash,
//	filterType FilterType) (*results.MsgCFHeaders, error) {
//	return c.GetCFilterHeaderAsync(blockHash, filterType).Receive()
//}
