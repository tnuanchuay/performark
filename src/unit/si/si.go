package si

import (
	"strings"
	"strconv"
)

const (
	k	=	1000.0
	M	=	1000000.0
	G	=	1000000000.0
)

var listUnits = []string{"k", "M", "G"}
var units = map[string]float64{
	"k":k,
	"M":M,
	"G":G,
}

func SIToFloat(s string)(float64, error){
	var result float64
	ss := string(s)
	for _, unit := range listUnits{
		if strings.Contains(ss, unit) {
			v := strings.Split(ss, unit)[0]
			vv, err := strconv.ParseFloat(v, 64)
			if err != nil{
				return 0, err
			}else{
				result = float64(vv) * units[unit]
				return result, nil
			}
		}
	}
	result, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
