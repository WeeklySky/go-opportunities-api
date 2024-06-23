
# Golang REST API for Openings

Welcome to the Golang REST API for managing job openings. This API allows you to create, read, update, and delete opening job opening records, as well as retrieve a list of all job openings.

This project was developed following the instructions from [@arthur404dev](https://github.com/arthur404dev). If you liked this project, give him a star [here](https://github.com/arthur404dev/gopportunities), he deserve it. If you want to learn how to develop this project, you can learn it [here](https://www.youtube.com/live/L6gk7FHBNkM?si=UpUak-8ghaAwcm5f).

## Table of Contents

-   [Endpoints](#endpoints)
    -   [GET /opening](#get-opening)
    -   [POST /opening](#post-opening)
    -   [PUT /opening](#put-opening)
    -   [DELETE /opening](#delete-opening)
    -   [GET /openings](#get-openings)
-   [Installation](#installation)
-   [Usage](#usage)
-   [Contributing](#contributing)
-   [License](#license)

## Endpoints

### GET /opening

Retrieve an existing opening by its ID.

**Request:**


```http
GET /opening?id={id}
``` 

**Response:**

```json
{
	"data":  {
		"ID":  2,
		"CreatedAt":  "2024-06-23T16:53:37.332476-03:00",
		"UpdatedAt":  "2024-06-23T18:18:24.178782-03:00",
		"DeletedAt":  null,
		"Role":  "Senior Ruby Engineer",
		"Company":  "BrasSoft",
		"Location":  "New York",
		"Remote":  true,
		"Link":  "vaga.com/vagazika/33298",
		"Salary":  190000
	},
	"message":  "operation from handler: show-opening successful"
}
``` 

### POST /opening

Create a new opening.

**Request:**

```http
POST /opening
Content-Type: application/json
```

Body
```json
{
    "role": "Senior React Engineer",
    "company": "UltraSoft",
    "location": "New York",
    "remote": true,
    "salary": 190000,
    "link": "vaga.com/vagazika/33298"
}
``` 

**Response:**

```json
{
    "data": {
        "ID": 5,
        "CreatedAt": "2024-06-23T18:34:37.188478-03:00",
        "UpdatedAt": "2024-06-23T18:34:37.188478-03:00",
        "DeletedAt": null,
        "Role": "Senior React Engineer",
        "Company": "UltraSoft",
        "Location": "New York",
        "Remote": true,
        "Link": "vaga.com/vagazika/33298",
        "Salary": 190000
    },
    "message": "operation from handler: create-opening successful"
}
``` 

### PUT /opening

Update an existing opening.

**Request:**

```http
PUT /opening?id={id}
Content-Type: application/json
```

Body:
```json
{
    "company": "Geegle"
}
``` 

**Response:**

```json
{
    "data": {
        "ID": 2,
        "CreatedAt": "2024-06-23T16:53:37.332476-03:00",
        "UpdatedAt": "2024-06-23T18:35:29.182524-03:00",
        "DeletedAt": null,
        "Role": "Senior Ruby Engineer",
        "Company": "Geegle",
        "Location": "New York",
        "Remote": true,
        "Link": "vaga.com/vagazika/33298",
        "Salary": 190000
    },
    "message": "operation from handler: update-opening successful"
}
``` 

### DELETE /opening

Delete an existing opening by its ID.

**Request:**

```http
DELETE /opening?id={id}
``` 

**Response:**

```json
{
    "data": {
        "ID": 4,
        "CreatedAt": "2024-06-23T16:54:47.778492-03:00",
        "UpdatedAt": "2024-06-23T16:54:47.778492-03:00",
        "DeletedAt": "2024-06-23T18:36:43.719817-03:00",
        "Role": "Senior React Engineer",
        "Company": "UltraSoft",
        "Location": "New York",
        "Remote": true,
        "Link": "vaga.com/vagazika/33298",
        "Salary": 190000
    },
    "message": "operation from handler: delete-opening successful"
}
``` 

### GET /openings

Retrieve a list of all openings.

**Request:**

```http
GET /openings
``` 

**Response:**

```json
{
    "data": [
        {
            "ID": 2,
            "CreatedAt": "2024-06-23T16:53:37.332476-03:00",
            "UpdatedAt": "2024-06-23T18:35:29.182524-03:00",
            "DeletedAt": null,
            "Role": "Senior Ruby Engineer",
            "Company": "Geegle",
            "Location": "New York",
            "Remote": true,
            "Link": "vaga.com/vagazika/33298",
            "Salary": 190000
        },
        {
            "ID": 3,
            "CreatedAt": "2024-06-23T16:54:44.010187-03:00",
            "UpdatedAt": "2024-06-23T16:54:44.010187-03:00",
            "DeletedAt": null,
            "Role": "Senior Python Engineer",
            "Company": "UltraSoft",
            "Location": "New York",
            "Remote": true,
            "Link": "vaga.com/vagazika/33298",
            "Salary": 190000
        },
        {
            "ID": 5,
            "CreatedAt": "2024-06-23T18:34:37.188478-03:00",
            "UpdatedAt": "2024-06-23T18:34:37.188478-03:00",
            "DeletedAt": null,
            "Role": "Senior React Engineer",
            "Company": "UltraSoft",
            "Location": "New York",
            "Remote": true,
            "Link": "vaga.com/vagazika/33298",
            "Salary": 190000
        }
    ],
    "message": "operation from handler: list-openings successful"
}
``` 

## Installation

To install and run the API, follow these steps:

1.  Clone the repository:
    
    `git clone git@github.com:guilhermemcardoso/go-opportunities-api.git` 
    
2.  Navigate to the project directory:
 
    `cd go-opportunities-api` 
    
3.  Run the API locally:   

    `go run main.go` 

4. Access the endpoints at `localhost:8080`
