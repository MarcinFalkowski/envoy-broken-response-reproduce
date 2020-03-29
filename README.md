docker-compose environment where we can reproduce envoy returning broken response when TLS http/1.1 cluster 
returns "upgrade: h2" and "transfer-encoding: chunked"

# Run

```
docker-compose up --build --abort-on-container-exit
```

# Test

Request to endpoint returning `transfer-encoding: chunked` and `upgrade: h2`

* through envoy: `curl -H "host: service" -v http://localhost:6002/upgrade-h2-chunked`
* directly: `curl -k --http1.1 -v https://localhost:8443/upgrade-h2-chunked`


Through envoy we are receiving invalid http response with error `transfer closed with outstanding read data remaining`.

Directly everything works OK

---
Request to endpoint returning `transfer-encoding: chunked` only works correctly through envoy:
```
curl -H "host: service" -v http://localhost:6002/chunked
```

---
Request to endpoint returning `upgrade: h2` only works correctly through envoy:
```
curl -H "host: service" -v http://localhost:6002/upgrade-h2
```
