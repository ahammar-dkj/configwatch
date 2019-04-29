
## Configwatch

Build Docker image, deploy 3 replicas of the service and tail logs:

```
$ make build-docker deploy
$ stern configwatch
```

From another terminal, edit the ConfigMap:

```
$ kubectl edit configmap configwatch
```

Once the Kubelet has refreshed the mounted ConfigMap volume, then the service should
emit an event that the file has changed. You should see something like this in the
logs:

```
configwatch-597d7857d4-8jj9x configwatch 2019/04/29 16:17:57 event:  "/etc/configwatch/..2019_04_29_16_17_57.560591983": CREATE
configwatch-597d7857d4-8jj9x configwatch 2019/04/29 16:17:57 event:  "/etc/configwatch/..2019_04_29_16_17_57.560591983": CHMOD
configwatch-597d7857d4-8jj9x configwatch 2019/04/29 16:17:57 event:  "/etc/configwatch/..data_tmp": RENAME
configwatch-597d7857d4-8jj9x configwatch 2019/04/29 16:17:57 event:  "/etc/configwatch/..data": CREATE
configwatch-597d7857d4-8jj9x configwatch 2019/04/29 16:17:57 event:  "/etc/configwatch/..2019_04_29_16_16_46.143564166": REMOVE
configwatch-597d7857d4-nc9x8 configwatch 2019/04/29 16:18:12 event:  "/etc/configwatch/..2019_04_29_16_18_12.084645634": CREATE
configwatch-597d7857d4-nc9x8 configwatch 2019/04/29 16:18:12 event:  "/etc/configwatch/..2019_04_29_16_18_12.084645634": CHMOD
configwatch-597d7857d4-nc9x8 configwatch 2019/04/29 16:18:12 event:  "/etc/configwatch/..data_tmp": RENAME
configwatch-597d7857d4-nc9x8 configwatch 2019/04/29 16:18:12 event:  "/etc/configwatch/..data": CREATE
configwatch-597d7857d4-nc9x8 configwatch 2019/04/29 16:18:12 event:  "/etc/configwatch/..2019_04_29_16_16_51.788360686": REMOVE
configwatch-597d7857d4-pbwmv configwatch 2019/04/29 16:18:14 event:  "/etc/configwatch/..2019_04_29_16_18_14.680660089": CREATE
configwatch-597d7857d4-pbwmv configwatch 2019/04/29 16:18:14 event:  "/etc/configwatch/..2019_04_29_16_18_14.680660089": CHMOD
configwatch-597d7857d4-pbwmv configwatch 2019/04/29 16:18:14 event:  "/etc/configwatch/..data_tmp": RENAME
configwatch-597d7857d4-pbwmv configwatch 2019/04/29 16:18:14 event:  "/etc/configwatch/..data": CREATE
configwatch-597d7857d4-pbwmv configwatch 2019/04/29 16:18:14 event:  "/etc/configwatch/..2019_04_29_16_16_48.648115377": REMOVE
```

https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/#mounted-configmaps-are-updated-automatically


### References

* fsnotify: https://github.com/fsnotify/fsnotify
* stern: https://github.com/wercker/stern
