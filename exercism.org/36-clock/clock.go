package clock

import (
    "math"
    "fmt"
)

type Clock struct {
    Minutes int
    Hours int
}

func convMins(m float64) (int, int) {
    var tmp, mFrac, h float64
    var res int
    
    if m >= 60  || m <= -60 {
    	tmp = float64(m) / float64(60)
        h,mFrac = math.Modf(tmp)
        res = int(math.Round(mFrac * float64(60)))
    } else {
    	res = int(m)
    }

    if res < 0 {
        res = 60 + res
        h -= 1
    }

    return int(h), res
}

func convHours(h float64) int {
    var tmp, hFrac float64
    var res int
    
    if h >= 24 || h <= -24 {
        tmp = float64(h) / float64(24)
        _,hFrac = math.Modf(tmp)
        res = int(math.Round(hFrac * float64(24)))
    } else {
    	res = int(h)
    }

    if res < 0 {
        res = 24 + res
    }

    return res
}

func New(h, m int) Clock {
    convH, convM := convMins(float64(m))
    convH = convHours(float64(h + convH))

	return Clock{
        Minutes: convM,
        Hours: convH,
    }
}

func (c Clock) Add(m int) Clock {
    convH, convM := convMins(float64(c.Minutes + m))
    convH = convHours(float64(c.Hours + convH))

    return Clock{
        Minutes: convM,
        Hours: convH,
    }
}

func (c Clock) Subtract(m int) Clock {
	convH, convM := convMins(float64(c.Minutes + -(m)))
    convH = convHours(float64(c.Hours + convH))

    return Clock{
        Minutes: convM,
        Hours: convH,
    }
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hours, c.Minutes)
}
