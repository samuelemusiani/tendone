# Tendone

A simple Go library for the Tenda i27 AP. It could be used to get and set 
parameters in a Tenda i27 AP from a Go program.

## Usage

First create a session based on the URI of your Tenda i27 AP:
```go
s := tendone.NewSession("http://tenda-ap.int")
```

Login the session:
```go
logged, err := s.Login("admin", "password")
if err != nil {
    log.Fatal(err)
}
fmt.Println("Logged in:", logged)
```

Get the system status:
```go
status, err := s.GetSysStatus()
if err != nil {
    log.Fatal(err)
}
fmt.Println(status)
```

Logout the session:
```go
slogged, err := s.Logout()
if err != nil {
    log.Fatal(err)
}
fmt.Println("Logged out:", slogged)
```

For the full docs see the [godoc](https://pkg.go.dev/github.com/samuelemusiani/tendone).

## Todo

The following list is based on the web interface of the Tenda i27 AP. I don't
have plans to implement all of them, but I will try to implement the most 
important ones.

- [x] Status
    - [x] System Status
    - [x] Wireless Status
    - [x] Traffic Statistics
    - [x] Client List
- [ ] Quick Setup
- [x] Internet settings
- [ ] Wireless
    - [x] SSID
    - [ ] RF Settings
    - [ ] RF Optimization
    - [x] Frequency Analysis
    - [x] WMM
    - [ ] Access Control
    - [ ] Advanced Settings
    - [ ] QVLAN Settings
    - [ ] WiFi Schedule
- [ ] Advanced
    - [ ] Traffic Control
    - [ ] Cloud Maintenance
    - [ ] Remote Management
- [ ] Tools
    - [ ] Date & Time
    - [ ] Maintenance
        - [x] Reboot
        - [ ] Reset
        - [ ] Upgrade
        - [x] Backup
        - [ ] Restore
        - [x] Led
    - [ ] Account
    - [x] System Log
    - [ ] Diagnostic Tool
    - [ ] Uplink Detection
