package main

import (
	"fmt"
	"time"
)

type Base struct{ ITime }

type Sub struct{ Base }

type RealTime struct{ ITime }

type FakeTime struct{ ITime }

func (t *RealTime) Time() uint32 {
	return uint32(time.Now().Unix()) // truncation is desired
}

func (t *FakeTime) Time() uint32 {
	// Monday, October 5, 2020 9:00:00 AM GMT-05:00
	return 1601906400
}

type ITime interface {
	Time() uint32
}

func (b *Base) AMethodThatUsesTime() {
	fmt.Println(b.ITime.Time())
}

func NewBase() *Base {
	return &Base{
		ITime: &RealTime{},
	}
}

func NewSub() *Sub {
	return &Sub{
		Base: Base{
			ITime: &FakeTime{},
		},
	}
}

func main() {
	base := NewBase()
	base.AMethodThatUsesTime()

	sub := NewSub()
	sub.AMethodThatUsesTime()
}

/*
Example output
go run .\overloading.go
1660769805
1601906400
*/
