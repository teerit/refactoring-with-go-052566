# Thaiwin API

3 route paths

~~~
/recently
    get currently checked in
/checkin
    check-in to any places
        input {
            "id": 1234,
            "place_id": "4321"
        }
        output {
            "density": "ok"
        }
/checkout
        input {
            "id": 1234,
            "place_id": "4321"
        }
        output {}        
~~~

```curl
curl --location 'http://127.0.0.1:8000/checkin' \
--header 'Content-Type: application/json' \
--data '{
    "id": 1234,
    "place_id": 4321
}'
```

```curl
curl --location 'http://127.0.0.1:8000/checkout' \
--header 'Content-Type: application/json' \
--data '{
    "id": 1234,
    "place_id": 4321
}'
```

```curl
curl --location --request POST 'http://127.0.0.1:8000/recently'
```

```curl

```curl