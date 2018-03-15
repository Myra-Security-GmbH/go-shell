package output

import (
	"fmt"
	"strconv"
)

//
// AsNumber formats a number for better readability
//
func AsNumber(num float64, decimalPlaces byte) string {
	var ret string
	tmp := fmt.Sprintf("%.0f", num)

	c := 0
	for i := len(tmp) - 1; i >= 0; i-- {
		if c > 0 && c%3 == 0 {
			ret = "," + ret
		}

		ret = string(tmp[i]) + ret

		c++
	}

	if decimalPlaces > 0 {
		ret += fmt.Sprintf(
			"%."+strconv.FormatInt(int64(decimalPlaces), 10)+"f",
			num-float64(uint64(num)),
		)[1:]
	}

	return ret
}
