package main

import "fmt"

type ICloneable interface {
	Clone() ICloneable
}

type Base struct {
	baseValue int
}

type Sub struct {
	Base
	subValue int
}

type Sub2 struct {
	Base
	subValue int
}

func NewBase() *Base {
	return &Base{
		baseValue: 10,
	}
}

func NewSub() *Sub {
	return &Sub{
		Base:     *NewBase(),
		subValue: 20,
	}
}

func NewSub2() *Sub2 {
	return &Sub2{
		Base:     *NewBase(),
		subValue: 20,
	}
}

func (b *Base) Clone() ICloneable {
	return &Base{
		baseValue: b.baseValue,
	}
}

func (s *Sub) Clone() ICloneable {
	return &Sub{
		Base:     *s.Base.Clone().(*Base),
		subValue: s.subValue,
	}
}

func (s *Sub2) Clone() *Sub2 {
	return &Sub2{
		Base:     *s.Base.Clone().(*Base),
		subValue: s.subValue,
	}
}

func main() {
	// Cloneing a base
	{
		base := NewBase()

		clone := base.Clone().(*Base)
		clone.baseValue += 10

		fmt.Println("original base")
		fmt.Printf("base: %v", base.baseValue)
		fmt.Println()

		fmt.Println("clone")
		fmt.Printf("base: %v", clone.baseValue)
		fmt.Println()
	}
	fmt.Println()

	// Cloning a sub
	{
		sub := NewSub()

		clone := sub.Clone().(*Sub)
		clone.baseValue += 10
		clone.subValue += 10

		fmt.Println("original sub")
		fmt.Printf("base: %v, sub: %v", sub.baseValue, sub.subValue)
		fmt.Println()

		fmt.Println("clone")
		fmt.Printf("base: %v, sub: %v", clone.baseValue, clone.subValue)
		fmt.Println()
	}
	fmt.Println()

	// Is Sub2 ICloneable?
	{
		var sub2 ICloneable
		sub2 = NewSub2() // Nope
		_ = sub2
	}
}
