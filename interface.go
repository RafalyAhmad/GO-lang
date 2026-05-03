package main

import "fmt"

type Persegi struct {
	sisi int
}

type BangunDatar interface {
	HitungLuas() int
}

func (p Persegi) HitungLuas() int {
	return p.sisi * p.sisi
}

func Cetak(b BangunDatar) {
	fmt.Println("Luasnya adalah", b.HitungLuas())
}

func main() {
	p1 := Persegi{sisi: 5}
	Cetak(p1)
}
