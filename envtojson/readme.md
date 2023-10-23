## env to json
### Objective
This will convert env formated secret to json format which can be used in aws `secret manager` or `app config`
### How to run it
* Edit `input.env` file
* run `go run envtojson/main.go`
* Output will be at `envtojson/output.json` file