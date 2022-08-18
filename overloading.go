package main

import (
	"fmt"
	"time"
)

type Base struct{ ITime }

type Sub struct{ Base }

func (b *Base) Time() uint32 {
	return uint32(time.Now().Unix()) // truncation is desired
}

func (b *Sub) Time() uint32 {
	// Monday, October 5, 2020 9:00:00 AM GMT-05:00
	return 1601906400
}

type ITime interface {
	Time() uint32
}

func (b *Base) GetITime() uint32 {
	return b.ITime.Time()
}

func (b *Base) AFuncThatUsesTime() {
	fmt.Println(b.GetITime())
}

func NewBase() *Base {
	base := &Base{}
	base.ITime = interface{}(base).(ITime)
	return base
}

func NewSub() *Sub {
	sub := &Sub{}
	sub.ITime = interface{}(sub).(ITime)
	return sub
}

func main() {
	base := NewBase()
	base.AFuncThatUsesTime()

	sub := NewSub()
	sub.AFuncThatUsesTime()
}

/*
Example output
go run .\overloading.go
1660769805
1601906400
*/
