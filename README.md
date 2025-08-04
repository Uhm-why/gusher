# 💻 Gusher (Go-Usher - The Go-based Networking Usher)

## ❗This Project is a Work in Progress❗

Gusher is a tool written in Go to provide the functionality needed to manage, configure, and update distributed systems.

## 🙏Goals

To provide an all-in-one system for managing, configuring, and update remote systems.

### 😕 What this includes

- Sending Files via a Secure Protocol
- Reading/Implementing Configuration Files (Windows, Linux and MacOS)
- Starting VMs and Containers
- Reporting on Remote System Status

## ❓Why Not Just Use [insert config manager name here]

Because I have tried them and they don't play well with the tools that I like to use or provide all of the functionality that I would like.

Also, I am doing a lot of this to learn about how different mechanisms work. The end product is a plus that I plan on providing to the open-source community.

## 📋 Current Development Progress/Roadmap

- [x] Implement Basic Chunk Sizing Logic
- [x] File Hashing
- [ ] Read Files and Write Chunks
- [ ] Chunk Hashing
- [ ] Re-integrate Chunks into Complete File
- [ ] Chunk and File Hash Verification
- [ ] Implementing a Secure Protocol (not SSH)
- [ ] Transferring Files Over Secure Protocol
- [ ] Reading and Implementing Configurations
- [ ] Spinning up VMs/Containers
- [ ] Automatically Adjusting Firewall Rules for VMs/Containers
- [ ] System Status (File Transfer Progress, Up/Down Time, Running VMs/Containers, etc.) Reports
- [ ] API
- [ ] GUI
- [ ] Code Optimization (to be done between each item in the roadmap). This will probably not be checked off for a while, if ever.

## Developed Using

- VSCodium
- [Cobra](github.com/spf13/cobra)
