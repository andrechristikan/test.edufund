### Prerequisites
1. Docker 
2. Docker Compose
3. PostMan (Rest Client)

### How to
Run bash script to up docker service
> sh run.sh

Run bash script to down docker service
> sh end.sh


### Endpoint
Let assume we running application on localhost port 8080. Detail endpoints will attach in `edufund.endpoint.json`, you must import into PostMan

#### Test
----
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


