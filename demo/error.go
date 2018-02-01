package main

import (
	"fmt"
	"math"
	"time"
)

type SError struct {
	ErrMsg string
	time   time.Time
}

func (err *SError) Error() string {
	return fmt.Sprintf("at %v --> %s", err.ErrMsg, err.time)
}

func main() {
	if f, e := Sqrt2(-1.0); e != nil {
		fmt.Println(f)

	} else {
		fmt.Println(e)
	}

}

func Sqrt2(x float64) (float64, error) {
	var res float64
	if x >= 0 {
		return Sqrt(x), nil
	} else {
		return res, &SError{"负数不能开方", time.Time{}}
	}

}

func Sqrt(x float64) float64 {
	z := 1.0
	for {
		tmp := z - (z*z-x)/(2*z)
		fmt.Println(tmp)
		if tmp == z || math.Abs(tmp-z) < 0.000000000001 {
			break
		}
		z = tmp
	}
	return z
}
