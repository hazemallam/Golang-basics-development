# Table Of Contents
1. [Beneficiary List](#beneficiary-list)
   1. [Add Beneficiary](#add-beneficiary)
        1. [Request](#add-beneficiary-request)
        2. [Success Response](#add-beneficiary-success-response)
        3. [Error Responses](#add-beneficiary-error-responses)
        4. [Notes](#add-beneficiary-notes)
   2. [Update Beneficiary](#update-beneficiary)
       1. [Request](#update-beneficiary-request)
       2. [Success Response](#update-beneficiary-sucess-response)
       3. [Error Responses](#update-beneficiary-error-responses)
       4. [Notes](#update-beneficiary-notes)
        

# Beneficiary List
<div id="beneficiary-list"></div>

## Add Beneficiary
<div id="add-beneficiary"></div>

**A user can add beneficiary list so can easily choose one of them to transfer money.**

**URL** : `http://localhost.com/product-backend-mule4`

**Method** : `POST`

**Auth required** : YES

### Request
<div id="add-beneficiary-request"></div>


```json
{
  "signature": "123123123",
  "message": {
    "header": {
      "apiCode": "150011",
      "timestamp": 15857102626979,
      "sessionId": "",
      "walletShortCode": "2020",
      "channelCode": "test_1",
      "walletId": "2421",
      "userId": "74894"
    },
    "payload": {
      "senderOwnerType": "1",
      "senderIdentifyValue": "Swifty1",
      "beneficiaryName": "testtt",
      "destinationPaymentMethod": {
        "paymentMethodType": "849",
        "currency_unit": "XOF",
        "dynamicGroup": {
          "groupId": 4,
          "groupName": "Destination GTP",
          "inputParameters": [
            {
              "parameterCode": 101,
              "parameterValue": "123450"
            },
            {
              "parameterCode": 100,
              "parameterValue": "9876"
            }
          ]
        }
      },
      "authParams": [
        {
          "authCode": "0",
          "authValue": "9632"
        }
      ]
    },
    "additionalData": {
    },
    "echoData": {
      "Sample": "Sample"
    }
  }
}
```

### Success Response
<div id="add-beneficiary-success-response"></div>

**Code** : `200 OK`

**Content example** : For the example above, when the new beneficiary is added successfully.

```json
{
  "signature": null,
  "response": {
    "timestamp": "1644319673169",
    "apiCode": "150011",
    "channelCode": "test_1",
    "traceId": null,
    "sessionId": null,
    "transactionReferenceId": null,
    "result": {},
    "echoData": {
      "sessionIdMacKey": null,
      "sessionIdWalletShortCode": null,
      "sessionIdSenderIdentifyValue": null
    },
    "responseStatus": "Succeeded",
    "responseError": null,
    "walletCurrency": "CDF",
    "successMessage": "Beneficiary has been added successfully!"
  }
}
```

### Error Responses
<div id="add-beneficiary-error-responses"></div>

**Condition** : Account already exists

**Response**

```json
{
  "response": {
    "timestamp": "1642368117263",
    "responseStatus": "Failed",
    "responseError": {
      "errorCode": "VAL202097",
      "errorMessage": "Beneficiary already exists"
    }
  }
}
```

**Or**

**Condition** : Special characters entered within beneficiary name.

**Response**

```json
{
  "response": {
    "timestamp": "1642368117263",
    "responseStatus": "Failed",
    "responseError": {
      "errorCode": "VAL202096",
      "errorMessage": "Special characters not allowed in Beneficiary name"
    }
  }
}
```
**Or**

**Condition** : No customer exists with provided alias.

**Response**

```json
{
  "response": {
    "timestamp": "1642368117263",
    "responseStatus": "Failed",
    "responseError": {
      "errorCode": "VAL202102",
      "errorMessage": "No Customer exists with provided Alias"
    }
  }
}
```
**Or**

**Condition** : No customer exists with provided alias.

**Response**

```json
{
  "response": {
    "timestamp": "1642368117263",
    "responseStatus": "Failed",
    "responseError": {
      "errorCode": "VAL202101",
      "errorMessage": "No Customer exists with provided mobile number"
    }
  }
}
```

### Notes
<div id="add-beneficiary-notes"></div>

**Mandatory Fields**

| Request Fields           | Notes |
|--------------------------|-------|
| Header’s parameters      | M     |
| payload                  | M     |
| additionalData           | M     |
| signature                | M     |
| message                  | M     |
| senderOwnerType          | M     |
| senderIdentifyValue      | M     |
| beneficiaryName          | M     |
| destinationPaymentMethod | M     |
| groupId                  | M     |
| inputParameters          | M     |

**Request Description**

```json
{
  "singnature": "<string>",
  "message": {
    "header": {
      "apiCode": "<string>",
      "timestamp": "<long>",
      "sessionId": "<string>",
      "walletId": "<string>",
      "walletShortCode": "<string>",
      "channelCode": "<string>",
      "userId": "<string>"
    },
    "payload": {
      "senderOwnerType": "<string>",
      "senderIdentfyValue": "<string>",
      "beneficiaryName": "<string>",
      "destinationPaymentMethod": {
        "paymentMethodType": "<string>",
        "dynamicGroup": {
          "groupId": "<integer>",
          "groupName": "<string>",
          "inputParameters": [
            {
              "parameterCode": "<integer>",
              "parameterValue": "<string>"
            }
          ]
        },
        "currency_unit": "<string>"
      }
    },
    "additionalData": "<object>",
    "echoData": {
      "Sample": "<string>"
    }
  }
}
```