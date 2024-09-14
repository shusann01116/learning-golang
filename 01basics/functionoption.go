package basics

import (
	"fmt"
	"log/slog"
)

type Portion int

const (
	Regular Portion = iota
	Small
	Large
)

type Udon struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

// 愚直なパターン、結構大変
func NewUdon1(p Portion, aburaage bool, ebiten uint) *Udon {
	return &Udon{
		men:      p,
		aburaage: aburaage,
		ebiten:   ebiten,
	}
}

var TempuraUdon = NewUdon1(Large, false, 2)

// Option パターン
// オプション引数を使ったパターン Fill Struct で構造体のメンバーを埋めることができるので
// 認知負荷が比較的低く、それ自体がドキュメンテーションの役割を果たす
type Option struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

func NewUdon2(o Option) *Udon {
	return &Udon{
		men:      o.men,
		aburaage: o.aburaage,
		ebiten:   o.ebiten,
	}
}

// builder パターン
// IDE による補完が効くため生産性がよいパターン
type FluentOpt struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

func NewUdon3(p Portion) *FluentOpt {
	return &FluentOpt{
		men:      p,
		aburaage: false,
		ebiten:   1,
	}
}

func (o *FluentOpt) Aburaage() *FluentOpt {
	o.aburaage = true
	return o
}

func (o *FluentOpt) Ebiten(n uint) *FluentOpt {
	o.ebiten = n
	return o
}

func (o *FluentOpt) Order() *Udon {
	return &Udon{
		men:      o.men,
		aburaage: o.aburaage,
		ebiten:   o.ebiten,
	}
}

func UseFluentInterface() {
	oomoriKitsune := NewUdon3(Large).Aburaage().Order()
	slog.Info(fmt.Sprint(oomoriKitsune))
}

// Functiona Option パターン
// 独立した関数が加工するパターン、戻り地の関数は型定義しておくことで go doc にまとまる
type OptFunc func(r *Udon)

func NewUdon4(opts ...OptFunc) *Udon {
	r := &Udon{
		men:      0,
		aburaage: false,
		ebiten:   0,
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

func OptMen(p Portion) OptFunc {
	return func(r *Udon) {
		r.men = p
	}
}

func OptAburaage() OptFunc {
	return func(r *Udon) {
		r.aburaage = true
	}
}

func OptEbiten(n uint) OptFunc {
	return func(r *Udon) {
		r.ebiten = n
	}
}
