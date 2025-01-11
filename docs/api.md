# API Endpoints
## Users
**Register**
* Endpoint: `POST /signup`
* Request Body:
    ```json
    {
        "name": "ian",
        "username": "ianabc",
        "password": "password",
        "gender": "male",
        "age": 27,
        "location": "jakarta"
    }
    ```
* Response Body:
    ```json
    {
        "message": "user registered successfully",
        "data": {
            "id": 23
        }
    }
    ```
**Login**
* Endpoint: `POST /signin`
* Request Body:
    ```json
    {
        "username": "ianabc",
        "password": "password"
    }
    ```
* Response Body:
    ```json
    {
        "message": "logged in successfully",
        "data": {
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzd2lwZXIiLCJleHAiOjE3MzY2NjE4NTQsImlhdCI6MTczNjU3NTQ1NCwianRpIjoiNDUyZjY2MWUtMzE3Zi00NzQ3LWE1ZWUtMzY0NzM0MTU2OTE2IiwidXNlcl9pZCI6MjN9.-dNknKaHLwlCIJfYepxHXQB9arMlR8uv8IU-muSQ6Qo",
            "refresh_token": ""
        }
    }
    ```
## Swipe-swipe
All the API below need token from login API in Authorization Bearer Header, such as
```
Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzd2lwZXIiLCJleHAiOjE3MzY2NjE4NTQsImlhdCI6MTczNjU3NTQ1NCwianRpIjoiNDUyZjY2MWUtMzE3Zi00NzQ3LWE1ZWUtMzY0NzM0MTU2OTE2IiwidXNlcl9pZCI6MjN9.-dNknKaHLwlCIJfYepxHXQB9arMlR8uv8IU-muSQ6Qo
```

**Get Specific Profile**
* Endpoint: `GET /profile/{id}`
* Response Body:
    ```json
    {
        "message": "profile fetched",
        "data": {
            "id": 23,
            "name": "ian",
            "username": "ianabc",
            "age": 27,
            "string": "jakarta",
            "is_premium": true,
            "is_verified": true,
            "created_at": "2025-01-11T06:04:08.656039Z",
            "updated_at": "2025-01-11T06:05:02.928049Z"
        }
    }
    ```

**Get Profile From Queue**
* Endpoint: `GET /queue`
* Response Body:
    ```json
    {
        "message": "profile fetched",
        "data": {
            "id": 18,
            "name": "Steve",
            "username": "Steve_18",
            "age": 35,
            "gender": "male",
            "string": "Dallas",
            "is_premium": true,
            "is_verified": true,
            "created_at": "2025-01-11T05:57:30.918386Z"
        }
    }
    ```
**Swipe Right**
* Endpoint: `POST /swipe/right/{id}`
* Response Body:
    ```json
    {
        "message": "swiped right",
        "data": {
            "is_matched": true
        }
    }
    ```
**Swipe Left**
* Endpoint: `POST /swipe/left/{id}`
* Response Body:
    ```json
    {
	    "message": "swiped left"
    }
    ```

## Premium
All the API below need token from login API in Authorization Bearer Header, such as
```
Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzd2lwZXIiLCJleHAiOjE3MzY2NjE4NTQsImlhdCI6MTczNjU3NTQ1NCwianRpIjoiNDUyZjY2MWUtMzE3Zi00NzQ3LWE1ZWUtMzY0NzM0MTU2OTE2IiwidXNlcl9pZCI6MjN9.-dNknKaHLwlCIJfYepxHXQB9arMlR8uv8IU-muSQ6Qo
```
**Enable Premium**
* Endpoint: `POST /premium`
* Request Body:
    ```json
    {
        "id":23
    }
    ```
* Response Body:
    ```json
    {
        "message": "premium user enabled successfully",
        "data": {
            "id": 23,
            "created_at": "0001-01-01T00:00:00Z"
        }
    }
    ```