/*
   GOLANG REST API Skeleton

   Copyright (C) 2017 Carsten Seeger

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.

   @author Carsten Seeger
   @copyright Copyright (C) 2017 Carsten Seeger
   @license http://www.gnu.org/licenses/gpl-3.0 GNU General Public License 3
   @link https://github.com/cseeger-epages/rest-api-go-skeleton
*/

package main

import (
	"github.com/BurntSushi/toml"
	"os"
)

type config struct {
	General   general   `toml:"general"`
	Certs     certs     `toml:"certs"`
	Tls       tlsconf   `toml:"tls"`
	Cors      cors      `toml:"cors"`
	Logging   logging   `toml:"logging"`
	RateLimit rateLimit `toml:"ratelimit"`
	DB        database  `toml:"database"`
	Users     []user    `toml:"user"`
}

type general struct {
	Listen    string
	Port      string
	BasicAuth bool
}

type certs struct {
	Public  string
	Private string
}

type tlsconf struct {
	Minversion          string
	CurvePrefs          []string
	Ciphers             []string
	PreferServerCiphers bool
	Hsts                bool
	HstsMaxAge          int
}

type cors struct {
	AllowCrossOrigin bool
	CorsMethods      []string
	AllowFrom        string
}

type logging struct {
	Type     string
	Loglevel string
	Output   string
	Logfile  string
}

type rateLimit struct {
	Limit int
	Burst int
}

type database struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

type user struct {
	Username string
	Password string
}

type jobConfig struct {
	Jobs []job `toml:"job"`
}

type job struct {
	Path string `toml:"path"`
	Cmd  string `toml:"cmd"`
}

func ParseConfig(fileName string, conf interface{}) error {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		Error("config error", err)
		os.Exit(1)
	}
	_, err := toml.DecodeFile(fileName, conf)
	return err
}
