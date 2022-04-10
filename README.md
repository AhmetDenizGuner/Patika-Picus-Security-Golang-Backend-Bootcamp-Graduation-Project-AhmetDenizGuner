# Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner

## Pre-Requirments
- MySQL: databaseName: picus_shop
- Redis: AddrURI: "localhost:6379"
```console
foo@bar:$ docker pull redis
foo@bar:$ docker run --name redis-test-instance -p 6379:6379 -d redis
foo
```


### ENVIRONMENT VARIABLES ###
- "ENV" = "local" --> to this auto, you can uncomment line25 in [router.go](internal/api/router.go)
- You can create own configuration and enviroment varible based on app.local.yaml[here](config/app.local.yaml)

## Notes
- You can get the sample request using by swagger [docs.](cmd/app/docs)
- For Add Bulk Category feature, there is example folder in [here.](resources/uploaded) First column name, second column is category id.

# Gin Picus-Shop API
This service provides basic e-commerce API.

## Version: 1.0

**Contact information:**  
Ahmet Deniz Guner  
ahmetdenizguner@gmail.com

**License:** [Apache 2.0](http://www.apache.org/licenses/LICENSE-2.0.html)

### /auth/login

#### POST
##### Summary

This endpoint used for login

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| signinRequest | body | signin information | Yes | [user.SigninRequest](internal/api/types/types.go) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 |  |  |
| 400 | Bad Request | [user.ApiErrorResponse](shared/types.go) |
| 404 | Not Found | [user.ApiErrorResponse](shared/types.go) |
| 507 | Insufficient Storage | [user.ApiErrorResponse](shared/types.go) |

### /auth/logout

#### POST
##### Summary

This endpoint used for logout

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| signoutRequest | body | signout information | Yes | [user.SignoutRequest](internal/api/types/types.go) |
| Authorization | header | Authorization | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 |  |  |
| 400 | Bad Request | [user.ApiErrorResponse](shared/types.go) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| ApiKeyAuth | |

### /auth/signup

#### POST
##### Summary

This endpoint used for register

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| signupRequest | body | signup information | Yes | [user.SignupRequest](internal/api/types/types.go) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 |  |  |
| 400 | Bad Request | [user.ApiErrorResponse](shared/types.go) |

### /cart/add-item

#### POST
##### Summary

This endpoint used for adding new element to user cart

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| stock_code | formData | stock code of adding element | Yes | string |
| Authorization | header | Authorization | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 201 |  |

##### Security

| Security Schema | Scopes |
| --- | --- |
| ApiKeyAuth | |

### /cart/list

#### GET
##### Summary

This endpoint used for see current cart

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| Authorization | header | Authorization | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 |  |

##### Security

| Security Schema | Scopes |
| --- | --- |
| ApiKeyAuth | |

### /cart/update-delete-item

#### PUT
##### Summary

This endpoint used for deleting or update the item that is in cart already

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| stock_code | formData | stock code of update element | Yes | string |
| update_quantity | formData | new cart quantity of element | Yes | integer |
| Authorization | header | Authorization | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 204 |  |

##### Security

| Security Schema | Scopes |
| --- | --- |
| ApiKeyAuth | |

### /category/add-all

#### POST
##### Summary

This endpoint used for uploading csv and creating categories from this csv file

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| file | formData | form data CSV | Yes | file |
| Authorization | header | Authorization | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 |  |  |
| 415 | Unsupported Media Type | [category.ApiErrorResponse](shared/types.go) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| ApiKeyAuth | |

### /category/list

#### GET
##### Summary

This endpoint used for getting category list with pagination

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| page | query | Page Index | No | integer |
| pageSize | query | Page Size | No | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 |  |  |
| 507 | Insufficient Storage | [category.ApiErrorResponse](shared/types.go) |

### /order/cancel

#### POST
##### Summary

This endpoint used for creating order with products in basket

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| order_delete_id | formData | id belongs order will be canceled | Yes | integer |
| Authorization | header | Authorization | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 201 |  |

##### Security

| Security Schema | Scopes |
| --- | --- |
| ApiKeyAuth | |

### /order/complete

#### POST
##### Summary

This endpoint used for creating order with products in basket

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| Authorization | header | Authorization | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 201 |  |

##### Security

| Security Schema | Scopes |
| --- | --- |
| ApiKeyAuth | |

### /order/list

#### GET
##### Summary

This endpoint used for see the active orders

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| Authorization | header | Authorization | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 201 |  |

##### Security

| Security Schema | Scopes |
| --- | --- |
| ApiKeyAuth | |

### /product/add

#### POST
##### Summary

This endpoint used for creating new product

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| requestModel | body | it is a new product model | Yes | [product.AddProductRequest](internal/api/types/types.go) |
| Authorization | header | Authorization | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 201 |  |

##### Security

| Security Schema | Scopes |
| --- | --- |
| ApiKeyAuth | |

### /product/delete

#### DELETE
##### Summary

This endpoint used for remove the product fromDB

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| stock_code | formData | stock code belongs product will be deleted | Yes | string |
| Authorization | header | Authorization | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 204 |  |

##### Security

| Security Schema | Scopes |
| --- | --- |
| ApiKeyAuth | |

### /product/list

#### GET
##### Summary

This endpoint used for getting product list with pagination

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| page | query | Page Index | No | integer |
| pageSize | query | Page Size | No | integer |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 |  |

### /product/search{searchKeyword}

#### POST
##### Summary

This endpoint used for searching product with pagination

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| page | query | Page Index | No | integer |
| pageSize | query | Page Size | No | integer |
| searchKeyword | query | word will be searched | No | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 |  |

### /product/update

#### PUT
##### Summary

This endpoint used for updates product in DB

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| requestModel | body | it is an updated product model | Yes | [product.AddProductRequest](internal/api/types/types.go) |
| Authorization | header | Authorization | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 204 |  |

##### Security

| Security Schema | Scopes |
| --- | --- |
| ApiKeyAuth | |

### Models

#### category.ApiErrorResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| error_message | string |  | No |
| is_success | boolean |  | No |

#### product.AddProductRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| category_id | integer |  | No |
| description | string |  | No |
| name | string |  | No |
| price | number |  | No |
| stock_code | string |  | No |
| stock_quantity | integer |  | No |

#### user.ApiErrorResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| error_message | string |  | No |
| is_success | boolean |  | No |

#### user.SigninRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| email | string |  | No |
| password | string |  | No |

#### user.SignoutRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| email | string |  | No |

#### user.SignupRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| email | string |  | No |
| name | string |  | No |
| password | string |  | No |
