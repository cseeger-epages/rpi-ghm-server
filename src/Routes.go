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
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Description string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes Routes

func init() {
	routes = Routes{
		Route{
			"Index",
			"GET",
			"/",
			"default page",
			Index,
		},
		Route{
			"Help",
			"GET",
			"/help",
			"help page",
			Help,
		},
	}

	for _, job := range JobConf.Jobs {
		routes = append(routes, Route{
			"job",
			"GET",
			job.Path,
			job.Cmd,
			JobHandler,
		})
	}
}
