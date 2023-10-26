# HOW TO RUN

1. install gorm, sqlite driver, and gin

```bash
$ go get -u gorm.io/gorm
$ go get -u gorm.io/driver/sqlite
$ go get -u github.com/gin-gonic/gin
```

2. Run

```bash
$ go run main.go
```

3. Import Postman collections from 

> ./rizqi_golang-assignment1_postman.json


### Note:

There is **an extra** API: PUT /orders/ to update both Order and **Upsert** Items.

**Pros** -> We don't need to delete all items, then re-create it everytime Client update an Order.

**Cons** -> There are extra field need for the Client to filled in JSON Body:

- **REQUIRE** Order.ID
- **OPTIONAL** Item.id *If present, it would update the item. Otherwise it would create new item*


# Projects : BUILD REST API IN GOLANG 
Setelah mempelajari seputar API dan Database, maka untuk membantu kamu lebih memahami ini lebih jauh, kerjakan project Assignment ini.

## REQUIREMENT :

- ORM: Gunakan native ORM dari database/sql atau GORM
- Web Server: Boleh gunakan gin framework untuk routing atau gunakan native http server golang

## DATABASE :

- Db name : score_assignment

## Table and fields :

### items
- id
- name
- description
- quantity
- order_id
- created_at
- updated_at

### orders
- id
- customer_name
- ordered_at
- created_at
- updated_at

## TASK :
### Make endpoint to :
#### Create Order
    URL: http://localhost:8080/orders
    Method: POST
    Body raw json: {"orderedAt":"2019-11-09T21:21:46+00:00","customerName":"Fitri","items":[{"name":"iPhone","description":"iPhone 11 Pro","quantity":2}]}
#### Get All Order
    URL: http://localhost:8080/orders
    Method: GET
#### Update Order
    URL: http://localhost:8080/orders/:id
    Method: PUT
    Body raw json: {"orderedAt":"2022-11-09T21:21:46+00:00","customerName":"Fitri Tesss Editttt","items":[{"name":"iPhone Edittt","description":"iPhone Pro","quantity":2000}]}
#### Delete Order
    URL: http://localhost:8080/orders/:id
    Method: DELETE



https://documenter.getpostman.com/view/7183159/2s9YR55Zh5#16ce62e8-712f-44fe-9bed-aeb8ff6e941c