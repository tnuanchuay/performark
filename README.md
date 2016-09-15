# Performark
[![Build Status](https://travis-ci.org/ntossapo/performark.svg?branch=master)](https://travis-ci.org/ntossapo/performark)

โปรแกรมทดสอบ Load ของเว็บแอปพลิเคชันต่างๆ หรือ เอพีไอ ที่ทำงานผ่านโปรโตคอล HTTP, HTTPS
สามารถบอกขีดความสามารถของโปรแกรมที่ท่านพัฒนาในการรองรับจำนวนผู้ใช้งานและข้อมูลต่างๆ
อธิบายข้อมูลจากผลการทดสอบด้วยกราฟชนิดต่างๆ เพื่อให้เห็นความสามารถของแอปพลิเคชันในแต่ละสภาพแวดล้อม
โปรแกรมทำงานอยู่บน wg/wrk ซึ่งเป็น benchmark ที่มีความนิยมในระดับหนึ่ง และเป็นโอเพนซอร์ส ให้ผลลัพธ์, ค่าตัวแปร เช่น
* Request/Second
* Latency
* Data-Transfer/Second
* Socket Error
* Non-2xx Response


### require
* [wg/wrk](https://github.com/wg/wrk)
* [golang](https://golang.org/)
* [mongodb](https://www.mongodb.com/)

### Available
* Linux
* OSX

### Installation
```
 ./install.sh
```

### Run
* ```./performark```
* Open browser and go http://127.0.0.1:8080

### Getting Start
* รันแอปพลิเคชันด้วย ```./performark```
* เปิด browser http://127.0.0.1:8080
* ที่ช่อง URL ใส่ url ที่ server ของท่านกำลังทำงานอยู่
* Run
* รอผลลัพธ์
* เมื่อทำงานเสร็จ สามารถดูผลลัพธ์การทดสอบได้ โดยคลิกดูที่การทดสอบที่ขึ้นมาด้านล่างของหน้าแรก

