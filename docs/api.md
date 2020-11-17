# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [accounts.proto](#accounts.proto)
    - [Account](#accounts.Account)
    - [CreateRequest](#accounts.CreateRequest)
    - [CreateResponse](#accounts.CreateResponse)
    - [DeleteRequest](#accounts.DeleteRequest)
    - [DeleteResponse](#accounts.DeleteResponse)
    - [GetRequest](#accounts.GetRequest)
    - [GetResponse](#accounts.GetResponse)
    - [ListRequest](#accounts.ListRequest)
    - [ListResponse](#accounts.ListResponse)
    - [UpdateRequest](#accounts.UpdateRequest)
    - [UpdateResponse](#accounts.UpdateResponse)
  
    - [currencyType](#accounts.currencyType)
  
    - [Accounts](#accounts.Accounts)
  
- [payments.proto](#payments.proto)
    - [CreateRequest](#payments.CreateRequest)
    - [CreateResponse](#payments.CreateResponse)
    - [GetRequest](#payments.GetRequest)
    - [GetResponse](#payments.GetResponse)
    - [ListRequest](#payments.ListRequest)
    - [ListResponse](#payments.ListResponse)
    - [Payment](#payments.Payment)
  
    - [directionType](#payments.directionType)
  
    - [Payments](#payments.Payments)
  
- [Scalar Value Types](#scalar-value-types)



<a name="accounts.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## accounts.proto



<a name="accounts.Account"></a>

### Account



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |
| name | [string](#string) |  |  |
| balance | [double](#double) |  |  |
| currency | [currencyType](#accounts.currencyType) |  |  |
| isAvailable | [bool](#bool) |  |  |






<a name="accounts.CreateRequest"></a>

### CreateRequest
create


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| balance | [double](#double) |  |  |
| currency | [currencyType](#accounts.currencyType) |  |  |






<a name="accounts.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |






<a name="accounts.DeleteRequest"></a>

### DeleteRequest
delete


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |






<a name="accounts.DeleteResponse"></a>

### DeleteResponse







<a name="accounts.GetRequest"></a>

### GetRequest
get


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |






<a name="accounts.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [Account](#accounts.Account) |  |  |






<a name="accounts.ListRequest"></a>

### ListRequest
list


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| isAvailable | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  |  |






<a name="accounts.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| accounts | [Account](#accounts.Account) | repeated |  |






<a name="accounts.UpdateRequest"></a>

### UpdateRequest
update


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |
| name | [string](#string) |  |  |
| balance | [double](#double) |  |  |
| currency | [currencyType](#accounts.currencyType) |  |  |
| isAvailable | [bool](#bool) |  |  |






<a name="accounts.UpdateResponse"></a>

### UpdateResponse






 


<a name="accounts.currencyType"></a>

### currencyType


| Name | Number | Description |
| ---- | ------ | ----------- |
| USD | 0 |  |
| EUR | 1 |  |
| RU | 2 |  |


 

 


<a name="accounts.Accounts"></a>

### Accounts


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| List | [ListRequest](#accounts.ListRequest) | [ListResponse](#accounts.ListResponse) |  |
| Get | [GetRequest](#accounts.GetRequest) | [GetResponse](#accounts.GetResponse) |  |
| Create | [CreateRequest](#accounts.CreateRequest) | [CreateResponse](#accounts.CreateResponse) |  |
| Update | [UpdateRequest](#accounts.UpdateRequest) | [UpdateResponse](#accounts.UpdateResponse) |  |
| Delete | [DeleteRequest](#accounts.DeleteRequest) | [DeleteResponse](#accounts.DeleteResponse) |  |

 



<a name="payments.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## payments.proto



<a name="payments.CreateRequest"></a>

### CreateRequest
create


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| amount | [double](#double) |  |  |
| accountFrom | [int64](#int64) |  |  |
| accountTo | [int64](#int64) |  |  |






<a name="payments.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |






<a name="payments.GetRequest"></a>

### GetRequest
get


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |






<a name="payments.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment | [Payment](#payments.Payment) |  |  |






<a name="payments.ListRequest"></a>

### ListRequest
list






<a name="payments.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payments | [Payment](#payments.Payment) | repeated |  |






<a name="payments.Payment"></a>

### Payment



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int64](#int64) |  |  |
| amount | [double](#double) |  |  |
| accountFrom | [int64](#int64) |  |  |
| accountTo | [int64](#int64) |  |  |
| direction | [directionType](#payments.directionType) |  |  |





 


<a name="payments.directionType"></a>

### directionType


| Name | Number | Description |
| ---- | ------ | ----------- |
| incoming | 0 |  |
| outgoing | 1 |  |


 

 


<a name="payments.Payments"></a>

### Payments


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| List | [ListRequest](#payments.ListRequest) | [ListResponse](#payments.ListResponse) |  |
| Get | [GetRequest](#payments.GetRequest) | [GetResponse](#payments.GetResponse) |  |
| Create | [CreateRequest](#payments.CreateRequest) | [CreateResponse](#payments.CreateResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

