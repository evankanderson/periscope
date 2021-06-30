# periscope

Local HTTP proxy to Kubernetes clusters. Lets you poke at cluster-internal HTTP
addresses using using `curl` or other local tools without needing to fuss with
`kubectl port-forward`, just set your `http_proxy` environment variable.

Also supports proxying requests on a port in the cluster back to a local server,
sort of like [telepresence](https://telepresence.io)/etc.

## Installation

```shell
go get github.com/evankanderson/periscope
```

## Sample Usage

The following command launches a pod on your kubernetes cluster called
`periscope-remote-proxy`, and then connects the local proxy to the remote
cluster.

```shell
$ periscope --setup -t localhost:1234
2021/06/30 15:06:23 Setting up pod on remote cluster...
2021/06/30 15:06:23 Connecting to pod on cluster to forward...
2021/06/30 15:06:24 Listening on "localhost:6080", forwarding to "localhost:5000". Incoming will connect to "localhost:1234"
```

If you don't have an existing HTTP service in your cluster, you can run the following image:

```shell
$ kubectl create deployment server --image=gcr.io/knative-samples/helloworld-go
deployment.apps/server created
$ kubectl expose deployment server --port=80 --target-port=8080
service/server exposed
```

And then, in another shell:

```shell
$ export http_proxy=localhost:6080
$ curl http://server.default/
Hello World!
```

### Other stuff to try: local proxying

Start an HTTP server on your machine. A simple python example:

```shell
python3 -m http.server 1234
```

Launch a shell in your cluster:

```shell
$ kubectl run shell --image ubuntu -- sleep 3600
pod/shell created
$ kubectl exec -it shell -- /bin/bash
```

And then within the shell:

```shell
apt-get update && apt-get install -y curl && curl http://periscope-remote-proxy/
```

You should be able to see results from the python server on your desktop!

## WARNING

THIS IS EXPERIMENTAL!

There are still at least a few hacky/TODO parts, mostly around managing the jobs
in the cluster:

- The `port-forward` and `apply` parts of the pod lifecycle aren't managed
  consistently, or really with any knowledge of each other.
- The `periscope-remote-proxy` pod is left running on the cluster after shutdown
  when started with `--setup`. `--setup` should really tear it down, too.

It shouldn't break your cluster, but during development, all of the following
have been observed:

- The `periscope-remote-proxy` pod in the cluster hangs on deletion, making it
  hard to remove.
- A `kubectl port-forward` process gets "lost" locally, and needs to be manually
  killed before the forwarding works again.
- The current GRPC process encapsulates request/reply HTTP connections, but
  won't work for websockets, HTTP/2 (in general), or streamed / chunked
  responses.
- Unusual error behavior on the HTTP forwarding could cause one or the other
  processes to panic (most of these should be fixed, and the rest are
  high-priority bugs)
