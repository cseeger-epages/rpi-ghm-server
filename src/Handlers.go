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
	"fmt"
	"net/http"
	"os/exec"
)

/*
Default Handler Template
func Handler(w http.ResponseWriter, r*http.Request) {
	// caching stuff is handler specific
	w.Header().Set("Cache-Control", "no-store")

	qs := ParseQueryStrings(r)
	msg := HelpMsg{Message: "im a default Handler"}
	EncodeAndSend(w, r, qs, msg)
}
*/

// root handler giving basic API information
func Index(w http.ResponseWriter, r *http.Request) {
	// dont know what should happen here
	w.Header().Set("Cache-Control", "no-store")

	qs := ParseQueryStrings(r)
	message := fmt.Sprintf("Welcome to GOLANG REST API SKELETON please take a look at https://%s/help", r.Host)
	msg := HelpMsg{Message: message}
	EncodeAndSend(w, r, qs, msg)
}

// help reference for all routes
func Help(w http.ResponseWriter, r *http.Request) {
	// never cache help commands
	w.Header().Set("Cache-Control", "no-store")

	qs := ParseQueryStrings(r)

	var msg []PathList

	for _, m := range routes {
		msg = append(msg, PathList{m.Pattern, m.Description})
	}

	EncodeAndSend(w, r, qs, msg)
}

func JobHandler(w http.ResponseWriter, r *http.Request) {
	// never cache help commands
	w.Header().Set("Cache-Control", "no-store")

	qs := ParseQueryStrings(r)

	var msg interface{}

	for _, job := range JobConf.Jobs {
		if job.Path == r.URL.String() {
			Debug("executing", fmt.Sprintf("%s", job.Cmd))
			out, err := exec.Command("bash", "-c", job.Cmd).Output()
			if err != nil {
				Error("cmd exec error", err)
				msg = ErrorMessage{err.Error()}
			} else {
				msg = HelpMsg{string(out)}
			}
		}
	}

	EncodeAndSend(w, r, qs, msg)
}
