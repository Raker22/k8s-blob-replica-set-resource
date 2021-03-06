#! /bin/bash
ROOT_PACKAGE="github.com/raker22/k8s-blob-replica-set-resource"
CUSTOM_RESOURCE_NAME="blobreplicaset"
CUSTOM_RESOURCE_VERSION="v1alpha1"

go get -u k8s.io/code-generator/...
cd "${GOPATH%%:*}/src/k8s.io/code-generator"

"${GOPATH%%:*}/src/k8s.io/code-generator/generate-groups.sh" \
  all \
  "${ROOT_PACKAGE}/pkg/client" \
  "${ROOT_PACKAGE}/pkg/apis" \
  "${CUSTOM_RESOURCE_NAME}:${CUSTOM_RESOURCE_VERSION}"

