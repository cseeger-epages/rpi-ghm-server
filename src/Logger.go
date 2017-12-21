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
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

func InitLogger() {
	switch Conf.Logging.Type {
	case LOGFORMATJSON:
		log.SetFormatter(&log.JSONFormatter{})
	case LOGFORMATTEXT:
		formatter := &log.TextFormatter{
			FullTimestamp: true,
		}
		log.SetFormatter(formatter)
	default:
		log.WithFields(log.Fields{
			"logformat": Conf.Logging.Type,
			"default":   LOGFORMATTEXT,
		}).Error("unknown logformat using default")
	}

	switch Conf.Logging.Loglevel {
	case INFO:
		log.SetLevel(log.InfoLevel)
	case ERROR:
		log.SetLevel(log.ErrorLevel)
	case DEBUG:
		log.SetLevel(log.DebugLevel)
	default:
		log.WithFields(log.Fields{
			"loglevel": Conf.Logging.Loglevel,
			"default":  INFO,
		}).Error("unknown loglevel using default")
		log.SetLevel(log.InfoLevel)
	}

	switch Conf.Logging.Output {
	case LOGFILE:
		logfile, err := os.OpenFile(Conf.Logging.Logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.WithFields(log.Fields{
				"filepath": Conf.Logging.Logfile,
			}).Error("can't open logfile use stdout")
			Conf.Logging.Output = LOGSTDOUT
		}
		log.SetOutput(logfile)
		log.WithFields(log.Fields{
			"output": LOGFILE,
			"format": Conf.Logging.Type,
		}).Debug("initialising logging")
	case LOGSTDOUT:
		log.WithFields(log.Fields{
			"output": LOGSTDOUT,
			"format": Conf.Logging.Type,
		}).Debug("using logging method")
	default:
		log.WithFields(log.Fields{
			"output":  Conf.Logging.Output,
			"default": LOGSTDOUT,
		}).Error("unknown log output using default")
		Conf.Logging.Output = LOGSTDOUT
	}
}

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.WithFields(log.Fields{
			"method":      r.Method,
			"request-uri": r.RequestURI,
			"duration":    time.Since(start),
			"name":        name,
			"ip":          r.RemoteAddr,
		}).Info("REQUEST")
	})
}
