# Simple API
A simple api for sending and retrieving message. 

## Installation
Make sure you have following software installed in your system
* go 1.10.4
* dep 

### Clone repository and dependency
Clone this repository in your `$GOPATH/src`
````
git clone https://github.com/Rikadewi/simple-api.git
cd simple-api
dep ensure
````
### Run on local server
```
go run api.go
```
### Access
* host: `127.0.0.1`
* port: `8000`

## API Endpoints
List of available endpoints

#### POST /send/{msg}
Send a message through path string parameter. If success, server will response with `OK`.
#### GET /fetch
Fetch all messages that has been sent out through API. 

**Response**
```
{
    "messages": [
        "messageOne",
        "messageTwo",
        "messageThree"
    ]
}
```
#### /ws
Establish long live connection using websocket.

## Demo
First start the server.
```
go run api.go
```
Open `client/index.html` with your favorite browser (Chrome, Firefox, etc). Try to send a message to server. You can use this request example that sent `hi` message.
```
curl -X POST \
  http://127.0.0.1:8000/send/hi \
```
Head over to the browser and see the result. You should see `hi` in your browser now. Try a couple of other curl's and see the difference.