# ftp-server

An implementation of an FTP server in Golang.

## Overview

Originally following [this article](https://medium.com/better-programming/how-to-write-a-concurrent-ftp-server-in-go-part-1-3904f2e3a9e5) though I found it's missing or incorrectly implementing the FTP commands and requires paying to view the follow up article.

If you're interested in writing your own Golang FTP server, it's an excellent starting point and the issues in the article encourage you to search and discover the correct implementation yourself.

## TODO
  - [ ] allow port and storage root to be passed as flags
  - [ ] testing
    - [ ] note -> since net uses file writes for response this can be mocked and caught
  - [ ] implement passive mode (EPSV)
  - [ ] implement IPv6 in EPRT
  - [ ] implement custom commands?
    - [ ] this requires a FTP client supporting custom commands
  - [ ] CWD bounds, restrict file system access
    - [ ] deny paths which evaulate outside of storage root 
  - [ ] users and passwords?
  - [ ] containerize
  - [ ] rename go module to not be 'music-server'