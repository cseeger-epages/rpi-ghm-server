# Simple job server to connect google home (mini) to your raspberry pi
The server is a simple command / path mapper. A so called job is a pair of 
two strings one containing the path and the other the command that will be executed.


## Configuration
The configuration files can be found under conf/api.conf and conf/job.conf.
They use the TOML format language.
See the config files for more details

### Example
conf/job.conf for controling tv power state via hdmi cec protocoll
```
[[job]]
path = "/hdmi-on"
cmd = "echo 'on 0' | cec-client -s -d 1"

[[job]]
path = "/hdmi-off"
cmd = "echo 'standby 0' | cec-client -s -d 1"
```

this will create the following path structure

```
https://<serverip>:<port>/hdmi-on
https://<serverip>:<port>/hdmi-off
```

curls

```
# switch tv on 
curl -k -H "Authorization: Basic $(echo '<username>:<password>' | base64)" https://<serverip>:<port>/hdmi-on
# switch tv off
curl -k -H "Authorization: Basic $(echo '<username>:<password>' | base64)" https://<serverip>:<port>/hdmi-off
```

## security 

These URL's can be used by [IFTTT](https://ifttt.com/google_assistant). Be carefull here you need to expose your API to the outside world so use some authentication and ip whitelisting and maybe some [Port knocking](https://wiki.archlinux.org/index.php/Port_knocking)

## build and run
for standard linux distributions 
```
make
```
for rpi
```
make rpi
```

### cli arguments
```
-c <configfile>
-j <jobconfigfile>
```
certificate path is configured under the conf/api.conf (default is certs/server{crt,key})

### generate certificates
```
cd certs
# Key considerations for algorithm "RSA" ≥ 2048-bit
openssl genrsa -out server.key 2048

# Key considerations for algorithm "ECDSA" ≥ secp384r1
# List ECDSA the supported curves (openssl ecparam -list_curves)
openssl ecparam -genkey -name secp384r1 -out server.key
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```

### configs
dont forget to copy configs

## some curls
```
curl -k -X GET -H "Authorization: Basic dGVzdHVzZXI6dGVzdHBhc3MK" https://localhost:8443/
curl -k -X GET -H "Authorization: Basic dGVzdHVzZXI6dGVzdHBhc3MK" https://localhost:8443/help
```

## basic auth test stuff
```
testuser:testpass - dGVzdHVzZXI6dGVzdHBhc3MK
username:password - dXNlcm5hbWU6cGFzc3dvcmQK
```
