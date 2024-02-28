# Golang Text Microservice Demo

## Running

`go run main.go`

## Interacting with the service

- `curl -X POST -d '{"input: "hello, world!"}' http://localhost:8080/spongebob`
- `curl -X POST -d '{"input: "wow"}' http://localhost:8080/is-palindrome`
- `curl -X POST -d '{"input: "this is a secret"}' http://localhost:8080/rot13`
    
## Deploying to Heroku

First set the heroku remote git repository: `heroku git:remote -a <heroku-app-name>`

Then push the code to heroku:
- `git push heroku main`

That's it. 😊 