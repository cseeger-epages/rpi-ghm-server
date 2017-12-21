package main

import (
	"crypto/tls"
	"fmt"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"strings"
)

func CreateTLSConf() *tls.Config {

	var minversion uint16 = tls.VersionTLS12
	switch Conf.Tls.Minversion {
	case SSL30:
		minversion = tls.VersionSSL30
	case TLS10:
		minversion = tls.VersionTLS10
	case TLS11:
		minversion = tls.VersionTLS11
	case TLS12:
		minversion = tls.VersionTLS12
	default:
		Debug("no tls minversion found using default", map[string]interface{}{"default": "tls12"})
	}

	var curves []tls.CurveID
	for _, v := range Conf.Tls.CurvePrefs {
		switch v {
		case CURVEP256:
			curves = append(curves, tls.CurveP256)
		case CURVEP384:
			curves = append(curves, tls.CurveP384)
		case CURVEP521:
			curves = append(curves, tls.CurveP521)
		case X25519:
			curves = append(curves, tls.X25519)
		}
	}

	if curves == nil {
		Debug("no tls curvepref found using default", map[string]interface{}{"default": "p256, p384, p521"})
		curves = []tls.CurveID{tls.CurveP256, tls.CurveP384, tls.CurveP521}
	}

	var ciphers []uint16
	for _, v := range Conf.Tls.Ciphers {
		ciphers = append(ciphers, CipherMap[v])
	}

	if ciphers == nil {
		Debug("no tls ciphers found using default", map[string]interface{}{
			"default": "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256, TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384, TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256, TLS_RSA_WITH_AES_256_GCM_SHA384,",
		})
		ciphers = []uint16{tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256, tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384, tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256, tls.TLS_RSA_WITH_AES_256_GCM_SHA384}
	}

	tlsCfg := &tls.Config{
		MinVersion:               minversion,
		CurvePreferences:         curves,
		PreferServerCipherSuites: Conf.Tls.PreferServerCiphers,
		CipherSuites:             ciphers,
	}
	return tlsCfg
}

func CreateServerAndListener(router *mux.Router, ip string, port string) (*http.Server, net.Listener, error) {

	if port == "" || router == nil {
		return nil, nil, fmt.Errorf("Router or Port is empty")
	}

	if ip != "" && !IsIPAddr(ip) {
		return nil, nil, fmt.Errorf("unknown IPAddress format")
	}

	// default initilization
	listen := fmt.Sprintf("%s:%s", "", port)

	if IsV4Addr(ip) {
		listen = fmt.Sprintf("%s:%s", ip, port)
	}
	if IsV6Addr(ip) {
		listen = fmt.Sprintf("[%s]:%s", ip, port)
	}
	l, err := net.Listen("tcp", listen)
	if err != nil {
		return nil, nil, err
	}

	tlsCfg := CreateTLSConf()
	srv := &http.Server{
		Addr:         listen,
		Handler:      router,
		TLSConfig:    tlsCfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}

	return srv, l, nil
}

func IsIPAddr(ip string) bool {
	if IsV4Addr(ip) {
		return true
	} else if IsV6Addr(ip) {
		return true
	}
	return false
}

func IsV4Addr(ip string) bool {
	trial := net.ParseIP(ip)
	if trial.To4() != nil {
		return true
	}
	return false
}

func IsV6Addr(ip string) bool {
	trial := net.ParseIP(ip)
	if trial.To16() != nil {
		return trial != nil && strings.Contains(ip, ":")
	}
	return false
}
