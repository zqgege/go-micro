package config

type MongoConfig interface {
	GetEnabled() bool
	GetConn() string
	GETDb() string
	//GETPoolLimit() int
}

type defaultMongoConfig struct {
	Enabled  bool          `json:"enabled"`
	Conn     string        `json:"conn"`
	Db       string        `json:"database"`
	//PoolLimit  int         `json:"poolLimit"`
}

// GetEnabled mongo 配置是否激活
func (r defaultMongoConfig) GetEnabled() bool {
	return r.Enabled
}

// GetConn mongo 地址
func (r defaultMongoConfig) GetConn() string {
	return r.Conn
}

func (r defaultMongoConfig) GETDb() string {
	return r.Db
}

/*func (r defaultMongoConfig)GETPoolLimit() int  {
	return r.PoolLimit
}*/