# podmanapi

The `podmanapi` package was generated using the Swagger Codegen 3.0.23 [downloaded from the Maven Repository](https://mvnrepository.com/artifact/io.swagger.codegen.v3/swagger-codegen-cli):

```
java -jar swagger-codegen-cli-3.0.23.jar generate -l go -i https://storage.googleapis.com/libpod-master-releases/swagger-latest-master.yaml -o pkg/swagger
```

The link to the [`swagger-latest-master.yaml`](https://storage.googleapis.com/libpod-master-releases/swagger-latest-master.yaml) was found in the auto-generated [API documentation page](https://podman.readthedocs.io/en/latest/_static/api.html).

Afterwards:

* the missing `os` imports were added
* `model_plugin_config_linux.go` was renamed to `model_plugin_config_linux_.go`
* `model_namespace.go`'s `Namespace`'s `String_` field JSON name was changed from `string` to `value` to accomodate the [upstream change](https://github.com/containers/podman/commit/8eb36320ca321ba514b5388cfdb11595e61a7d49)

... to fix the compilation.
