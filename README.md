### Prerequisites
1. Docker 
2. Docker Compose
3. PostMan (Rest Client)

### How to
* Change the config in `.env.example` for app config, and database config in `database.env`. Standard config will be like
```
// .env.example
APP_ENV=production
APP_PORT=8080

DB_HOST=edufund-db
DB_PORT=5432
DB_USER=andrechristikan
DB_PASS=123456
DB_NAME=edufund


// database.env
POSTGRES_USER=andrechristikan
POSTGRES_PASSWORD=123456
POSTGRES_DB=edufund
```

* Optional. If you change `APP_PORT`, don't forget to change EXPORT PORT in `docker-compose.yml` and `dockerfile` for the app.

* Run bash script to up docker service
> sh run.sh

* Run bash script to down docker service
> sh end.sh


### Endpoint
Let assume we running application on localhost port 8080. Detail endpoints will attach in `edufund.endpoint.json`, you must import into PostMan

Test endpoint

* **URL**

  localhost:8080/test

* **Method:**
  
  `GET` 
  
*  **URL Params**

   Don't have params

* **Body**

  Don't have body

* **Success Response:**

  * **Code:** `200` <br />
    **Content:** 
    ```
    {
        "message": "test endpoints"
    }
    ```
 
* **Error Response:**

  Don't have error


---
