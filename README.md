# 🔎 Port Scanner in Go

A simple concurrent port scanner written in Go.  
This tool scans a given host for open TCP ports within a specified range.

---

## 🚀 Features
- ⚡ Concurrent scanning with goroutines
- 🎯 Customizable port range
- ⏱️ Timeout for each connection attempt (3 seconds)
- 🪶 Lightweight and easy to use

---

## 🛠️ Usage

### Run directly
```bash
go run portscan.go <host> <start_port> <end_port>
````

### Example

```bash
go run portscan.go example.com 1 1000
```

This will scan ports **1 to 1000** on `example.com`.

### Compile binary

```bash
go build portscan.go
./portscan <host> <start_port> <end_port>
```

---

## 📌 Sample Output

```text
Scanning example.com ports 1-1000...

Port 22 open
Port 80 open
Port 443 open

Scan completed!
```

---

## 📦 Requirements

* Go 1.18 or newer
* Internet access for scanning remote hosts

---

## ⚠️ Disclaimer

This project is for **educational purposes only**.
Do not use this tool to scan systems without proper authorization.
Unauthorized scanning may be illegal.
