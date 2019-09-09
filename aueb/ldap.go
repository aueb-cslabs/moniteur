package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"gopkg.in/ldap.v3"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var ldapConf LdapConfiguration

type LdapConfiguration struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
	UseTLS  bool   `yaml:"use_tls"`

	BaseDN       string `yaml:"base_dn"`
	BindUsername string `yaml:"bind_username"`
	BindPassword string `yaml:"bind_password"`

	SearchClass string `yaml:"search_class"`

	UsernameAttribute string `yaml:"username_attribute"`

	AuthorizedUsers map[string]string `yaml:"authorized_users"`
}

func loadLdapConfig(ldapConf *LdapConfiguration) error {
	all, err := ioutil.ReadFile("ldap.yml")
	if err != nil {
		return err
	}
	return yaml.Unmarshal(all, ldapConf)
}

func initLdapConnection() (*ldap.Conn, error) {
	var l *ldap.Conn
	var err error

	if ldapConf.Address == "" {
		return nil, errors.New("ldap has not been enabled")
	}

	if ldapConf.UseTLS {
		l, err = ldap.DialTLS("tcp4",
			fmt.Sprintf("%s:%d", ldapConf.Address, ldapConf.Port),
			&tls.Config{ServerName: ldapConf.Address, InsecureSkipVerify: true})
	} else {
		l, err = ldap.Dial("tcp4", fmt.Sprintf("%s:%d", ldapConf.Address, ldapConf.Port))
	}

	if err != nil {
		return nil, err
	}
	return l, l.Bind(ldapConf.BindUsername, ldapConf.BindPassword)
}

func authenticateLdap(username string, password string) (bool, error) {
	l, err := initLdapConnection()
	if err != nil {
		return false, err
	}
	defer l.Close()

	if ldapConf.AuthorizedUsers[username] == "" {
		return false, errors.New("user is not authenticated")
	}

	// Search for the given username
	filter := fmt.Sprintf("(&(objectClass=%s)(%s=%s))",
		ldapConf.SearchClass, ldapConf.UsernameAttribute, username)

	searchRequest := ldap.NewSearchRequest(
		ldapConf.BaseDN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases,
		0, 0, false, filter, []string{"dn"}, nil)

	sr, err := l.Search(searchRequest)
	if err != nil {
		return false, err
	}

	if len(sr.Entries) != 1 {
		return false, errors.New("user does not exist or too many entries returned")
	}

	targetDN := sr.Entries[0].DN
	// Bind as the user to verify their password
	err = l.Bind(targetDN, password)
	if err != nil {
		return false, err
	}

	return true, nil
}
