# Tendone

A simple Go library for the Tenda i27 AP. This could be used to get and set parameteres in a Tenda i27 AP from a go program.

## Usage

First create a session based on the uri of your Tenda i27 AP:
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


## Todo

- [x] Status
    - [x] System Status
    - [x] Wireless Status
    - [x] Traffic Statistics
    - [x] Client List
- [ ] Quick Setup
- [x] Internet settings
- [ ] Wireless
    - [ ] SSID
    - [ ] RF Settings
    - [ ] RF Optimization
    - [ ] Frequency Analysis
    - [ ] WMM
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
    - [ ] Account
    - [ ] System Log
    - [ ] Diagnostic Tool
    - [ ] Uplink Detection
