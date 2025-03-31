package config

// reids 配置
type RedisConfig struct {
	RedisConnString string `mapstructure:"redis_conn_string" yaml:"redis_conn_string"`
	RedisPassword   string `mapstructure:"redis_password" yaml:"redis_password"`
	Database        int    `mapstructure:"database" yaml:"database"` // 新增 Redis 数据库选择
	RedisMasterName string `mapstructure:"redis_master_name" yaml:"redis_master_name"`
	SyncFrequency   string `mapstructure:"sync_frequency" yaml:"sync_frequency"`
}

// 数据库配置
type DatabaseConfig struct {
	Options           string `mapstructure:"options" yaml:"options"` // 数据库连接参数
	IsMasterNode      bool   `mapstructure:"is_master_node" yaml:"is_master_node"`
	SQLitePath        string `mapstructure:"sqlite_path" yaml:"sqlite_path"`
	SQLiteBusyTimeout int    `mapstructure:"sqlite_busy_timeout" yaml:"sqlite_busy_timeout"`
	UsingSQLite       bool   `mapstructure:"using_sqlite" yaml:"using_sqlite"`
	UsingPostgreSQL   bool   `mapstructure:"using_postgresql" yaml:"using_postgresql"`
	UsingMySQL        bool   `mapstructure:"using_mysql" yaml:"using_mysql"`
}
