import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
    stages: [
        { duration: '10s', target: 10000 }, // Ramp up to 500 users
        { duration: '20s', target: 25000 }, // Stay at 1000 users
        { duration: '10s', target: 0 }, // Ramp down
    ],
};

const userIds = [
    "7aface47-7ce7-4b6a-ba16-392d19aa2785",
    "f51ee719-58fc-4f5f-9b32-3ebaa7648e52"
];

export default function () {
    let userId = userIds[Math.floor(Math.random() * userIds.length)];
    let headers = { 'User-ID': userId };
    let res = http.get(`http://cs:8080/api/v1/user/me?caching=true`, { headers: headers });

    check(res, {
        'is status 200': (r) => r.status === 200,
        'response time < 200ms': (r) => r.timings.duration < 200,
    });

    sleep(1);
}
