package brute

import (
	"errors"
	"fmt"
	"github.com/yhy0/hydra/hydra"
	"strings"
)

/**
  @author: yhy
  @since: 2022/7/1
  @desc: //TODO
**/

func Hydra(host string, port int, service string) (string, error) {
	var (
		msg  string
		err  error
		flag bool
	)

	service = strings.ToLower(service)

	switch service {
	case "redis":
		msg, err = hydra.Redis(host, port)
	case "ftp":
		msg, err = hydra.FTP(host, port)
	case "memcached":
		msg, err = hydra.Memcached(host, port)
	case "ldap", "rsh-spx", "ssh":
		msg, err = hydra.SSH(host, port)
	case "mysql":
		msg, err = hydra.Mysql(host, port)
	case "oracle", "oracle-tns":
		msg, err = hydra.Oracle(host, port)
	case "mongod", "mongodb":
		msg, err = hydra.Mongodb(host, port)
	case "postgresql":
		msg, err = hydra.Postgresql(host, port)
	case "ms-sql-s":
		msg, err = hydra.SQLServer(host, port)
	case "ms-wbt-server", "ssl/ms-wbt-server":
		msg, err = hydra.RDP(host, port)
	case "microsoft-ds":
		msg, err = hydra.SMB(host, port)
	default:
		flag = true
	}

	if flag {
		return "", errors.New(fmt.Sprintf("Service mismatch: %s %d", service, port))
	}

	return msg, err
}
