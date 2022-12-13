package arm

import (
	"gonum.org/v1/gonum/stat/distuv"
)

type Arm struct {
	P        float64
	Success  int
	Fail     int
	Binomial distuv.Binomial
}

func New(p float64) *Arm {
	return &Arm{
		P:        p,
		Success:  0,
		Fail:     0,
		Binomial: distuv.Binomial{N: 1, P: p},
	}
}

func (arm *Arm) Play() int {
	result := int(arm.Binomial.Rand())

	if result == 1 {
		arm.Success += 1
	} else {
		arm.Fail += 1
	}

	return result
}

func (arm *Arm) CalcSuccess() float64 {
	if arm.Success+arm.Fail == 0 {
		return 0.0
	} else {
		return float64(arm.Success) / float64(arm.Success+arm.Fail)
	}
}
