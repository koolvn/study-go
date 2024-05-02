package config

import "os"

type Config struct {
	Host         string
	Port         string
	CertFile     string
	KeyFile      string
	LdapAddr     string
	LdapBindUser string
	LdapBindPass string
	LdapBaseDN   string
	Https        bool
}

func NewConfig() *Config {
	host := getEnv("APP_HOST", "0.0.0.0")
	port := getEnv("APP_PORT", "8080")
	certFile := getEnv("APP_CERT_FILE", "")
	keyFile := getEnv("APP_KEY_FILE", "")
	ldapAddr := getEnv("APP_LDAP_ADDRESS", "")
	ldapBindUser := getEnv("APP_LDAP_BIND_USER", "")
	ldapBindPass := getEnv("APP_LDAP_BIND_PASS", "")
	ldapBaseDN := getEnv("APP_LDAP_BASE_DN", "")
	https := false
	if certFile != "" && keyFile != "" {
		https = true
	}
	return &Config{
		Host:         host,
		Port:         port,
		CertFile:     certFile,
		KeyFile:      keyFile,
		LdapAddr:     ldapAddr,
		LdapBindUser: ldapBindUser,
		LdapBindPass: ldapBindPass,
		LdapBaseDN:   ldapBaseDN,
		Https:        https,
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	} else {
		return fallback
	}
}
