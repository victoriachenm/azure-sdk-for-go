## Go

``` yaml
clear-output-folder: false
export-clients: true
go: true
input-file: https://github.com/Azure/azure-rest-api-specs/blob/ac155b972d0619a6e5bf665a863fb05ee7eeb30f/specification/keyvault/data-plane/Microsoft.KeyVault/stable/7.3/keys.json
license-header: MICROSOFT_MIT_NO_VERSION
module: github.com/Azure/azure-sdk-for-go/sdk/keyvault/azkeys
module-version: 0.6.0
openapi-type: "data-plane"
output-folder: ../azkeys
override-client-name: Client
security: "AADToken"
security-scopes:  "https://vault.azure.net/.default"
use: "@autorest/go@4.0.0-preview.42"
version: "^3.0.0"

directive:
  # make vault URL a parameter of the client constructor
  - from: swagger-document
    where: $["x-ms-parameterized-host"]
    transform: $.parameters[0]["x-ms-parameter-location"] = "client"

  # rename parameter models to match their methods
  - rename-model:
      from: KeyCreateParameters
      to: CreateKeyParameters
  - rename-model:
      from: KeyExportParameters
      to: ExportKeyParameters
  - rename-model:
      from: KeyImportParameters
      to: ImportKeyParameters
  - rename-model:
      from: KeyReleaseParameters
      to: ReleaseParameters
  - rename-model:
      from: KeyRestoreParameters
      to: RestoreKeyParameters
  - rename-model:
      from: KeySignParameters
      to: SignParameters
  - rename-model:
      from: KeyUpdateParameters
      to: UpdateKeyParameters
  - rename-model:
      from: KeyVerifyParameters
      to: VerifyParameters

  # rename paged operations from Get* to List*
  - rename-operation:
      from: GetDeletedKeys
      to: ListDeletedKeys
  - rename-operation:
      from: GetKeys
      to: ListKeys
  - rename-operation:
      from: GetKeyVersions
      to: ListKeyVersions

  # delete unused error models
  - from: models.go
    where: $
    transform: return $.replace(/(?:\/\/.*\s)+type (?:Error|KeyVaultError).+\{(?:\s.+\s)+\}\s/g, "");

  # delete the Attributes model defined in common.json (it's used only with allOf)
  - from: models.go
    where: $
    transform: return $.replace(/(?:\/\/.*\s)+type Attributes.+\{(?:\s.+\s)+\}\s/, "");
  - from: models_serde.go
    where: $
    transform: return $.replace(/(?:\/\/.*\s)+func \(a \*?Attributes\).*\{\s(?:.+\s)+\}\s/g, "");

  # delete generated constructor
  - from: client.go
    where: $
    transform: return $.replace(/(?:\/\/.*\s)+func NewClient.+\{\s(?:.+\s)+\}\s/, "");

  # delete the version path param check (version == "" is legal for Key Vault but indescribable by OpenAPI)
  - from: client.go
    where: $
    transform: return $.replace(/\sif keyVersion == "" \{\s+.+keyVersion cannot be empty"\)\s+\}\s/g, "");

  # delete client name prefix from method options and response types
  - from:
      - client.go
      - models.go
      - response_types.go
    where: $
    transform: return $.replace(/Client(\w+)((?:Options|Response))/g, "$1$2");

  # rename Kid ("Key ID") fields "KID" and make their type handwritten so we can add parsing methods
  - from: models.go
    where: $
    transform: return $.replace(/(\s)Kid \*string(\s+.*)/g, "$1KID *ID$2")
  - from: models_serde.go
    where: $
    transform: return $.replace(/Kid/g, "KID")

  # Maxresults -> MaxResults
  - from:
      - client.go
      - models.go
    where: $
    transform: return $.replace(/Maxresults/g, "MaxResults")

  # keyName, keyVersion -> name, version
  - from: client.go
  - where: $
  - transform: return $.replace(/keyName/g, "name").replace(/keyVersion/g, "version")
```
