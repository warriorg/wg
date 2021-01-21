package util

import (
	"github.com/warriorg/wg/snowflake"
)

var idNode *snowflake.Node

func init() {
	idNode, _ = snowflake.NewNode(0)
}

// NextID next
func NextID() int64 {
	return idNode.Generate().Int64()
}
