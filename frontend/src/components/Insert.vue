<template>
    <div class="container-fluid">
       <div class="row insertform">
            <div>
                <h6>SET</h6>
                <label for="plainkey">Key</label><input type="text" id="plainkey" v-model="setKey" />
                <label for="setValue">Value</label><input id="setValue" type="text" v-model="setValue" />
                <button @click="execSET">SET</button>
                <div >{{ latestSET }}</div>
            </div>
        </div>
        <div class="row insertform">
            <div>
                <h6>HSET</h6>
                <label for="hasKey">Key</label><input type="text" id="hashKey" v-model="hsetKey" />
                <label for="hsetField">Field</label><input type="text" id="hsetField" v-model="hsetField" />
                <label for="hsetValue">Value</label><input type="text" v-model="hsetValue" id="hsetValue" />
                <button @click="execHSET">HSET</button>
                <div>{{ latestHSET }}</div>
            </div>
        </div>
        <div class="row insertform">
            <div>
                <h6>SADD</h6>
                <label for="saddKey">Key</label><input type="text" id="saddKey" v-model="saddKey" />
                <label for="saddValue">Value</label><input type="text" id="saddValue" v-model="saddValue" />
                <button @click="execSADD">SET</button>
                <div>{{ latestSADD }}</div>
            </div>
        </div>
    </div>
</template>

<script>
import api from '../services.js';

const { redisHSET, redisSADD, redisSET } = api;

export default {
    name: 'Insert',
    methods: {
        execSET() {
            redisSET(this.setKey.trim(), this.setValue.trim())
                .then(({ statusText }) => {
                    this.latestSET = statusText;
                })
                .catch(error => {
                    this.latestSET = error.response.data.error;
                });
        },
        execHSET() {
            redisHSET(this.hsetKey.trim(), this.hsetField.trim(), this.hsetValue.trim())
                .then(({ statusText }) => {
                    this.latestHSET = statusText;
                })
                .catch(error => {
                    this.latestHSET = error.response.data.error;
                });
        },
        execSADD() {
            redisSADD(this.saddKey.trim(), this.saddValue.trim())
                .then(({ statusText }) => {
                    this.latestSADD = statusText;
                })
                .catch(error => {
                    this.latestSADD = error.response.data.error;
                });
        }
    },
    data () {
        return {
            setKey: '',
            setValue: '',
            hsetKey: '',
            hsetField: '',
            hsetValue: '',
            saddKey: '',
            saddValue: '',
            latestSET: '',
            latestHSET: '',
            latestSADD: ''
        };
    }

};
</script>
