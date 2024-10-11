## Motivation

- A number of scenarios with 3rd party integrations necessiate the need for a web hook to receive response.
- Case where a certain action is not completed immediately or is asynchronous, some partners opt to send a final response to a web hook.
- For example (Safaricom Daraja)[https://developer.safaricom.co.ke/] sends the final transaction status and info as a postback/callback to a designated http endpoint, therefore in development scenarios debugging can be hard as you need to create an internet reachable http endpoint where you'll receive your postbacks.
- This simple tools helps you t-shoot this scenario by providing a convenient way for you to quickly receive this postbacks.
- Have your 3rd party send the postbacks to this app via the endpoint `/api/v1/listener` and they are streamed real-time to a friendly web UI where you can inspect and t-shoot.

![A Screenshot of the Running Frontend on Browser](https://github.com/ItsCosmas/skiza/blob/main/demo/webpage.png) <br />

## Running and Developing locally

1. Create `.env` at server root, i.e.

```sh
cp server/.env.example server/.env
```

2. Download Swag for generating docs

```sh
go get -u github.com/swaggo/swag/cmd/swag
```

3. Run

- NOTE: You have to generate swagger docs before running the app.

```sh
# Terminal 1
cd server # Navigate to the go app
swag init # Generates Swagger
go run main.go # Run the go app

# Terminal 2
cd client # Navigate to the react app
npm run dev  # Run the react app
```

- API Route `http://localhost:8000/api/v1`
- Swagger Doc `http://localhost:8000/api/v1/docs`
- Web Socket URL `ws://127.0.0.1:8000/api/v1/ws`

---

![A Screenshot of the Running Frontend and Backend on Terminal](https://github.com/ItsCosmas/skiza/blob/main/demo/terminal.png) <br />

---

## Testing

1. Toggle Connection on the web app using the switch and if connected status changes to connected

2. Send a Sample Postback to your listener endpoint

```sh
curl -X 'POST' \\
  'http://127.0.0.1:8000/api/v1/listener' \\
  -H 'accept: application/json' \\
  -H 'Content-Type: application/json' \\
  -d '{\
        "event": "user_signup",\
        "data": {\
            "user_id": 12345,\
            "username": "example_user",\
            "email": "user@example.com"\
        },\
        "timestamp": "2024-10-09T12:34:56Z"\
      }'\
```

```sh
curl -X 'POST' \\
  'http://127.0.0.1:8000/api/v1/listener' \\
  -H 'accept: application/xml' \\
  -H 'Content-Type: application/xml' \\
  -d '<request>\
        <event>user_signup</event>\
        <data>\
            <user_id>12345</user_id>\
            <username>example_user</username>\
            <email>user@example.com</email>\
        </data>\
        <timestamp>2024-10-09T12:34:56Z</timestamp>\
      </request>'\
```
