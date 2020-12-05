package att

import (
	"encoding/json"
	"math"
	"strconv"
)

type Float32 float32

func (f *Float32) Set(v float32, chg Change) Change {
	if *f == Float32(v) {
		return 0
	}
	*f = Float32(v)
	return chg
}

func (f Float32) Get() float32 { return float32(f) }

func (f Float32) MarshalJSON() ([]byte, error) {
	x := float64(f)
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
		*f = Float32(math.NaN())
	case `"+∞"`:
		*f = Float32(math.Inf(1))
	case `"-∞"`:
		*f = Float32(math.Inf(-1))
	default:
		x, err := strconv.ParseFloat(str, 32)
		if err != nil {
			return err
		}
		*f = Float32(x)
	}
	return nil
}

type Float64 float64

func (f *Float64) Set(v float64, chg Change) Change {
	if *f == Float64(v) {
		return 0
	}
	*f = Float64(v)
	return chg
}

func (f Float64) Get() float64 { return float64(f) }

func (f Float64) MarshalJSON() ([]byte, error) {
	x := float64(f)
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

func (f *Float64) UnmarshalJSON(data []byte) error {
	str := string(data)
	switch str {
	case `"NaN"`:
		*f = Float64(math.NaN())
	case `"+∞"`:
		*f = Float64(math.Inf(1))
	case `"-∞"`:
		*f = Float64(math.Inf(-1))
	default:
		x, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return err
		}
		*f = Float64(x)
	}
	return nil
}
