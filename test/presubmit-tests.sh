#!/bin/bash

# Copyright 2020 Paulhindemith
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# /bin/bash

readonly STRICT_TEST
readonly ROOT_DIR=$(git rev-parse --show-toplevel)
source ${ROOT_DIR}/vendor/knative.dev/test-infra/scripts/library.sh

set -o errexit

echo ">> $(dirname $0)/../vendor/github.com/paulhindemith/dev-infra/hack/boilerplate/ensure-boilerplate.sh"
$(dirname $0)/../vendor/github.com/paulhindemith/dev-infra/hack/boilerplate/ensure-boilerplate.sh Paulhindemith \
  || ( (( STRICT_TEST )) && abort "ensure-boilerplate.sh is aborted" )

echo ">> go fmt $(dirname $0)/../..."
go fmt $(dirname $0)/../...

service grafana-server start
trap 'service grafana-server stop' EXIT

go test $(dirname $0)/../... -v

echo "success"
