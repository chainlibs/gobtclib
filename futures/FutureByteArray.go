package futures

import (
	"encoding/json"
	"github.com/chainlibs/gobtclib/base"
)

/*
Description:
FutureByteArray is a future promise to deliver the result of []Byte type.
 * Author: architect.bian
 * Date: 2018/10/15 10:57
 */
type FutureByteArray chan *base.Response

/*
Description:
Receive waits for the response promised by the future and returns the
result of Byte type.
 * Author: architect.bian
 * Date: 2018/10/15 10:57
 */
func (f FutureByteArray) Receive() ([]byte, error) {
	res, err := ReceiveFuture(f)
	if err != nil {
		return nil, err
	}

	// Unmarshal the result as a Byte.
	var result []byte
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
