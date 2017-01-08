package jayson

import (
	"encoding/json"
)

//定义一个类型Data类型为JSONObject的JSONNode
type JSONNode struct {
	DataNode
	data JSONObjectType
}

//装换为JsonNode节点
func CastJSONNode(node *DataNode, err error) (*JSONNode, error) {
	if err != nil {
		return nil, err
	}

	o, err := node.JSONObject()

	if err != nil {
		return nil, err
	}

	jNode := new(JSONNode)
	jNode.data = o
	jNode.exists = true
	return jNode, nil
}

//序列化
func (node *JSONNode) MarshaJSON() ([]byte, error) {
	return json.Marshal(node.data)
}

//返回Json node的data
func (node *JSONNode) Data() map[string]*DataNode {
	return node.data
}

//JSONNode的string格式
func (node *JSONNode) String() string {
	f, err := json.Marshal(node.data)

	if err != nil {
		return err.Error()
	}
	return string(f)
}

//遍历节点
func (node *JSONNode) walkNode(keys ...string) (*DataNode, error) {
	return node.walkNde(keys)
}

//序列化一个JsonNode节点
func (node *JSONNode) MarshalJSONNode() ([]byte, error) {
	return json.Marshal(node.data)
}

//遍历到一个JSONObjectType节点的数据
func (node *JSONNode) WalkJSONObject(keys ...string) (JSONObjectType, error) {
	endNode, err := node.walkNde(keys)

	if err != nil {
		return nil, err
	} else {
		obj, err := endNode.JSONObject()

		if err == nil {
			return obj, nil
		} else {
			return nil, err
		}
	}

}

//遍历到一个String节点的数据
func (node *JSONNode) WalkString(keys ...string) (string, error) {
	endNode, err := node.walkNde(keys)

	if err != nil {
		return EMPTY_STRING, err
	} else {
		return endNode.String()
	}
}

//遍历到一个数据类型为Null的节点的数据
func (node *JSONNode) WalkNull(keys ...string) error {
	endNode, err := node.walkNde(keys)

	if err != nil {
		return err
	}
	return endNode.Null()
}

//遍历到一个数据类型为json.Number的节点的数据
func (node *JSONNode) WalkNumber(keys ...string) (json.Number, error) {
	endNode, err := node.walkNde(keys)

	if err != nil {
		n, err := endNode.Number()

		if err != nil {
			return n, nil
		} else {
			return "", nil
		}
	} else {
		return "", nil
	}

}

//遍历到一个数据类型为float64的数据节点的数据
func (node *JSONNode) WalkFloat64(keys ...string) (float64, error) {
	endNode, err := node.walkNde(keys)

	if err != nil {
		return 0, err
	} else {
		n, err := endNode.Float64()

		if err != nil {
			return 0, err
		} else {
			return n, nil
		}
	}
}

//遍历好一个数据类型为int64的数据节点的数据
func (node *JSONNode) WalkInt64(keys ...string) (int64, error) {

	endNode, err := node.walkNde(keys)

	if err != nil {
		n, err := endNode.Int64()

		if err != nil {
			return 0, err
		} else {
			return n, nil
		}
	} else {
		return 0, err
	}
}

//遍历到一个为泛型数据类型节点的数据
func (node *JSONNode) WalkT(keys ...string) (interface{}, error) {
	endNode, err := node.walkNde(keys)

	if err != nil {
		return nil, err
	} else {
		return endNode.T(), nil
	}
}

//遍历到一个bool类型节点的数据
func (node *JSONNode) WalkBoolean(keys ...string) (bool, error) {
	endNode, err := node.walkNde(keys)

	if err != nil {
		return false, err
	} else {
		return endNode.Boolean()
	}
}

//遍历到一个节点的数据为Array
func (node *JSONNode) WalkDataNodeArray(keys ...string) ([]*DataNode, error) {
	endNode, err := node.walkNde(keys)

	if err != nil {
		return nil, err
	} else {
		return endNode.Array()
	}
}

//遍历好的节点为JsonObjectType的节点数据
func (node *JSONNode) WalkJSONObjectArray(keys ...string) ([]JSONObjectType, error) {
	endNode, err := node.walkNde(keys)

	if err != nil {
		return nil, err
	} else {
		array, err := endNode.Array()

		if err != nil {
			return nil, err
		} else {
			jsonObjectArray := make([]JSONObjectType, len(array))

			for index, arrayItem := range array {
				jsonObject, err := arrayItem.JSONObject()

				if err != nil {
					return nil, err
				} else {
					jsonObjectArray[index] = jsonObject
				}
			}

			return jsonObjectArray, nil
		}
	}
}

func (node *JSONNode) WalkNumberArray(keys ...string) ([]json.Number, error) {
	endNode, err := node.walkNde(keys)

	if err != nil {
		return nil, err
	} else {
		array, err := endNode.Array()

		if err != nil {
			return nil, err
		} else {
			numberArray := make([]json.Number, len(array))

			for index, element := range array {
				number, err := element.Number()

				if err != nil {
					return nil, err
				} else {
					numberArray[index] = number
				}
			}
			return numberArray, nil
		}

	}

}

func (node *JSONNode) WalkNullArray(keys ...string) (int64, error) {
	endNode, err := node.walkNde(keys)

	if err != nil {
		return 0, err
	} else {

		array, err := endNode.Array()

		if err != nil {
			return 0, err
		} else {

			var length int64 = 0

			for _, arrayItem := range array {
				err := arrayItem.Null()

				if err != nil {
					return 0, err
				} else {
					length++
				}

			}
			return length, nil
		}
	}
}

func (node *JSONNode) WalkInt64Array(keys ...string) ([]int64, error) {
	endNode, err := node.walkNde(keys)

	if err != nil {
		return nil, err
	} else {

		array, err := endNode.Array()

		if err != nil {
			return nil, err
		} else {

			typedArray := make([]int64, len(array))

			for index, arrayItem := range array {
				typedArrayItem, err := arrayItem.Int64()

				if err != nil {
					return nil, err
				} else {
					typedArray[index] = typedArrayItem
				}

			}
			return typedArray, nil
		}
	}
}

func (node *JSONNode) WalkBooleanArray(keys ...string) ([]bool, error) {
	endNode, err := node.walkNde(keys)

	if err != nil {
		return nil, err
	} else {

		array, err := endNode.Array()

		if err != nil {
			return nil, err
		} else {

			typedArray := make([]bool, len(array))

			for index, arrayItem := range array {
				typedArrayItem, err := arrayItem.Boolean()

				if err != nil {
					return nil, err
				} else {
					typedArray[index] = typedArrayItem
				}

			}
			return typedArray, nil
		}
	}
}

func (node *JSONNode) GetFloat64Array(keys ...string) ([]float64, error) {
	child, err := node.walkNde(keys)

	if err != nil {
		return nil, err
	} else {

		array, err := child.Array()

		if err != nil {
			return nil, err
		} else {

			typedArray := make([]float64, len(array))

			for index, arrayItem := range array {
				typedArrayItem, err := arrayItem.Float64()

				if err != nil {
					return nil, err
				} else {
					typedArray[index] = typedArrayItem
				}

			}
			return typedArray, nil
		}
	}
}

func (node *JSONNode) GetStringArray(keys ...string) ([]string, error) {
	endNode, err := node.walkNde(keys)

	if err != nil {
		return nil, err
	} else {

		array, err := endNode.Array()

		if err != nil {
			return nil, err
		} else {

			typedArray := make([]string, len(array))

			for index, arrayItem := range array {
				typedArrayItem, err := arrayItem.String()

				if err != nil {
					return nil, err
				} else {
					typedArray[index] = typedArrayItem
				}

			}
			return typedArray, nil
		}
	}
}
