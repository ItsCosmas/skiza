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
![A Screenshot of the Running Frontend on Browser](https://github.com/ItsCosmas/skiza/blob/main/demo/webpage.png) <br />
