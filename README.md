# Tendone

A simple Go library for the Tenda i27 AP. This could be used to get and set parameteres in a Tenda i27 AP from a go program.

## Usage

First create a session based on the uri of your Tenda i27 AP:
```go
s := tendone.NewSession("http://tenda-ap.int")
```

Login the session:
```go
logged, err := s.Login("admin", "uq1hMoCHEZt3NEsFxXT2")
if err != nil {
    log.Fatal(err)
}
fmt.Println("Logged in:", logged)
```


## Todo

- [x] Status
- [ ] Quick Setup
- [x] Internet settings
- [ ] Wireless
- [ ] Advanced
- [ ] Tools
