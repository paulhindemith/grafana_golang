# test
Only run following script in docker container: `paulhindemith/dev-infra/third_party/grafana/6.6.0/Dockerfile`.
```
docker run -it -w/go/src/github.com/paulhindemith/grafana_golang  -v ~/go/src/github.com/paulhindemith/grafana_golang:/go/src/github.com/paulhindemith/grafana_golang paulhindemith/grafana-dev:6.6.0 ./test/presubmit-tests.sh
```
