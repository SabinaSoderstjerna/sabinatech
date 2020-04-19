# sabina.tech

Website hosted on [Google Cloud Run](https://cloud.google.com/run)
with a golang server based on [net/http](https://golang.org/pkg/net/http/) package
with the help of the [text/template](https://golang.org/pkg/text/template/) package

## Setup

The folder structure is the following

```markdown
sabinatech
+-- cmd
    +-- main.go
+-- internal    // .go server files
+- src          // .json source files
+- static       // all static resources e.g images and stylesheets
+- templates    // .html template files
```

## Deployment

Deployment is done from a docker image built,
pushed and deployed on [Google Cloud Build](https://cloud.google.com/cloud-build).
The configuration can be found in `cloudbuild.deploy.yaml`.
