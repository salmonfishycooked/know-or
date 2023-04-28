package snowflake

import (
	sf "github.com/bwmarrin/snowflake"
	"time"
)

var node *sf.Node

// Init 用来初始化分布式ID生成器 node
func Init(startTime string, machineID int64) error {
	st, err := time.Parse("2006-01-02", startTime)
	if err != nil {
		return err
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return err
}

// GenID 用来生成一个 int64 的 ID
func GenID() int64 {
	return node.Generate().Int64()
}
