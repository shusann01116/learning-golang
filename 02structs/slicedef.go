package structs

import "time"

// コンシューマ型
type Consumer struct {
	Name      string
	Active    bool
	ExpiredAt time.Time
}

// Slice に対して型を定義する
type Consumers []Consumer

// 定義した型に対して Func を定義することでロジックをレポジトリに混入させない
// フィルタリング操作などなど
func (c Consumers) ActiveConsumer() Consumers {
	resp := make([]Consumer, 0, len(c))
	for _, v := range c {
		if v.Active {
			resp = append(resp, v)
		}
	}
	return resp
}

// 戻り値を `Consumers` で揃えておくことでメソッドチェーンが作れる
func (c Consumers) Exipres(end time.Time) Consumers {
	resp := make(Consumers, 0, len(c))
	for _, v := range c {
		if end.After(v.ExpiredAt) {
			resp = append(resp, v)
		}
	}
	return resp
}

func exampleActiveConsumer() {
	consumers, err := GetConsumers()
	if err != nil {
		// error handling
	}

	activeConsumers := consumers.ActiveConsumer().Exipres(time.Now().AddDate(0, 0, -1).Truncate(24 * time.Hour))
	println(len(activeConsumers))
}

func GetConsumers() (Consumers, error) {
	panic("not implemented")
}
