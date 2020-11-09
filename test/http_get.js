import { check, sleep } from 'k6';
import http from 'k6/http';

export let options = {
    duration: '1m',
    rps: 2,
};

export default function () {
    let res = http.get(__ENV.API_SERVER_URL);
    check(res, {
        'status was 200 and response data.count = is 1 ~ 60': (r) => r.status == 200 && 0<r.json("data.count")<=60,
        'status was 429 and response error_message = "Error"': (r) => r.status == 429 && r.json("error_message") == "Error",
    });
}   
