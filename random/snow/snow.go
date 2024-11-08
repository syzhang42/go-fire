package snow

import (
	"net"
	"sync"
	"time"
)

type Snowflake struct {
	mu        sync.Mutex
	epoch     int64
	nodeID    int64
	seq       int64
	lastStamp int64
}

const (
	epoch     = 1622505600000 // 自定义起始时间戳
	nodeBits  = 10            // 节点 ID 位数
	seqBits   = 12            // 序列号位数
	maxNodeID = (1 << nodeBits) - 1
	maxSeq    = (1 << seqBits) - 1
)

var DefaultSf *Snowflake

func NewSnowflake() {
	DefaultSf = &Snowflake{
		epoch:  epoch,
		nodeID: getNodeID(),
	}
}

func (s *Snowflake) NextID() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().UnixNano() / 1e6

	if now == s.lastStamp {
		s.seq = (s.seq + 1) & maxSeq
		if s.seq == 0 {
			for now <= s.lastStamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		s.seq = 0
	}

	s.lastStamp = now
	id := (now-epoch)<<22 | (s.nodeID << seqBits) | s.seq
	return id
}

// 从机器的 IP 地址生成节点 ID
func getNodeID() int64 {
	ips, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	var nodeID int64
	for _, ip := range ips {
		if ipnet, ok := ip.(*net.IPNet); ok && ipnet.IP.IsGlobalUnicast() {
			ipParts := ipnet.IP.To4()
			if ipParts != nil {
				// 将 IP 地址的四个部分组合成一个整数
				nodeID = int64(ipParts[0])<<24 | int64(ipParts[1])<<16 | int64(ipParts[2])<<8 | int64(ipParts[3])
				return nodeID & maxNodeID // 确保在有效范围内
			}
		}
	}
	panic("getNodeID error")
}
