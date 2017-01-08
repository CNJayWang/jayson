package jayson

import (
	"encoding/json"
	"io"
)

const (
	KEY_NOT_FOUND string = "key noy found"
	EMPTY_STRING  string = ""
)

//从io中读取一个Json数据
func NewDataNodefromReader(reader io.Reader) (*DataNode, error) {
	node := new(DataNode)
	d := json.NewDecoder(reader)
	d.UseNumber()
	error := d.Decode(&node.data)
	return node, error
}

//从字节数组中获取一个JsonNode节点
func NewJSONNodeFromBytes(b []byte) (*JSONNode, error) {
	return CastJSONNode(UnMarshaBytes(b))
}

//从io流中获取一个JsonNode节点
func NewJSONNodeFromReader(reader io.Reader) (*JSONNode, error) {

	return CastJSONNode(NewDataNodefromReader(reader))
}
