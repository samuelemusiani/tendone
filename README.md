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

The following list is based on the web interface of the Tenda i27 AP. I don't
have plans to implement all of them, but I will try to implement the most important ones.

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
        - [x] Reboot
        - [ ] Reset
        - [ ] Upgrade
        - [x] Backup
        - [ ] Restore
        - [x] Led
    - [ ] Account
    - [ ] System Log
    - [ ] Diagnostic Tool
    - [ ] Uplink Detection
