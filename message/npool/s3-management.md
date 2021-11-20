# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/s3-management.proto](#npool/s3-management.proto)
    - [GetKycImgRequest](#s3.management.v1.GetKycImgRequest)
    - [GetKycImgResponse](#s3.management.v1.GetKycImgResponse)
    - [UploadKycImgRequest](#s3.management.v1.UploadKycImgRequest)
    - [UploadKycImgResponse](#s3.management.v1.UploadKycImgResponse)
    - [VersionResponse](#s3.management.v1.VersionResponse)
  
    - [S3Management](#s3.management.v1.S3Management)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/s3-management.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/s3-management.proto



<a name="s3.management.v1.GetKycImgRequest"></a>

### GetKycImgRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ImgID | [string](#string) |  |  |






<a name="s3.management.v1.GetKycImgResponse"></a>

### GetKycImgResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  | return ImgBase64 |






<a name="s3.management.v1.UploadKycImgRequest"></a>

### UploadKycImgRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserID | [string](#string) |  |  |
| ImgType | [string](#string) |  |  |
| ImgBase64 | [string](#string) |  |  |






<a name="s3.management.v1.UploadKycImgResponse"></a>

### UploadKycImgResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  | return ImgID |






<a name="s3.management.v1.VersionResponse"></a>

### VersionResponse
request body and response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |





 

 

 


<a name="s3.management.v1.S3Management"></a>

### S3Management
Service Name

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [VersionResponse](#s3.management.v1.VersionResponse) | Method Version |
| UploadKycImg | [UploadKycImgRequest](#s3.management.v1.UploadKycImgRequest) | [UploadKycImgResponse](#s3.management.v1.UploadKycImgResponse) |  |
| GetKycImg | [GetKycImgRequest](#s3.management.v1.GetKycImgRequest) | [GetKycImgResponse](#s3.management.v1.GetKycImgResponse) |  |

 



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

