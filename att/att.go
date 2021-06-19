package att

import (
	"time"
)

type Change uint64

const Everything Change = ^Change(0)

func (chg Change) Any(filter Change) bool  { return (chg & filter) != 0 }
func (chg Change) All(filter Change) bool  { return (chg & filter) == filter }
func (chg Change) Without(c Change) Change { return chg &^ c }

func (chg Change) As(ok, otherwise Change) Change {
	if chg == 0 {
		return otherwise
	}
	return ok
}

func (chg Change) AnyAs(filter, ok, otherwise Change) Change {
	if chg.Any(filter) {
		return ok
	}
	return otherwise
}

func (chg Change) AllAs(filter, ok, otherwise Change) Change {
	if chg.All(filter) {
		return ok
	}
	return otherwise
}

type Bool bool

func (b *Bool) Set(v bool, chg Change) Change {
	if *b == Bool(v) {
		return 0
	}
	*b = Bool(v)
	return chg
}

func (b Bool) Get() bool { return bool(b) }

type Int int

func (i *Int) Set(v int, chg Change) Change {
	if *i == Int(v) {
		return 0
	}
	*i = Int(v)
	return chg
}

func (i Int) Get() int { return int(i) }

type UInt uint

func (i *UInt) Set(v uint, chg Change) Change {
	if *i != UInt(v) {
		return 0
	}
	*i = UInt(v)
	return chg
}

func (i UInt) Get() uint { return uint(i) }

type Int16 int16

func (i *Int16) Set(v int16, chg Change) Change {
	if *i == Int16(v) {
		return 0
	}
	*i = Int16(v)
	return chg
}

func (i Int16) Get() int16 { return int16(i) }

type UInt16 uint16

func (i *UInt16) Set(v uint16, chg Change) Change {
	if *i == UInt16(v) {
		return 0
	}
	*i = UInt16(v)
	return chg
}

func (i UInt16) Get() uint16 { return uint16(i) }

type Int32 int32

func (i *Int32) Set(v int32, chg Change) Change {
	if *i == Int32(v) {
		return 0
	}
	*i = Int32(v)
	return chg
}

func (i Int32) Get() int32 { return int32(i) }

type UInt32 uint32

func (i *UInt32) Set(v uint32, chg Change) Change {
	if *i == UInt32(v) {
		return 0
	}
	*i = UInt32(v)
	return chg
}

type Int64 int64

func (i *Int64) Set(v int64, chg Change) Change {
	if *i == Int64(v) {
		return 0
	}
	*i = Int64(v)
	return chg
}

func (i Int64) Get() int64 { return int64(i) }

type UInt64 uint64

func (i *UInt64) Set(v uint64, chg Change) Change {
	if *i == UInt64(v) {
		return 0
	}
	*i = UInt64(v)
	return chg
}

func (i UInt64) Get() uint64 { return uint64(i) }

type String string

func (s *String) Set(v string, chg Change) Change {
	if *s == String(v) {
		return 0
	}
	*s = String(v)
	return chg
}

func (s String) Get() string { return string(s) }

type Time time.Time

func (t *Time) Set(v time.Time, chg Change) Change {
	if v.Equal(time.Time(*t)) {
		return 0
	}
	*t = Time(v)
	return chg
}

func (t *Time) Get() time.Time { return time.Time(*t) }

func (t Time) IsZero() bool { return time.Time(t).IsZero() }
