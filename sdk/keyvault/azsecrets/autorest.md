## Go

These settings apply only when `--go` is specified on the command line.

<!-- Original autorest command used by Chris Scott -->
<!-- autorest --use=@autorest/go@4.0.0-preview.27 https://github.com/Azure/azure-rest-api-specs/tree/1e2c9f3ec93078da8078389941531359e274f32a/specification/keyvault/data-plane --tag=package-7.2 --output-folder=internal --module-azsecrets --module-version=0.1.0 --openapi-type="data-plane" --security="AADToken" --security-scopes="https://vault.azure.net/.default" -->

``` yaml
go: true
version: "^3.0.0"
input-file: https://github.com/Azure/azure-rest-api-specs/blob/1e2c9f3ec93078da8078389941531359e274f32a/specification/keyvault/data-plane/Microsoft.KeyVault/stable/7.2/secrets.json
license-header: MICROSOFT_MIT_NO_VERSION
clear-output-folder: false
output-folder: internal
tag: package-7.2
credential-scope: none
use: "@autorest/go@4.0.0-preview.27"
module-version: 0.1.0
```