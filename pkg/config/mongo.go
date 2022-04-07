package config

import (
	"fmt"
	"net/url"
	"strings"
)

type MongoDB struct {
	User     *string     `toml:"username" yaml:"username"`
	Password *string     `toml:"password" yaml:"password"`
	Database string      `toml:"database" yaml:"database"`
	Nodes    []MongoNode `toml:"nodes" yaml:"nodes"`

	Compressors    []string `toml:"compressors" yaml:"compressors"`
	ReplicaSet     *string  `toml:"replica_set" yaml:"replica_set"`
	MaxPoolSize    *uint64  `toml:"max_pool_size" yaml:"max_pool_size"`
	MinPoolSize    *uint64  `toml:"min_pool_size" yaml:"min_pool_size"`
	ReadPreference *string  `toml:"read_preference" yaml:"read_preference"`
}

// MongoNode represents information about replica in MongoDB replicaSet.
type MongoNode struct {
	Host string `toml:"host" yaml:"host"`
	Port int    `toml:"port" yaml:"port"`
}

func (m *MongoDB) URI() string {
	nodes := make([]string, len(m.Nodes))
	for i, node := range m.Nodes {
		nodes[i] = fmt.Sprintf("%s:%d", node.Host, node.Port)
	}

	addr := "mongodb://"
	values := make(url.Values)

	if m.User != nil && m.Password != nil {
		addr += *m.User + ":" + *m.Password + "@"
		values.Add("authSource", m.Database)
	}

	addr += strings.Join(nodes, ",") + "/" + m.Database

	if len(m.Compressors) > 0 {
		values.Add("compressors", strings.Join(m.Compressors, ","))
	}

	return addr + "?" + values.Encode()
}
