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
	"github.com/gorilla/mux"
	"gopkg.in/throttled/throttled.v2"
	"gopkg.in/throttled/throttled.v2/store/memstore"
	"net/http"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	// latest
	AddRoutes(router)

	//v1
	AddV1Routes(router.PathPrefix("/v1").Subrouter())

	//v2 only dummy yet
	AddV2Routes(router.PathPrefix("/v2").Subrouter())

	return router
}

// add default routes + ratelimit
func AddRoutes(router *mux.Router) {
	store, err := memstore.New(65536)
	Error("ROUTES: could not create memstore", err)

	// rate limiter
	quota := throttled.RateQuota{throttled.PerMin(Conf.RateLimit.Limit), Conf.RateLimit.Burst}
	rateLimiter, err := throttled.NewGCRARateLimiter(store, quota)
	Error("ROUTES: error in ratelimiting", err)

	httpRateLimiter := throttled.HTTPRateLimiter{
		RateLimiter: rateLimiter,
		VaryBy:      &throttled.VaryBy{Path: true},
	}

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(use(handler, BasicAuthHandler, httpRateLimiter.RateLimit))
	}
}

// Middleware chainer
func use(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for _, m := range middleware {
		h = m(h)
	}

	return h
}

// version 1 routes
func AddV1Routes(router *mux.Router) {
	AddRoutes(router)
}

// dummy for version 2 routes
func AddV2Routes(router *mux.Router) {
	AddRoutes(router)
}
