## env to yaml
### Objective
This will convert env formated secret to yaml format which can be used in `values.yaml` file for k8s deployment
### How to run it
* Edit `input.env` file
* run `go run envtoyaml/main.go`
* Output will be at `envtoyaml/output.yaml` file