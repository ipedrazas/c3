# c3

This is what happens when you don't like d2

ConfigMap / Docker compose controller, `c3` watches for new ConfigMaps with the annotation `compose.docker.io/filename`

When the controller reads a CM, it schedules a Kubernetes job that will do a `d2` transformation like the one below (as an example).

### TODO

[ ] create a container that runs the helm transformation
[ ] add the OCI thing to push data out of the job

### d2 thingsy things

d2 handled this with the dockerfiles instead of compose files

```
# Creates the chart
FROM harbor.alacasa.uk/library/defaults:latest as builder
RUN helm create dapi --starter common-configmap
```

The real interface is the `values.yaml` not the templates. This is where the real transformation
happens

```
FROM harbor.alacasa.uk/library/hgv:v0.1.11 as values-builder

WORKDIR /workspace

COPY ./.meta.yaml /workspace/.meta.yaml
COPY ./gp/targets/helm/target.yaml /targets/target.yaml
COPY ./data /data

USER appuser

RUN python /home/appuser/app/helm.py /workspace/.meta.yaml /targets/target.yaml /workspace/values.yaml

# Spit out the artifacts
FROM scratch as result

COPY --from=builder /home/appuser/dapi ./charts/dapi
COPY --from=values-builder /workspace/values.yaml ./charts/dapi%
```

To generate the artifacts you can do:

```
docker buildx build -f ./target-helm.Dockerfile -o type=local,dest=. . --progress=plain --no-cache
```

Nothing of this has anything to do with `c3`, but I cannot find the markdown file where I have to put it.
