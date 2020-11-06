import { check, sleep } from 'k6';
import http from 'k6/http';

export let options = {
    duration: '1m',
    rps: 2,
};

export default function () {
    let res = http.get(__ENV.API_SERVER_URL);
    check(res, {
        'status was 200': (r) => r.status == 200,
        'status was 429': (r) => r.status == 429,
    });
}
