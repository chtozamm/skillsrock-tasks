# Apache Benchmark Output

### Command

```bash
ab -n 10000 -c 1000 http://localhost:8080/tasks
```

**Server Hostname:** localhost  
**Server Port:** 8080

**Document Path:** /tasks  
**Document Length:** 4 bytes

**Concurrency Level:** 1000  
**Time taken for tests:** 2.592 seconds  
**Complete requests:** 10000  
**Failed requests:** 0  
**Total transferred:** 1300000 bytes  
**HTML transferred:** 40000 bytes  
**Requests per second:** 3858.68 \[#/sec\] (mean)  
**Time per request:** 259.156 \[ms\] (mean)  
**Time per request:** 0.259 \[ms\] (mean, across all concurrent requests)  
**Transfer rate:** 489.87 \[Kbytes/sec\] received

## Connection Times (ms)

|             | min | mean | +/-sd | median | max |
| ----------- | --- | ---- | ----- | ------ | --- |
| Connect:    | 0   | 7    | 11.4  | 1      | 50  |
| Processing: | 90  | 240  | 59.4  | 230    | 507 |
| Waiting:    | 40  | 237  | 60.0  | 226    | 507 |
| Total:      | 90  | 247  | 63.5  | 233    | 525 |

## Percentage of Requests Served Within a Certain Time (ms)

| Percentile | Time (ms)             |
| ---------- | --------------------- |
| 50%        | 233                   |
| 66%        | 251                   |
| 75%        | 259                   |
| 80%        | 265                   |
| 90%        | 304                   |
| 95%        | 410                   |
| 98%        | 469                   |
| 99%        | 490                   |
| 100%       | 525 (longest request) |
