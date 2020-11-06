import { check, sleep } from 'k6';
import http from 'k6/http';

export let options = {
    duration: '60s',
    vus: 1,
    rps: 2
};
export default function () {
    let res = http.get('http://0.0.0.0:8080/v1/hello');
    check(res, {
        'status was 200': (r) => r.status == 200,
        'status was 429': (r) => r.status == 429,
    });
}
