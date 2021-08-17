package loader

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type Loader struct {
	Header []string
	Types  []string
	Data   [][]string

	HeaderIndex map[string]int
}

func New(filename string) (*Loader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	_, _ = reader.Read()

	header, err := reader.Read()
	if err != nil {
		return nil, err
	}
	types, err := reader.Read()
	if err != nil {
		return nil, err
	}
	data, err := reader.ReadAll()

	headerIndex := make(map[string]int)
	for i, value := range header {
		if len(value) == 0 {
			continue
		}
		headerIndex[value] = i
	}

	return &Loader{
		Header:      header,
		Types:       types,
		Data:        data,
		HeaderIndex: headerIndex,
	}, nil
}

func (ld *Loader) Size() int {
	return len(ld.Data)
}

func (ld *Loader) HeaderSize() int {
	return len(ld.Header)
}

func (ld *Loader) GetString(row int, name string) string {
	index, found := ld.HeaderIndex[name]
	if !found {
		return ""
	}
	return ld.Data[row][index]
}

func (ld *Loader) GetInt(row int, name string) int {
	index, found := ld.HeaderIndex[name]
	if !found {
		return 0
	}
	val, err := strconv.Atoi(ld.Data[row][index])
	if err != nil {
		return 0
	}
	return val
}

func (ld *Loader) GetInts(row int, name string) []int {
	now := 0
	ret := []int{}
	for {
		key := fmt.Sprintf("%s[%d]", name, now)
		index, found := ld.HeaderIndex[key]
		if !found {
			break
		}
		val, err := strconv.Atoi(ld.Data[row][index])
		if err != nil {
			val = 0
		}
		ret = append(ret, val)
		now++
	}
	return ret
}

func (ld *Loader) GetBool(row int, name string) bool {
	index, found := ld.HeaderIndex[name]
	if !found {
		return false
	}
	return ld.Data[row][index] == "True"
}

func (ld *Loader) Load(row int, ptr interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	refPtr := reflect.ValueOf(ptr)
	refValue := refPtr.Elem()
	refType := refValue.Type()
	numField := refValue.NumField()
	for i := 0; i < numField; i++ {
		fieldType := refType.Field(i)
		name := fieldType.Tag.Get("name")
		if len(name) == 0 {
			continue
		}
		fieldPtr := refValue.Field(i).Addr().Interface()
		switch obj := fieldPtr.(type) {
		case *string:
			*obj = ld.GetString(row, name)
		case *int:
			*obj = ld.GetInt(row, name)
		case *bool:
			*obj = ld.GetBool(row, name)
		case *[]int:
			*obj = ld.GetInts(row, name)
		}
	}

	return nil
}
