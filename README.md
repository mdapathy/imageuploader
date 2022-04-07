## ImageUploader

Image uploader allows users to upload, delete and view pictures they had previously uploaded.

To run the project:

1. run `docker-compose up -d` to create a mongoDB instance  (if the latter is not already present).
2. run `make` to build the project
3. run `./bin/server`

----

### API description


#### POST `api/v1/private/gallery/images` 
Route to upload images

**Headers**:

*X-User-Id*: required, implies that the system has a separate microservice for authorization to which the proxy redirects the requests to set this header 

**Request Body**: 

        {  
            "content" : "image_in_base64_format"
        }

**Request Body Parameters:**

*content*: required field, is a base64 representation of the image, supported formats are png and jpeg

**Response:**


201 Created - if the upload was successful

---
500 Internal Server Error - if an internal error occurs

**Response Body:**

    {
        "code": 500,
        "messages": [
        "system.unhealthy"
        ]   
    }

----
400 Bad Request - if the request cannot be completed due to user's error in input

**Response Body:**  

    `{
        "code": 400,
        "messages": [
        "content.invalid_value"
        ] 
    }`

**Response Body Fields:**

*code*: code of the error

*messages*: string enum, values are "content.required", "content.invalid_value"

----
#### GET `api/v1/private/gallery/images`
Route to view a list of images

**Parameters:**

*limit*: max amount of requests to return, default value 5

*offset*: offset for the records to return, used for pagination, default value 0.

*size_from*: indicates min size of the image

*size_to*: indicates max size (including this value) of the image

*created_from*: indicates min time of the image creation, in RFC339 date format

*created_to*: indicates max time (including this value) of the image creation, in RFC339 date format

**Headers**:

*X-User-Id*: required, implies that the system has a separate microservice for authorization to which the proxy redirects the requests to set the header

**Responses:**

200 OK - if the request was successful

**Response Body**

    {
        "total": 1,
        "data": [
            "id": "624f032eb6cff800172bc33e",
            "user_id": "default",
            "content": "base64_image",
            "size": 52980,
            "created_at": "2022-04-06T13:57:47.182Z"
        ]
    }

**Response Body Fields**
   
total: total amount of records available for this user to access

data: array of image objects

*Object in data array*:

id: id of the image, is an objectID

user_id: id of the user who uploaded the image

content: image in base64 format

size: size of the uploaded image

created_at: time when the image was created, in RFC339 date format


----
500 Internal Server Error - if an internal error occurs

**Response Body:**

    {
        "code": 500,
        "messages": [
        "system.unhealthy"
        ]   
    }

----
400 Bad Request - if the request cannot be completed due to user's error in input

**Response Body:**

    `{
        "code": 400,
        "messages": [
        "created_from.invalid"
        ] 
    }`

**Response Body Fields:**

*code*: code of the error

*messages*: string enum, values are "size_to.invalid", "size_from.invalid", "created_from.invalid",  "created_to.invalid"


---

#### GET `api/v1/private/gallery/images/{id}` 
Route to view image with a specified id

**Headers**:

*X-User-Id*: required, implies that the system has a separate microservice for authorization to which the proxy redirects the requests to set the header

**Path Parameters**

id : id of the image, is an objectID

**Responses:**

200 OK - if the request was successful

**Response Body**

    {
        "id": "624f032eb6cff800172bc33e",
        "user_id": "default",
        "content": "base64_image",
        "size": 52980,
        "created_at": "2022-04-06T13:57:47.182Z"
    }

**Response Body Fields**

id: id of the image, is an objectID

user_id: id of the user who uploaded the image

content: image in base64 format

size: size of the uploaded image

created_at: time when the image was created, in RFC339 date format


----
500 Internal Server Error - if an internal error occurs

**Response Body:**

    {
        "code": 500,
        "messages": [
        "system.unhealthy"
        ]   
    }

----
400 Bad Request - if the request cannot be completed due to user's error in input

**Response Body:**

    `{
        "code": 400,
        "messages": [
        "image.not_found"
        ] 
    }`

**Response Body Fields:**

*code*: code of the error

*messages*: string enum, values are "image.not_found", "id.invalid_value"


#### DELETE `api/v1/private/gallery/images/{id}` 
Route to delete image with a specified id 


**Headers**:

*X-User-Id*: required, implies that the system has a separate microservice for authorization to which the proxy redirects the requests to set the header 

**Path Parameters**

id : id of the image, is an objectID


**Responses:**

200 OK - if the request was successful

----
500 Internal Server Error - if an internal error occurs

**Response Body:**

    {
        "code": 500,
        "messages": [
        "system.unhealthy"
        ]   
    }

----
400 Bad Request - if the request cannot be completed due to user's error in input

**Response Body:**

    `{
        "code": 400,
        "messages": [
        "image.not_found"
        ] 
    }`

**Response Body Fields:**

*code*: code of the error

*messages*: string enum, values are "image.not_found", "id.invalid_value"
