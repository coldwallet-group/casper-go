package deploy

import (
	"container/list"
	cl "github/casper-go/clvalue"
)

type Keyer interface {
	GetKey() string
	GetValue() *cl.CLValue
}

type MapList struct {
	dataMap  map[string]*list.Element
	dataList *list.List
}

func NewMapList() *MapList {
	return &MapList{
		dataMap:  make(map[string]*list.Element),
		dataList: list.New(),
	}
}

func (mapList *MapList) Exists(data Keyer) bool {
	_, exists := mapList.dataMap[string(data.GetKey())]
	return exists
}

func (mapList *MapList) Push(data Keyer) bool {
	if mapList.Exists(data) {
		return false
	}
	elem := mapList.dataList.PushBack(data)
	mapList.dataMap[data.GetKey()] = elem
	return true
}

func (mapList *MapList) Remove(data Keyer) {
	if !mapList.Exists(data) {
		return
	}
	mapList.dataList.Remove(mapList.dataMap[data.GetKey()])
	delete(mapList.dataMap, data.GetKey())
}

func (mapList *MapList) Size() int {
	return mapList.dataList.Len()
}

func (mapList *MapList) Walk(cb func(data Keyer)) {
	for elem := mapList.dataList.Front(); elem != nil; elem = elem.Next() {
		cb(elem.Value.(Keyer))
	}
}

type Elements struct {
	key   string
	value *cl.CLValue
}

func (e Elements) GetKey() string {
	return e.key
}

func (e Elements) GetValue() *cl.CLValue {
	return e.value
}
