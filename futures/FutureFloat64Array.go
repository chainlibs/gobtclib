package futures

import (
	"encoding/json"
	"github.com/chainlibs/gobtclib/base"
)

/*
Description:
FutureFloat64Array is a future promise to deliver the result of []float64 type.
 * Author: architect.bian
 * Date: 2018/10/15 10:57
 */
type FutureFloat64Array chan *base.Response

/*
Description:
Receive waits for the response promised by the future and returns the
result of float64 type.
 * Author: architect.bian
 * Date: 2018/10/15 10:57
 */
func (f FutureFloat64Array) Receive() ([]float64, error) {
	res, err := ReceiveFuture(f)
	if err != nil {
		return nil, err
	}

	// Unmarshal the result as a float64.
	var result []float64
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
