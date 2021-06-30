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
periscope --setup -t localhost:1234
```

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
