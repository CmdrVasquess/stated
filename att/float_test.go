package att

import (
	"encoding/json"
	"math"
	"testing"
)

func TestFloatF32_json(t *testing.T) {
	test := func(v float32) {
		cf := ToFloat32(v)
		data, err := json.Marshal(cf)
		if err != nil {
			t.Error(err)
		}
		var rf Float32
		if err = json.Unmarshal(data, &rf); err != nil {
			t.Error(err)
		}
		if math.IsNaN(float64(rf.Get())) && math.IsNaN(float64(v)) {
			return
		}
		if float32(rf.Get()) != v {
			t.Errorf("expect %f, got %f", v, rf)
		}
	}
	test(4711)
	test(float32(math.Inf(-1)))
	test(float32(math.Inf(1)))
	test(float32(math.NaN()))
}
