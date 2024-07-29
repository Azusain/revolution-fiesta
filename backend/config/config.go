package config

import (
	"crypto/rand"
	"crypto/rsa"
	"strings"
	"time"
)

const (
	Version         string        = "0.0.1"
	Port            string        = "8080"
	DatabaseUsr     string        = "postgres"
	DatabasePasswd  string        = "270153"
	DatabaseHost    string        = "localhost"
	DatabasePort    string        = "5432"
	DatabaseName    string        = "mydb"
	SchemaFilePath  string        = "../schema.sql"
	RedisAddr       string        = "127.0.0.1:6379"
	RedisPassword   string        = ""
	RedisDB                       = 0
	TokenExpiration time.Duration = time.Hour / 4
)

var (
	DatabaseUrl string
	PublicKey   *rsa.PublicKey
	PrivateKey  *rsa.PrivateKey
)

func createKey() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	publicKey := &privateKey.PublicKey

	return privateKey, publicKey, nil
}
func init() {
	DatabaseUrl = getPgConnUrl(DatabaseUsr, DatabasePasswd, DatabaseHost, DatabasePort, DatabaseName)
	PrivateKey, PublicKey, _ = createKey()
}

func getPgConnUrl(usr, passwd, host, port, database string) string {
	builder := strings.Builder{}

	_, _ = builder.WriteString("postgres://")
	_, _ = builder.WriteString(usr)

	if passwd != "" {
		_, _ = builder.WriteRune(':')
		_, _ = builder.WriteString(passwd)
	}

	_, _ = builder.WriteRune('@')
	_, _ = builder.WriteString(host)

	if port != "" {
		_, _ = builder.WriteRune(':')
		_, _ = builder.WriteString(port)
	}

	_, _ = builder.WriteRune('/')
	_, _ = builder.WriteString(database)
	_, _ = builder.WriteString("?sslmode=disable")

	return builder.String()
}
