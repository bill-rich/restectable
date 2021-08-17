# Restectable REST Message API
Restectable is a respectable REST API for storing and retrieving messages.

## About
Restectable stores messages in memory to be retrieved at a later time using the
SHA256 hash of the message content. If a hash collision occurs, the newest
instance of the message will be stored and the old one will be lost.

The restectable server always runs locally on port 8080.

## How To Run

```
go get github.com/bill-rich/restectable
restectable
```

## API Documentation

* ***POST /message***: Add a new message to be stored.
  * Body: JSON ```{ "content": "coolmessage" }```
  * Response: 200, JSON ```{ "hash": "f719a89a21181e181e16e3f42c1740b712a2061b72cd5fe41a41da472110ec89" }```
* ***GET /message/{sha256hash}***: Retrieves the message with the associated
  hash.
  * Response: 200, Message retrieved. BODY: JSON ```{ "message": "coolmessage" }```
  * Response: 404, Message not found.


## Examples

Create a message on the server:
```
# curl -X POST http://127.0.0.1:8080/message -H "Content-Type: application/json" -d "{ \"content\": \"coolmessage\" }"
{"hash":"f719a89a21181e181e16e3f42c1740b712a2061b72cd5fe41a41da472110ec89"}
```

Retrieve a stored message:
```
# curl https://127.0.0.1:8080/message/f719a89a21181e181e16e3f42c1740b712a2061b72cd5fe41a41da472110ec89
{"content":"coolmessage"}
```
