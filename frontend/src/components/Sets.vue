<template>
    <div class="container-fluid">
        <div class="row insertform">
            <div>
                <label for="search">Search</label>
                <input type="text" id="search" v-model="search" />
                <button @click="reset">CLEAR</button>
                <button @click="invert">INVERT SELECTION</button>
            </div>
        </div>
        <div class="row">
            <div class="col-md-4 column">
                <h6>SETS</h6>
                <div class="evt selectable" @click="select(e.id)" v-for="e in entries" v-bind="e" v-bind:key="e.id">
                    {{ e.key }}
                </div>
            </div>
            
            <div class="col-md-4 column">
                <h6>SELECTED</h6>
                <div v-show="sel1" class="evt">
                    <div><strong>{{ sel1key }}</strong></div>
                    <div>{{ sel1.value }}</div>
                </div>
                <div v-show="sel2" class="evt">
                    <div><strong>{{ sel2key }}</strong></div>
                    <div>{{ sel2.value }}</div>
                </div>
            </div>
            
            <div class="col-md-4" v-show="sel2key !== ''">
                <div class="insertform">
                    <h6>UNION</h6>
                    <div class="evt">{{ union }}</div>
                    <div class="command"><pre>{{ unioncommand }}</pre></div>
                </div>
                <div class="insertform">
                    <h6>DIFF</h6>
                    <div class="evt">{{ diff }}</div>
                    <div class="command"><pre>{{ diffcommand }}</pre></div>
                </div>
                <div class="insertform">
                    <h6>INTERSECT</h6>
                    <div class="evt">{{ intersect }}</div>
                    <div class="command"><pre>{{ intersectcommand }}</pre></div>
                </div>

            </div>
        </div>
    </div>
</template>

<style>
.selectable {
    cursor: pointer ;
}
.column {
    border-right: 1px solid #ccc;
}
.command {
    font-size: 0.8em;
    background: #ccc;
}
</style>

<script>
export default {
    name: 'Sets',
    computed: {
        sel1key() {
            return !!this.sel1 ? this.sel1.key : '';
        },
        sel2key() {
            return this.sel2.key;
        },
        entries(){
            const sets = this.$store.getters.EVENTS.filter(e => e.type === 'SET').filter(e => e.key.indexOf(this.search) !== -1);
            this.sets = {};
            sets.forEach(s => {
                this.sets[s.id] = s;
                console.log('SET', s.id, this.sets[s.id]);
            })
            return this.sets;
        },
        union() {
            const a = this.sel1.value.split(",").map(v => v.trim());
            const b = this.sel2.value.split(",").map(v => v.trim());
            const s = new Set(a.concat(b));
            console.log(a, b, s)
            return [...s].join(", ");
        },
        unioncommand() {
            const a = this.sel1.key;
            const b = this.sel2.key;
            return `SUNION ${a} ${b}`;
        },
        diff() {
            const a = this.sel1.value.split(",").map(v => v.trim());
            const b = this.sel2.value.split(",").map(v => v.trim());
            const c = new Set(a.filter(e => b.indexOf(e) === -1));
            return [...c].join(", ");
        },
        diffcommand() {
            const a = this.sel1.key;
            const b = this.sel2.key;
            return `SDIFF ${a} ${b}`;
        },
        intersect() {
            const a = this.sel1.value.split(",").map(v => v.trim());
            const b = this.sel2.value.split(",").map(v => v.trim());
            const c = new Set(a.filter(e => b.indexOf(e) !== -1));
            return [...c].join(", ");
        },
        intersectcommand() {
            const a = this.sel1.key;
            const b = this.sel2.key;
            return `SINTER ${a} ${b}` 
        }
    },
    methods: {
       invert() {
           const x = this.sel1;
           this.sel1 = this.sel2;
           this.sel2 = x;
       },
       reset(){
            this.sel1 = {
                key: '',
                type: '',
                value: '',
                id: ''
            };
            this.sel2 = {
                key: '',
                type: '',
                value: '',
                id: ''
            };
            this.sets = {};
        },
        select(id) {
            console.log(id);

            if(this.sel1.key === '') {
                this.sel1 = this.sets[id];
                console.log(this.sel1); 
                return;
            }
            this.sel2 = this.sets[id];
        }
    },
    data() {
        return {
            search: '',
            sel1: {
                key: '',
                type: '',
                value: '',
                id: ''
            },
            sel2: {
                key: '',
                type: '',
                value: '',
                id: ''
            },
            sets: {}
        };
    }
}
</script>
