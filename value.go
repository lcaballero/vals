package vals

import (
	"errors"
	"reflect"
	"strings"
)


type Value struct {
	key interface{}
	data interface{}
	previous *Value
	err error
}

func New(m interface{}) *Value {
	return &Value{
		data: m,
	}
}

func (v *Value) Val(key string) interface{} {
	mp, ok := v.data.(map[string]interface{})
	if !ok {
		v.err = errors.New("Couldn't assert data is map[string]interface{}")
		return nil
	} else {
		return mp[key]
	}
}

func (v *Value) At(k string) *Value {
	return &Value {
		key: k,
		data: v.Val(k),
		previous: v,
	}
}

func (v *Value) Data() interface{} {
	return v.data
}

func (v *Value) HasValue() bool {
	return v.data != nil
}

func (v *Value) IsArray() bool {
	t := reflect.TypeOf(v.data)
	return t != nil && t.Kind() == reflect.Array
}

func (v *Value) IsSlice() bool {
	t := reflect.TypeOf(v.data)
	return t != nil && t.Kind() == reflect.Slice
}

func (v *Value) IsMap() bool {
	t := reflect.TypeOf(v.data)
	return t != nil && t.Kind() == reflect.Map
}

func (v *Value) to(n int) interface{} {
	var data interface{} = nil
	if v.IsArray() || v.IsSlice() {
		actual := reflect.ValueOf(v.data)
		if 0 <= n && n <= actual.Len() {
			data = actual.Index(n).Interface()
		}
	}
	return data
}

func (v *Value) Len() int {
	if v.IsArray() || v.IsSlice() {
		actual := reflect.ValueOf(v.data)
		return actual.Len()
	}
	return -1
}

func (v *Value) In(n int) *Value {
	return &Value {
		key: n,
		data: v.to(n),
		previous: v,
	}
}

func (v *Value) AsString() string {
	s, ok := v.data.(string)
	if ok {
		return s
	} else {
		v.err = errors.New("Couldn't convert")
		return ""
	}
}

// Requires a pointer to fill with Value data
func (v *Value) Fill(b interface{}) {
	biz := reflect.ValueOf(b).Elem()
	n := biz.NumField()
	for i := 0; i < n; i++ {
		prop := biz.Type().Field(i).Name
		name := strings.ToLower(prop)
		val := v.At(name)
		if val.HasValue() {
			pv := biz.FieldByName(prop)
			pv.Set(reflect.ValueOf(val.Data()))
		}
	}
}
