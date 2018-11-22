<template>
    <div id="query">
        <div class="insertform">
            <h6>Query</h6>
            <label for="term">Search: </label><input type="text" id="term" v-on:keyup="search" v-model="term"/>
            <div class="row" v-if="entries.length > 0">
                <div class="col-md-2"><button @click="selectAll">SELECT ALL</button></div>
                <div class="col-md-2"><button @click="addEntries">ADD TO EVENTS</button></div>
            </div> 
            <div v-for="e in entries" v-bind:entry="e" v-bind:key="e.key" class="row">
                <div class="col-md-1">{{ e.type }}</div>
                <div class="col-md-1"><input type="checkbox" v-model="selected" :value="e.key" /></div>
                <div class="col-md-3">{{ e.key }}</div>
                <div class="col-md-7">{{ e.value }}</div>
            </div>
            
        </div>
    </div>
</template>

<script>
import api from '../services';

const { query, redisKey } = api; 

export default {
    methods: {
        selectAll() {
            let selected = [];
            this.entries.forEach(e => {
                selected.push(e.key)
            });
            this.selected = selected;
        },
        addEntries() {
            const filtered = this.entries.filter(e => this.selected.indexOf(e.key) !== -1);
            filtered.map(e => this.$store.dispatch('addEvent', e));
        },
        search() {
            this.entries = [];
            if (this.term.length > 3) {
                query(this.term)
                    .then(async ({ data }) => {
                        const strkeys = await Promise.all(data.string.map(async s => {
                            const { data } = await redisKey(s, 'STRING');
                            return data;
                        }));
                        const hashkeys = await Promise.all(data.hash.map(async h => { 
                            const { data } = await redisKey(h, 'HASH'); 
                            return data;
                        }));
                        const setkeys = await Promise.all(data.set.map(async s => {
                            const { data } = await redisKey(s, 'SET');
                            return data;
                        }));
                        console.log(strkeys, hashkeys, setkeys);
                        const results = strkeys.concat(hashkeys, setkeys);
                        this.entries = results;
                    })
            }
        }
    },
    data () {
        return {
            term: '',
            entries: [],
            selected: []
        }
    }
};
</script>