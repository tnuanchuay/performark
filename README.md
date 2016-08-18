# Performark
[![Build Status](https://travis-ci.org/ntossapo/performark.svg?branch=master)](https://travis-ci.org/ntossapo/performark)

ทดสอบโหลดของเว็บหรือ API โดยใช้ performark ติดตั้งไว้ในเครื่องที่ใช้ในการพัฒนา

Web Benchmark to test your API performance. Install in your development machine and test your api before deployment.

### require
* [wg/wrk](https://github.com/wg/wrk)
* [golang](https://golang.org/)
* [mongodb](https://www.mongodb.com/)

### Available
* Linux

### Installation
```
 ./install.sh
```

### Run
* ```./performark```
* Open browser and go http://127.0.0.1:8080

### Usage-Concept
* Enter URL to benchmark
* Enter http body-load if it have
* Run and wait for result
* See the report
