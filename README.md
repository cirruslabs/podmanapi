# podmanapi

The `podmanapi` package was generated using the Swagger Codegen 3.0.25 [downloaded from the Maven Repository](https://mvnrepository.com/artifact/io.swagger.codegen.v3/swagger-codegen-cli):

```
java -jar swagger-codegen-cli-3.0.25.jar generate -l go -i https://storage.googleapis.com/libpod-master-releases/swagger-latest.yaml -o pkg/swagger
```

The link to the [`swagger-latest.yaml`](https://storage.googleapis.com/libpod-master-releases/swagger-latest.yaml) was found in the auto-generated [API documentation page](https://podman.readthedocs.io/en/latest/_static/api.html).

Afterwards:

* the missing `os` imports were added
* `ContainerRenameLibpod`'s and `ContainerRename`'s `name` arguments were renamed to `nameInPath` and `nameInQuery`
  * with their usage renames according to the context (i.e. whether the argument was used to craft a path or a query)
* `model_plugin_config_linux.go` was renamed to `model_plugin_config_linux_.go`

... to fix the compilation.
