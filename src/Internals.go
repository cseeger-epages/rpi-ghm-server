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
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Parse filter functions
func ParseQueryStrings(r *http.Request) QueryStrings {
	vals := r.URL.Query()

	// set defaults
	qs := QueryStrings{false}

	// Parse
	_, ok := vals["prettify"]
	if ok {
		qs.prettify = true
	}

	return qs
}

// Handles some filters and does what the name says
func EncodeAndSend(w http.ResponseWriter, r *http.Request, qs QueryStrings, msg interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if Conf.Cors.AllowCrossOrigin {
		w.Header().Set("Access-Control-Allow-Origin", Conf.Cors.AllowFrom)
		w.Header().Set("Access-Control-Allow-Methods", strings.Join(Conf.Cors.CorsMethods, ","))
	}

	if Conf.Tls.Hsts {
		hsts := fmt.Sprintf("max-age=%d; includeSubDomains", Conf.Tls.HstsMaxAge)
		w.Header().Add("Strict-Transport-Security", hsts)
	}

	var err error
	// i need to encode the data twice for checking etag
	// and for sending with/without prettyfy maybe there
	// is a better way
	etagdata, err := json.Marshal(msg)
	Error("json marshal error etag", err)
	etagsha := sha256.Sum256([]byte(etagdata))
	etag := fmt.Sprintf("%x", etagsha)
	w.Header().Set("ETag", etag)

	if match := r.Header.Get("If-None-Match"); match != "" {
		if strings.Contains(match, etag) {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}

	w.WriteHeader(http.StatusOK)

	if qs.prettify {
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", " ")
		err = encoder.Encode(msg)
	} else {
		err = json.NewEncoder(w).Encode(msg)
	}
	Error("json parse error", err)
}
