package futures

import (
	"encoding/json"
	"github.com/chainlibs/gobtclib/base"
)

/*
Description:
FutureStringArray is a future promise to deliver the result of []string type.
 * Author: architect.bian
 * Date: 2018/10/15 10:56
 */
type FutureStringArray chan *base.Response

/*
Description:
Receive waits for the response promised by the future and returns the
result of string type.
 * Author: architect.bian
 * Date: 2018/10/15 10:56
 */
func (f FutureStringArray) Receive() (*[]string, error) {
	res, err := ReceiveFuture(f)
	if err != nil {
		return nil, err
	}

	// Unmarshal the result as a string.
	var result []string
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
