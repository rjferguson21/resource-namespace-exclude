# resource-namespace-exclude

### Overview

<!--mdtogo:Short-->

Kustomize KRM Plugin for removing namespaces from ClusterResources


```
apiVersion: v1
kind: ResourceExcludeTransformer
metadata:
  name: cluster-isssuer-exclude
  annotations:
    config.kubernetes.io/function: |
      container: 
        image: rjferguson21/resource-namespace-exclude:latest
clusterResources: 
  - ClusterIssuer
```
```
kustomize build --enable-alpha-plugins example
```


example/Kustomization.yaml
```
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
namespace: foobar
resources:
  - pod.yaml
  - cluster-issuer.yaml
transformers:
  - resource-exclude-transformer.yaml
```

Output
```
apiVersion: cert-manager.io/v1alpha2
kind: ClusterIssuer
metadata:
  name: letsencrypt-prod
spec: null
---
apiVersion: v1
kind: Pod
metadata:
  labels:
    run: nginx
  name: nginx
  namespace: foobar
spec:
  containers:
  - image: nginx
    name: nginx
```

### Why is this neccesary?
https://github.com/kubernetes-sigs/kustomize/issues/880