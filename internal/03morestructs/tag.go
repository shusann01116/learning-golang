package morestructs

import (
	"log"
	"reflect"
	"strconv"
	"unsafe"
)

type Person struct {
	// <field> <type> `<key:"meta">`
	Name string `json:"name" validate:"required" gorm:"column:name"`
}

type MapStruct struct {
	Str     string  `map:"str"`
	StrPtr  *string `map:"strPtr"`
	Bool    bool    `map:"bool"`
	BoolPtr *bool   `map:"boolPtr"`
	Int     int     `map:"int"`
	IntPtr  *int    `map:"intPtr"`
}

func MapStructSample() {
	src := map[string]string{
		"str":     "string data",
		"strPtr":  "string data",
		"bool":    "true",
		"boolPtr": "true",
		"int":     "12345",
		"intPtr":  "12345",
	}
	var ms MapStruct
	_ = Decode(&ms, src)
	log.Println(ms)
}

func Decode(target *MapStruct, src map[string]string) error {
	v := reflect.ValueOf(target)
	e := v.Elem()
	return decode(e, src)
}

//nolint:all
func decode(e reflect.Value, src map[string]string) error {
	t := e.Type()
	for i := range t.NumField() {
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
		default:
			panic("not implemented")
		}
	}
	return nil
}

//nolint:all
func Encode(target map[string]string, src interface{}) error {
	v := reflect.ValueOf(src)
	e := v.Elem()
	t := e.Type()
	for i := range t.NumField() {
		f := t.Field(i)

		if f.Anonymous {
			_ = Encode(target, e.Field(i).Addr().Interface())
			continue
		}

		key := f.Tag.Get("map")
		if key == "" {
			key = f.Name
		}

		if f.Type.Kind() == reflect.Struct {
			_ = Encode(target, e.Field(i).Addr().Interface())
			continue
		}

		var k reflect.Kind
		var isP bool
		if f.Type.Kind() != reflect.Ptr {
			k = f.Type.Kind()
		} else {
			k = f.Type.Elem().Kind()
			isP = true
			if k == reflect.Ptr {
				continue
			}
		}

		switch k {
		case reflect.String:
			if isP {
				if e.Field(i).Pointer() != 0 {
					target[key] = *(*string)(unsafe.Pointer(e.Field(i).Pointer()))
				}
			} else {
				target[key] = e.Field(i).String()
			}
		case reflect.Bool:
			var b bool
			if isP {
				if e.Field(i).Pointer() != 0 {
					b = *(*bool)(unsafe.Pointer(e.Field(i).Pointer()))
				}
			} else {
				b = e.Field(i).Bool()
			}
			target[key] = strconv.FormatBool(b)
		case reflect.Int:
			var n int64
			if isP {
				if e.Field(i).Pointer() != 0 {
					n = int64(*(*int)(unsafe.Pointer(e.Field(i).Pointer())))
				}
			} else {
				n = e.Field(i).Int()
			}
			target[key] = strconv.FormatInt(n, 10)
		default:
			panic("not implemented")
		}
	}
	return nil
}
