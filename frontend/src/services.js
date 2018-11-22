import { post } from 'axios';

const BASE_URL = 'http://localhost:1666';
const setURL = (key, value) => `${BASE_URL}/set/${key}/${value}`;
const hsetURL =  (key, field, value) => `${BASE_URL}/hset/${key}/${field}/${value}`;
const saddURL = (key, value) => `${BASE_URL}/sadd/${key}/${value}`;

const redisSET = (key, value) => post(setURL(key, value));
const redisHSET = (key, field, value) => post(hsetURL(key, field, value));
const redisSADD = (key, value) => post(saddURL(key, value));

export default {
    redisSET,
    redisHSET,
    redisSADD
};