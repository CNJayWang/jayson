package jayson

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// 定义一些类型错误信息，用与在读写数据时验证返回信息
var (
	ErrNoNull             = errors.New("is not null")
	ErrNotArray           = errors.New("Not an array")
	ErrNotNumber          = errors.New("not a number")
	ErrNotBool            = errors.New("not a bool")
	ErrNotJSONObject      = errors.New("not and JSONObject")
	ErrNotJSONObjectArray = errors.New("not an JSONObject array")
	ErrNotString          = errors.New("not a string")
)

//定义一个T来标识一个泛型
type T interface{}

//定义一个Json Object Type
type JSONObjectType map[string]*DataNode

//定义一个Data的Node
type DataNode struct {
	data   T
	exists bool
}

// ey不存在的错误
type KeyNotFundError struct {
	Key string
}

//key不存在的错误
func (k KeyNotFundError) Error() string {
	if k.Key != "" {
		return fmt.Sprint("key '%s' not found", k.Key)
	}

	return KEY_NOT_FOUND
}

//DataNode节点的数据为Nil时获取节点数据
func (node *DataNode) Null() error {
	valid := false
	switch node.data.(type) {
	case nil:
		valid = node.exists
	}

	if valid {
		return nil
	}

	return ErrNoNull
}

//数据节点的数据类型为数组时，获取数据
func (node *DataNode) Array() ([]*DataNode, error) {
	valid := false

	switch node.data.(type) {
	case []T:
		valid = true
	}

	var slice []*DataNode

	if valid {
		for _, element := range node.data.([]T) {
			child := DataNode{element, true}
			slice = append(slice, &child)
		}
		return slice, nil
	}
	return slice, ErrNotArray
}

//数据节点的数据类型为json.Number类型时，获取数据
func (node *DataNode) Number() (json.Number, error) {
	valid := false

	switch node.data.(type) {
	case json.Number:
		valid = true
	}
	if valid {
		return node.data.(json.Number), nil
	}
	return nil, ErrNotNumber
}

//数据节点的数据类型为float64时，获取数据
func (node *DataNode) Float64() (float64, error) {
	n, err := node.Number()

	if err != nil {
		return 0, err
	}

	return n.Float64()
}

//数据节点类型为int64时，获取节点数据
func (node *DataNode) Int64() (int64, error) {
	n, err := node.Number()

	if err != nil {
		return 0, err
	}

	return n.Int64()
}

//数据节点类型为bool类型，获取节点数据
func (node *DataNode) Boolean() (bool, error) {

	valid := false

	switch node.data.(type) {
	case bool:
		valid = true
	}

	if valid {
		return node.data.(bool), nil
	}

	return false, ErrNotBool
}

//DataNode节点的数据类型为JSONObjectType类型
func (node *DataNode) JSONObject() (JSONObjectType, error) {
	var valid bool

	//检查数据的类型
	switch node.data.(type) {
	case JSONObjectType:
		valid = true
		break

	}

	if valid {

		m := make(JSONObjectType)

		for key, element := range node.data.(JSONObjectType) {
			m[key] = &DataNode{element, true}
		}
		return m, nil
	}
	return nil, ErrNotJSONObject
}

//数据节点为JsonObjectType的数组类型，获取节点数据
func (node *DataNode) JSONObjectArray() ([]JSONObjectType, error) {
	valid := false

	switch node.data.(type) {
	case JSONObjectType:
		valid = true
		break
	}

	var slice []JSONObjectType

	if valid {
		for _, element := range node.data.([]JSONObjectType) {
			slice = append(slice, element)
		}
		return slice, nil
	}

	return nil, ErrNotJSONObjectArray
}

//数据节点为string类型
func (node *DataNode) String() (string, error) {
	valid := false

	switch node.data.(type) {
	case string:
		valid = true
	}

	if valid {
		return node.data.(string), nil
	}
	//类型错误时返回空字符串
	return EMPTY_STRING, ErrNotString
}

//Json节点的泛型数据
func (node *DataNode) T() T {
	return node.data
}

//反序列号byte组为DataNode
func UnMarshaBytes(b []byte) (*DataNode, error) {
	r := bytes.NewReader(b)
	return NewDataNodefromReader(r)
}

//序列化DataNode节点的数据
func (node *DataNode) Marshal() ([]byte, error) {
	return json.Marshal(node.data)
}

// 获取子DataNode
func (node *DataNode) getChildNode(key string) (*DataNode, error) {
	obj, err := node.JSONObject()

	if err == nil {
		child, ok := obj[key]
		if ok {
			return child, nil
		} else {
			return nil, KeyNotFundError{key}
		}
	}

	return nil, err
}

//private方法,遍历Node节点
func (node *DataNode) walkNde(keys []string) (*DataNode, error) {
	currentNode := node
	var err error
	for _, key := range keys {
		currentNode, err = currentNode.getChildNode(key)
		if err != nil {
			return nil, err
		}
	}

	return currentNode, nil
}
