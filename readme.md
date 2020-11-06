## 作法
1. 使用 redis 達成 rate limiting
2. request ip 會產生一組 redis key 為 ratelimit:{{IP}}:{{UUID}}
3. 每一組 key 都會有 60 秒的 TTL
4. 每次 request 會取得 ratelimit:{{IP}}: * 共有多少個, 若小於設定值(一分鐘60個) 則通過, 否則為 Error
5. 由 middleware 為阻擋, 若通過則進入 /v1/hello 去取得當前 ip 總共的 request 數量 進行回傳
 


## 建置
1. 測試使用 k6 (https://github.com/loadimpact/k6)
2. 需使用 redis 
3. 有關於 server 的設定擋在 deploy/config/app.yaml

### Local
```shell
docker run --name redis-lab -p 6379:6379 -d redis
make server
API_SERVER_URL=http://0.0.0.0:8080/v1/hello k6 run test/http_get.js 
```

### Cloud (Heroku)
* 第一次 request 須等一段時間 
```shell
API_SERVER_URL=https://dcard-test.herokuapp.com/v1/hello k6 run test/http_get.js 
```

## 結果
一分鐘總共打 121 次 (k6 最後看起來會多打一次) , rps 為 2
```
running (1m00.0s), 0/1 VUs, 121 complete and 0 interrupted iterations
default ✓ [======================================] 1 VUs  1m0s

    ✗ status was 200
     ↳  49% — ✓ 60 / ✗ 61
    ✗ status was 429
     ↳  50% — ✓ 61 / ✗ 60

    checks.....................: 50.00% ✓ 121 ✗ 121
    data_received..............: 67 kB  1.1 kB/s
    data_sent..................: 10 kB  173 B/s
    http_req_blocked...........: avg=9.97µs   min=4µs     med=6µs      max=328µs    p(90)=10µs    p(95)=21µs
    http_req_connecting........: avg=2µs      min=0s      med=0s       max=242µs    p(90)=0s      p(95)=0s
    http_req_duration..........: avg=5.69ms   min=2.47ms  med=5.69ms   max=12.29ms  p(90)=8.63ms  p(95)=9.32ms
    http_req_receiving.........: avg=52.45µs  min=36µs    med=46µs     max=257µs    p(90)=65µs    p(95)=80µs
    http_req_sending...........: avg=30.3µs   min=15µs    med=26µs     max=143µs    p(90)=44µs    p(95)=60µs
    http_req_tls_handshaking...: avg=0s       min=0s      med=0s       max=0s       p(90)=0s      p(95)=0s
    http_req_waiting...........: avg=5.61ms   min=2.42ms  med=5.6ms    max=11.89ms  p(90)=8.57ms  p(95)=9.24ms
    http_reqs..................: 121    2.01585/s
    iteration_duration.........: avg=495.89ms min=13.06ms med=500.05ms max=506.51ms p(90)=502.8ms p(95)=503.65ms
    iterations.................: 121    2.01585/s
    vus........................: 1      min=1 max=1
    vus_max....................: 1      min=1 max=1
```