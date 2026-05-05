# ftp-server

An implementation of an FTP server in Golang.

## Overview

Originally following [this article](https://medium.com/better-programming/how-to-write-a-concurrent-ftp-server-in-go-part-1-3904f2e3a9e5) though I found it's missing or incorrectly implementing the FTP commands and requires paying to view the follow up article.

I've had trouble with my implementation when using termscp and Cyberduck as the client but not had any issues with using ncftp on the command line.

## TODO
  - [ ] allow port and storage root to be passed as flags
  - [ ] testing
    - [ ] note -> since net uses file writes for response this can be mocked and caught
  - [x] CWD bounds, restrict file system access
  - [ ] users and passwords?
  - [ ] containerize
  - [x] rename go module to not be 'music-server'
  - [ ] implement commands:
    - [x] passive mode (PASV)
      - [x] extended passive mode (EPSV)
    - [x] LIST
    - [x] IPv6 in EPRT 
    - [ ] STAT
    - [ ] FEAT
    - [ ] DELE