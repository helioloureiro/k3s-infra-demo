# k3s-infra-demo
Small helm chart to run tests on k3s.


## Steps

You can replace `namespace` by some other suitable name.

192.168.122.50 is the k3s ingress interface that I can access.

  ```bash
  export namespace=helio
  create namespace $namespace
  helm install demo ./demo -n $namespace --debug
  curl -s http://192.168.122.50/api

  ```

## removal

 ```bash
 helm delete demo -n $namespace
 ```
