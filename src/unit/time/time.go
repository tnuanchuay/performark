package time

import (
	"strings"
	"strconv"
	"math"
)

const (
	us	=	1e-6
	ms	=	1e-3
	s	=	1
	m	=	60
	h	=	3600
)
var listUnits = []string{"us", "ms", "s", "m", "h"}
var units = map[string]float64{
	"us":us,
	"ms":ms,
	"s":s,
	"m":m,
	"h":h,
}

func StringToInt(s string) (float64, error){
	var result float64 = 0
	for _, unit := range listUnits{
		if strings.Contains(s, unit){
			value := strings.Split(s, unit)
			result, err := strconv.ParseFloat(value[0], 64)
			if err != nil {
				return result, error(err)
			}

			result = result * units[unit]
			result = math.Floor(result)
			return result, nil
		}
	}
	return result, nil
}
