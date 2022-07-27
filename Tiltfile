# -*- mode: Python -*-
# allow_k8s_contexts('rancher-desktop')

local_resource('testkube-watch-controller-build',
    'GOOS=linux GOARCH=amd64 goreleaser build --single-target --snapshot --rm-dist',
    dir='.', deps=['./main.go', './go.mod', './cmd/', './config/', './pkg/', './utils/'])

# to disable push with rancher desktop we need to use custom_build instead of docker_build
custom_build('lreimer/testkube-watch-controller', 'docker build -t $EXPECTED_REF .', ['./'], disable_push=True)
k8s_yaml(kustomize('./k8s/'))