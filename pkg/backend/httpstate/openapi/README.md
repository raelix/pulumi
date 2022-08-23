# Go API client for openapi

The Pulumi Service API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.pulumi.com/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**GetCLIVersionInfo**](docs/DefaultApi.md#getcliversioninfo) | **Get** /cli/version | getCLIVersionInfo asks the service for information about versions of the CLI (the newest version as well as the oldest version before the CLI should warn about an upgrade).
*StackApi* | [**CreateStack**](docs/StackApi.md#createstack) | **Post** /{organization}/{project} | CreateStack creates a stack with the given cloud and stack name in the scope of the indicated project.
*StackApi* | [**DecryptValue**](docs/StackApi.md#decryptvalue) | **Post** /{organization}/{project}/{stack}/decrypt | DecryptValue decrypts a ciphertext value in the context of the indicated stack.
*StackApi* | [**DeleteStack**](docs/StackApi.md#deletestack) | **Delete** /{organization}/{project}/{stack} | DeleteStack deletes the indicated stack. If force is true, the stack is deleted even if it contains resources.
*StackApi* | [**DoesProjectExist**](docs/StackApi.md#doesprojectexist) | **Head** /stacks/{organization}/{project} | Returns true if a project with the given name exists, or false otherwise.
*StackApi* | [**EncryptValue**](docs/StackApi.md#encryptvalue) | **Post** /{organization}/{project}/{stack}/encrypt | EncryptValue encrypts a plaintext value in the context of the indicated stack.
*StackApi* | [**GetStack**](docs/StackApi.md#getstack) | **Get** /{organization}/{project}/{stack} | GetStack retrieves the stack with the given name.
*StackApi* | [**UpdateStackTags**](docs/StackApi.md#updatestacktags) | **Patch** /{organization}/{project}/{stack}/tags | UpdateStackTags updates the stacks&#39;s tags, replacing all existing tags.
*UserApi* | [**GetCurrentUser**](docs/UserApi.md#getcurrentuser) | **Get** /user | Returns the current user.
*UserApi* | [**ListStacks**](docs/UserApi.md#liststacks) | **Get** /user/stacks | Lists all stacks the current user has access to, optionally filtered by project.


## Documentation For Models

 - [CreateStackRequest](docs/CreateStackRequest.md)
 - [DecryptValueRequest](docs/DecryptValueRequest.md)
 - [DecryptValueResponse](docs/DecryptValueResponse.md)
 - [EncryptValueRequest](docs/EncryptValueRequest.md)
 - [EncryptValueResponse](docs/EncryptValueResponse.md)
 - [GetCLIVersionInfo200Response](docs/GetCLIVersionInfo200Response.md)
 - [ListStacks200Response](docs/ListStacks200Response.md)
 - [OperationStatus](docs/OperationStatus.md)
 - [ServiceUser](docs/ServiceUser.md)
 - [ServiceUserInfo](docs/ServiceUserInfo.md)
 - [Stack](docs/Stack.md)
 - [StackSummary](docs/StackSummary.md)
 - [UpdateKind](docs/UpdateKind.md)


## Documentation For Authorization



### token

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


