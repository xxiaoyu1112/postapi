package view_model

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {
	a := `{     "feature": [         [             499.75,             120.180503,             30.323697,             2.0,             2.0194721229177763,             118.0,             160.25,             71.86666666666667         ],         [             499.75,             120.180503,             30.323697,             2.0,             2.0194721229177763,             118.0,             160.25,             71.86666666666667         ],         [             499.75,             120.180503,             30.323697,             2.0,             2.0194721229177763,             118.0,             160.25,             71.86666666666667         ],         [             477.63333333333327,             120.180503,             30.323697,             2.0,             2.0194721229177763,             118.0,             182.36666666666673,             71.86666666666667         ],         [             563.3,             120.180031,             30.322666,             2.0,             2.823838137978246,             165.0,             119.98333333333335,             95.14999999999998         ],         [             474.15,             120.180031,             30.322666,             2.0,             2.823838137978246,             165.0,             185.85000000000002,             71.86666666666667         ],         [             527.6,             120.180031,             30.322666,             2.0,             2.823838137978246,             165.0,             252.39999999999998,             191.86666666666667         ],         [             496.4,             120.178986,             30.32284000000001,             2.0,             2.3104130219822014,             135.0,             163.60000000000002,             71.86666666666667         ]     ],     "start": 2 }`
	pt := PredictInput{}

	json.Unmarshal([]byte(a), &pt)
	fmt.Print(pt)
}
