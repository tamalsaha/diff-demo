# diff-demo

```console
$ go install -v

$ diff-demo diff -s ./examples/voyager/src.yaml -d ./examples/voyager/on-master.yaml 
spec:
  template:
    spec:
      nodeSelector:
        node-role.kubernetes.io/master: ""
      tolerations:
      - key: CriticalAddonsOnly
        operator: Exists
      - effect: NoSchedule
        key: node-role.kubernetes.io/master
        operator: Exists
```
