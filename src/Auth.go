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
	"net/http"
	"strings"
)

func BasicAuthHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			h.ServeHTTP(w, r)
			return
		}

		if !Conf.General.BasicAuth {
			h.ServeHTTP(w, r)
			return
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

		username, password, authOK := r.BasicAuth()
		if authOK == false {
			http.Error(w, "Not authorized", 401)
			xff := r.Header.Get("X-FORWARDED-FOR")
			if xff == "" {
				xff = "not set"
			}
			Debug("Authorization error", map[string]interface{}{
				"RemoteAddr":     r.RemoteAddr,
				"X-FORWARDD-FOR": xff,
			})
			return
		}

		valid := false

		for _, v := range Conf.Users {
			if username == v.Username && strings.TrimSuffix(password, "\n") == v.Password {
				valid = true
			}
		}
		if !valid {
			http.Error(w, "Not authorized", 401)
			return
		}

		h.ServeHTTP(w, r)
	})
}
