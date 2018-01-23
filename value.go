package vals

import (
	"errors"
	"reflect"
	"strings"
)

// A value instance provides an API over generic data.  The functions
// At() and In() push new Values onto the end of the chain, drilling
// down to the desired value.
type Value struct {
	key      interface{}
	data     interface{}
	previous *Value
	err      error
}

// Creates a new instance with the given 'backing' value.
func New(m interface{}) *Value {
	return &Value{
		data: m,
	}
}

// Attempts to access the value as though it is a map using the given key.
// If the backing data is not a map this method returns nil.
func (v *Value) val(key string) interface{} {
	mp, ok := v.data.(map[string]interface{})
	if ok {
		return mp[key]
	} else {
		v.err = errors.New("Couldn't assert data is map[string]interface{}")
		return nil
	}
}

// Access the backing data which this value now contains.
func (v *Value) Data() interface{} {
	return v.data
}

// Returns true if the backing data is non-nil
func (v *Value) HasValue() bool {
	return v.data != nil
}

// Returns true iff backing data is non-nil and it is an array.
func (v *Value) IsArray() bool {
	t := reflect.TypeOf(v.data)
	return t != nil && t.Kind() == reflect.Array
}

// Returns true iff backing data is non-nil and it is an slice.
func (v *Value) IsSlice() bool {
	t := reflect.TypeOf(v.data)
	return t != nil && t.Kind() == reflect.Slice
}

// Returns true iff backing data is non-nil and it is an map.
func (v *Value) IsMap() bool {
	t := reflect.TypeOf(v.data)
	return t != nil && t.Kind() == reflect.Map
}

// Provides the value at the given index as though the backing
// data is an array or slice (nil if the data is not a slice or array).
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

// If data is an array or slice then this function returns the
// size of the array or slice, otherwise -1
func (v *Value) Len() int {
	if v.IsArray() || v.IsSlice() {
		actual := reflect.ValueOf(v.data)
		return actual.Len()
	}
	return -1
}

// Saves the index, and loads the value for that index as the backing store
// if possible.  If the index doesn't exist the new data will be nil.
func (v *Value) In(n int) *Value {
	return &Value{
		key:      n,
		data:     v.to(n),
		previous: v,
	}
}

// Saves the key, and loads the value for that key from the backing store
// if possible.  If the key doesn't exist the new Data/value will be nil.
func (v *Value) At(key string) *Value {
	return &Value{
		key:      key,
		data:     v.val(key),
		previous: v,
	}
}

// Attempts to cast data to a string, if data cannot be cast to a string
// this function returns "" else the value as cast to string.
func (v *Value) AsString() string {
	s, ok := v.data.(string)
	if ok {
		return s
	} else {
		v.err = errors.New("Couldn't convert")
		return ""
	}
}

// Requires a pointer to fill with Value data.
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
