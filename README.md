# Performark
[![Build Status](http://jenkins.ntossapo.me:8080/buildStatus/icon?job=performark-master)](http://jenkins.ntossapo.me:8080/job/performark-master/)

Web Benchmark to test your API performance. Install in your develop machine and test your api before deployment.
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
* run and wait for result
* see the report

### Screen Shot
![Home](https://raw.githubusercontent.com/ntossapo/performark/master/screenshot/1.png)
![Report](https://raw.githubusercontent.com/ntossapo/performark/master/screenshot/2.png)
