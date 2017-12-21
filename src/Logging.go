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
)

// generic logging function
func Log(msg string, params map[string]interface{}, loglevel string) {
	switch loglevel {
	case INFO:
		log.WithFields(params).Info(msg)
	case DEBUG:
		log.WithFields(params).Debug(msg)
	case ERROR:
		log.WithFields(params).Error(msg)
	}
}

// Error logs out errors using fields
// e.g. Error("error msg", fmt.Errorf("my error"))
// or using log.Fields / map[string]interace type
// Error("error msg", map[string]interface{}{"val1": "foo", "val2": "bar"}
func Error(msg string, params interface{}) {
	switch params.(type) {
	case map[string]interface{}:
		fields := params.(map[string]interface{})
		if fields["err"] != nil {
			Log(msg, fields, ERROR)
		}
	case error:
		if params != nil {
			fields := map[string]interface{}{"msg": params}
			Log(msg, fields, ERROR)
		}
	}
}

// Debug uses the same syntax as Error function but does
// not support error type and does not check for errors
// e.g. Debug("debug", fmt.Errorf("my debug msg"))
// or using log.Fields / map[string]interace type
// Debug("debug msg", map[string]interface{}{"val1": "foo", "val2": "bar"}
func Debug(msg string, params interface{}) {
	switch params.(type) {
	case map[string]interface{}:
		fields := params.(map[string]interface{})
		Log(msg, fields, DEBUG)
	case string, error:
		fields := map[string]interface{}{"msg": params}
		Log(msg, fields, DEBUG)
	}
}

// Info uses the same syntax as Error function but does
// not support error type and does not check for errors
// e.g. Info("info", fmt.Errorf("my info msg"))
// or using log.Fields / map[string]interace type
// Info("info msg", map[string]interface{}{
//	"err": "foo",
//	"someField": "bar"
// }
func Info(msg string, params interface{}) {
	switch params.(type) {
	case map[string]interface{}:
		fields := params.(map[string]interface{})
		Log(msg, fields, INFO)
	case string, error:
		fields := map[string]interface{}{"msg": params}
		Log(msg, fields, INFO)
	}
}

// simple error logging without fields
func ErrorMsg(msg string) {
	log.Error(msg)
}

// simple debug logging without fields
func DebugMsg(msg string) {
	log.Debug(msg)
}

// simple info logging without fields
func InfoMsg(msg string) {
	log.Info(msg)
}
