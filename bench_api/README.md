All results are running locally.

# Simple GET Request

## Python

```
❰chase❙~/Dev/playground(git:bench_go_vs_python_api)❱✔≻ ab -n 10000 -c 5 http://0.0.0.0:8000/
This is ApacheBench, Version 2.3 <$Revision: 1903618 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 0.0.0.0 (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        uvicorn
Server Hostname:        0.0.0.0
Server Port:            8000

Document Path:          /
Document Length:        17 bytes

Concurrency Level:      5
Time taken for tests:   1.994 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      1420000 bytes
HTML transferred:       170000 bytes
Requests per second:    5016.28 [#/sec] (mean)
Time per request:       0.997 [ms] (mean)
Time per request:       0.199 [ms] (mean, across all concurrent requests)
Transfer rate:          695.62 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       5
Processing:     0    1   4.2      1     209
Waiting:        0    1   3.6      1     209
Total:          0    1   4.2      1     209

Percentage of the requests served within a certain time (ms)
  50%      1
  66%      1
  75%      1
  80%      1
  90%      1
  95%      1
  98%      2
  99%      2
 100%    209 (longest request)
 ```

 ## Go

 ```
 ❰chase❙~/Dev/playground(git:bench_go_vs_python_api)❱✔≻ ab -n 10000 -c 5 http://0.0.0.0:8080/
This is ApacheBench, Version 2.3 <$Revision: 1903618 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 0.0.0.0 (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        0.0.0.0
Server Port:            8080

Document Path:          /
Document Length:        17 bytes

Concurrency Level:      5
Time taken for tests:   0.477 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      1400000 bytes
HTML transferred:       170000 bytes
Requests per second:    20944.47 [#/sec] (mean)
Time per request:       0.239 [ms] (mean)
Time per request:       0.048 [ms] (mean, across all concurrent requests)
Transfer rate:          2863.50 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       1
Processing:     0    0   0.1      0       1
Waiting:        0    0   0.1      0       1
Total:          0    0   0.1      0       1

Percentage of the requests served within a certain time (ms)
  50%      0
  66%      0
  75%      0
  80%      0
  90%      0
  95%      0
  98%      1
  99%      1
 100%      1 (longest request)
 ```

# GET all data from a small (18 rows) postgres table

## Python

```
❰chase❙~/Dev/playground(git:go_api_updates)❱✔≻ ab -n 10000 -c 5 http://0.0.0.0:8000/items/
This is ApacheBench, Version 2.3 <$Revision: 1903618 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 0.0.0.0 (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        uvicorn
Server Hostname:        0.0.0.0
Server Port:            8000

Document Path:          /items/
Document Length:        1219 bytes

Concurrency Level:      5
Time taken for tests:   13.695 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      13460000 bytes
HTML transferred:       12190000 bytes
Requests per second:    730.17 [#/sec] (mean)
Time per request:       6.848 [ms] (mean)
Time per request:       1.370 [ms] (mean, across all concurrent requests)
Transfer rate:          959.78 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0      10
Processing:     2    7  33.9      6    1497
Waiting:        2    6  30.4      5    1497
Total:          3    7  33.9      6    1497

Percentage of the requests served within a certain time (ms)
  50%      6
  66%      6
  75%      6
  80%      6
  90%      7
  95%      8
  98%     10
  99%     13
 100%   1497 (longest request)
```

## Go

```
❰chase❙~/Dev/playground(git:go_api_updates)❱✔≻ ab -n 10000 -c 5 http://0.0.0.0:8080/items/
This is ApacheBench, Version 2.3 <$Revision: 1903618 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 0.0.0.0 (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        0.0.0.0
Server Port:            8080

Document Path:          /items/
Document Length:        1219 bytes

Concurrency Level:      5
Time taken for tests:   1.973 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      13440000 bytes
HTML transferred:       12190000 bytes
Requests per second:    5068.46 [#/sec] (mean)
Time per request:       0.986 [ms] (mean)
Time per request:       0.197 [ms] (mean, across all concurrent requests)
Transfer rate:          6652.35 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   3.0      0     297
Processing:     0    1   6.1      0     303
Waiting:        0    1   6.1      0     302
Total:          0    1   6.8      0     303

Percentage of the requests served within a certain time (ms)
  50%      0
  66%      0
  75%      1
  80%      1
  90%      2
  95%      3
  98%      4
  99%      4
 100%    303 (longest request)
 ```