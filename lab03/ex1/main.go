package main

import (
	"fmt"
	"math/rand"
	"time"
)

// stocheaza deviatia fata de timpul real
type Clock struct{ offset time.Duration }

func (c *Clock) Now() time.Time {
	return time.Now().Add(c.offset)
}

// TODO:
// Calculeaza decalajul theta si round -trip -ul delta

func estimate(t1, t2, t3, t4 time.Time) (theta, delta time.Duration) {
	theta = (t2.Sub(t1) + t3.Sub(t4)) / 2
	delta = t4.Sub(t1) - t3.Sub(t2)
	return theta, delta
}

func measure(c_clock, s_clock *Clock) (theta, delta time.Duration) {
	// Request
	t1 := c_clock.Now()
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	t2 := s_clock.Now()
	// Procesare intre 5 si 20 ms
	time.Sleep(5*time.Millisecond + time.Duration(rand.Intn(20))*time.Millisecond)
	t3 := s_clock.Now()
	// Response
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	t4 := c_clock.Now()
	return estimate(t1, t2, t3, t4)
}

func main() {
	c_clock := &Clock{offset: 0}
	s_clock := &Clock{offset: 1000 * time.Millisecond}

	theta, delta := measure(c_clock, s_clock)
	fmt.Printf("inainte: theta = %v, delta = %v\n", theta, delta)

	c_clock.offset += theta

	theta_a, delta_a := measure(c_clock, s_clock)
	fmt.Printf("dupa:    theta = %v, delta = %v\n", theta_a, delta_a)

	fmt.Println("")
	c_clock_b := &Clock{offset: 0}
	var theta_min, delta_min time.Duration
	for i := 0; i < 10; i++ {
		theta_b, delta_b := measure(c_clock_b, s_clock)
		fmt.Printf("%2d: theta = %v, delta = %v\n", i+1, theta_b, delta_b)
		if i == 0 || delta_b < delta_min {
			theta_min, delta_min = theta_b, delta_b
		}
	}
	fmt.Printf("theta = %v, delta = %v \n", theta_min, delta_min)


	fmt.Printf("delta : (a) %v vs (b) %v\n", delta, delta_min)
	fmt.Printf("offset :  (a) %v vs (b) %v\n", theta_a, theta_min)
}
