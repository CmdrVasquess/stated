package att

import (
	"encoding/json"
	"math"
	"strconv"

	"github.com/fractalqb/change"
)

type Float32 struct {
	change.Val[float32]
}

func ToFloat32(v float32) Float32 {
	return Float32{change.NewVal(v)}
}

func (f Float32) MarshalJSON() ([]byte, error) {
	x := float64(f.Get())
	switch {
	case math.IsNaN(x):
		return json.Marshal("NaN")
	case math.IsInf(x, 1):
		return json.Marshal("+∞")
	case math.IsInf(x, -1):
		return json.Marshal("-∞")
	default:
		return strconv.AppendFloat(nil, x, 'f', -1, 32), nil
	}
}

func (f *Float32) UnmarshalJSON(data []byte) error {
	str := string(data)
	switch str {
	case `"NaN"`:
		*f = ToFloat32(float32(math.NaN()))
	case `"+∞"`:
		*f = ToFloat32(float32(math.Inf(1)))
	case `"-∞"`:
		*f = ToFloat32(float32(math.Inf(-1)))
	default:
		x, err := strconv.ParseFloat(str, 32)
		if err != nil {
			return err
		}
		*f = ToFloat32(float32(x))
	}
	return nil
}

type Float64 struct {
	change.Val[float64]
}

func ToFloat64(v float64) Float64 {
	return Float64{change.NewVal(v)}
}

func (f Float64) MarshalJSON() ([]byte, error) {
	x := f.Get()
	switch {
	case math.IsNaN(x):
		return json.Marshal("NaN")
	case math.IsInf(x, 1):
		return json.Marshal("+∞")
	case math.IsInf(x, -1):
		return json.Marshal("-∞")
	default:
		return strconv.AppendFloat(nil, x, 'f', -1, 64), nil
	}
}

func (v *Float64) UnmarshalJSON(data []byte) error {
	str := string(data)
	switch str {
	case `"NaN"`:
		*v = ToFloat64(math.NaN())
	case `"+∞"`:
		*v = ToFloat64(math.Inf(1))
	case `"-∞"`:
		*v = ToFloat64(math.Inf(-1))
	default:
		x, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return err
		}
		*v = ToFloat64(x)
	}
	return nil
}
