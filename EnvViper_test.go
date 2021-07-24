package envviper

import (
	"github.com/stretchr/testify/suite"
	"os"
	"strings"
	"testing"
)

const (
	DbUriPgsql  = "pgsql://test:test@localhost:5432/test?sslmode=disabled"
	DbUriMongo  = "mongo://test:test@localhost:27017/test?authMechanism=scram-sha-1"
	DbUriNats   = "nats://localhost:4222"
	DbUriPgsql2 = "pgsql://test:test@localhost:5433/test"
)

type Configuration struct {
	Http HttpConfig
	Db   map[string]DBConf
}

type HttpConfig struct {
	Port    string
	Timeout int
}

type DBConf struct {
	Uri         string
	Maxlifetime int
}

type viperTestSuite struct {
	suite.Suite
}

func TestViper(t *testing.T) {
	suite.Run(t, new(viperTestSuite))
}

func (s *viperTestSuite) SetupSuite() {
	_ = os.Setenv("VIPER_DB_PGSQL_URI", DbUriPgsql)
	_ = os.Setenv("VIPER_DB_NATS_URI", DbUriNats)
	_ = os.Setenv("VIPER_DB_MONGO_URI", DbUriMongo)
	_ = os.Setenv("VIPER_DB_PGSQL2_URI", DbUriPgsql2)
}

func (s *viperTestSuite) TestUnmarshal() {
	vp := NewEnvViper()
	vp.SetConfigType("json")
	vp.SetConfigFile("config.json")
	err := vp.ReadInConfig()
	s.Nil(err)
	vp.SetEnvParamsSimple("viper")
	var cfg Configuration
	err = vp.Unmarshal(&cfg)
	s.Nil(err)
	s.Equal("8080", cfg.Http.Port)
	s.Equal(30, cfg.Http.Timeout)
	s.Contains(cfg.Db, "pgsql")
	s.Contains(cfg.Db, "nats")
	s.Contains(cfg.Db, "mongo")
	s.Contains(cfg.Db, "pgsql2")
	s.Contains(cfg.Db, "mysql")
	s.Equal(60, cfg.Db["pgsql"].Maxlifetime)
	s.Equal(0, cfg.Db["pgsql2"].Maxlifetime)
	s.Equal(5, cfg.Db["nats"].Maxlifetime)
	s.Equal(5, cfg.Db["mongo"].Maxlifetime)
	s.Equal(10, cfg.Db["mysql"].Maxlifetime)
	s.Equal(DbUriPgsql, cfg.Db["pgsql"].Uri)
	s.Equal(DbUriPgsql2, cfg.Db["pgsql2"].Uri)
	s.Equal(DbUriNats, cfg.Db["nats"].Uri)
	s.Equal(DbUriMongo, cfg.Db["mongo"].Uri)
	s.Equal("", cfg.Db["mysql"].Uri)
}

func (s *viperTestSuite) TestPanics() {
	vp := NewEnvViper()
	s.Panics(func() {
		vp.AutomaticEnv()
	})
	s.Panics(func() {
		vp.SetEnvPrefix("viper")
	})
	s.Panics(func() {
		vp.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	})
}

func (s *viperTestSuite) TestUnmarshalGlobal() {
	vp := DefEnvViper()
	vp.SetConfigType("json")
	vp.SetConfigFile("config.json")
	err := vp.ReadInConfig()
	s.Nil(err)
	vp.SetEnvParamsSimple("viper")
	var cfg Configuration
	err = vp.Unmarshal(&cfg)
	s.Nil(err)
	s.Equal("8080", cfg.Http.Port)
	s.Equal(30, cfg.Http.Timeout)
	s.Contains(cfg.Db, "pgsql")
	s.Contains(cfg.Db, "nats")
	s.Contains(cfg.Db, "mongo")
	s.Contains(cfg.Db, "pgsql2")
	s.Contains(cfg.Db, "mysql")
	s.Equal(60, cfg.Db["pgsql"].Maxlifetime)
	s.Equal(0, cfg.Db["pgsql2"].Maxlifetime)
	s.Equal(5, cfg.Db["nats"].Maxlifetime)
	s.Equal(5, cfg.Db["mongo"].Maxlifetime)
	s.Equal(10, cfg.Db["mysql"].Maxlifetime)
	s.Equal(DbUriPgsql, cfg.Db["pgsql"].Uri)
	s.Equal(DbUriPgsql2, cfg.Db["pgsql2"].Uri)
	s.Equal(DbUriNats, cfg.Db["nats"].Uri)
	s.Equal(DbUriMongo, cfg.Db["mongo"].Uri)
	s.Equal("", cfg.Db["mysql"].Uri)
}
