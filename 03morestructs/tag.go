package morestructs

import (
	"log"
	"reflect"
	"strconv"
)

type Person struct {
	// <field> <type> `<key:"meta">`
	Name string `json:"name" validate:"required" gorm:"column:name"`
}

type MapStruct struct {
	Str     string  `map:"str"`
	StrPtr  *string `map:"str"`
	Bool    bool    `map:"bool"`
	BoolPtr *bool   `map:"bool"`
	Int     int     `map:"int"`
	IntPtr  *int    `map:"int"`
}

func mapStruct() {
	src := map[string]string{
		"str":  "string data",
		"bool": "true",
		"int":  "12345",
	}
	var ms MapStruct
	Decode(&ms, src)
	log.Println(ms)
}

func Decode(target *MapStruct, src map[string]string) error {
	v := reflect.ValueOf(target)
	e := v.Elem()
	return decode(e, src)
}

func decode(e reflect.Value, src map[string]string) error {
	t := e.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		// 埋め込まれた構造体は再帰処理
		if f.Anonymous {
			if err := decode(e.Field(i), src); err != nil {
				return err
			}
			continue
		}

		// 子供が構造体だったら再帰処理
		if f.Type.Kind() == reflect.Struct {
			if err := decode(e.Field(i), src); err != nil {
				return err
			}
			continue
		}

		// タグがなければフィールド名をそのまま使う
		key := f.Tag.Get("map")
		if key == "" {
			key = f.Name
		}

		// 元データになければスキップ
		sv, ok := src[key]
		if !ok {
			continue
		}

		// フィールドの型を取得
		var k reflect.Kind
		var isP bool
		if f.Type.Kind() != reflect.Ptr {
			k = f.Type.Kind()
		} else {
			k = f.Type.Elem().Kind()
			// ポインターのポインターは無視
			if k == reflect.Ptr {
				continue
			}
			isP = true
		}

		switch k {
		case reflect.String:
			if isP {
				e.Field(i).Set(reflect.ValueOf(&sv))
			} else {
				e.Field(i).SetString(sv)
			}
		case reflect.Bool:
			b, err := strconv.ParseBool(sv)
			if err != nil {
				continue
			}
			if isP {
				e.Field(i).Set(reflect.ValueOf(&b))
			} else {
				e.Field(i).SetBool(b)
			}
		case reflect.Int:
			n64, err := strconv.ParseInt(sv, 10, 64)
			if err != nil {
				continue
			}
			if isP {
				e.Field(i).Set(reflect.ValueOf(&n64))
			} else {
				e.Field(i).SetInt(n64)
			}
		}
	}
	return nil
}
