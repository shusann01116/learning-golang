package structs

import (
	"log/slog"
	"slices"
	"strconv"
	"time"
)

// コンシューマ型
type Consumer struct {
	Name      string
	Active    bool
	ExpiredAt time.Time
}

// Slice に対して型を定義する
type Consumers []Consumer

func (c Consumers) RequiredFollows() Consumers {
	return c.activeConsumer().expires(time.Now().AddDate(0, 0, 10)).sortByExpiredAt()
}

// 定義した型に対して Func を定義することでロジックをレポジトリに混入させない
// フィルタリング操作などなど
func (c Consumers) activeConsumer() Consumers {
	resp := make([]Consumer, 0, len(c))
	for _, v := range c {
		if v.Active {
			resp = append(resp, v)
		}
	}
	return resp
}

// 戻り値を `Consumers` で揃えておくことでメソッドチェーンが作れる
func (c Consumers) expires(end time.Time) Consumers {
	resp := make(Consumers, 0, len(c))
	for _, v := range c {
		if end.After(v.ExpiredAt) {
			resp = append(resp, v)
		}
	}
	return resp
}

// see: https://pkg.go.dev/slices#SortFunc for impl
func (c Consumers) sortByExpiredAt() Consumers {
	slices.SortFunc(c, func(a, b Consumer) int {
		return a.ExpiredAt.Compare(b.ExpiredAt)
	})
	return c
}

func ExampleActiveConsumer() {
	consumers, err := GetConsumers()
	if err != nil { //nolint:staticcheck,revive
		// error handling
	}

	requiredFollows := consumers.RequiredFollows()
	slog.Info(strconv.Itoa(len(requiredFollows)))
}

func GetConsumers() (Consumers, error) {
	panic("not implemented")
}
