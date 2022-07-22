# Testkube Watch Controller

A K8s controller that watches for Kubernetes change events to trigger the execution of Testkube tests and suites. 

The idea, logic and configuration is inspired by the original
[Kubewatch](https://github.com/robusta-dev/kubewatch) project. The big difference is that
on change events the controller will call the Testkube API server to trigger test and test suite execution based on label and annotation metadata.

## Usage

```bash
kubectl testkube create test --file examples/nginx-test.json --name nginx-curl-test --type "curl/test"

kubectl apply -f examples/nginx-service.yaml
kubectl apply -f examples/nginx-deployment.yaml
```

## Maintainer

M.-Leander Reimer (@lreimer), <mario-leander.reimer@qaware.de>

## License

This software is provided under the Apache v2.0 open source license, read the `LICENSE`
file for details.
