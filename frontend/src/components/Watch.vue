<template>
    <div class="container-fluid">
        <div class="row insertform">
            <label for="channel">Filter</label> <input type="text" id="channel" v-model="expr" /> Type: <input type="text" id="type" v-model="type" />
            <button @click="reset">RESET</button>
        </div>
        <div class="row">
            <div class="col-md-1"><strong>TYPE</strong></div>
            <div class="col-md-3"><strong>KEY</strong></div>
            <div class="col-md-8"><strong>VALUE</strong></div>
            <Event v-for="e in entries" v-bind:event="e" v-bind:key="e.id"/>
        </div>
    </div>
</template>
<script>
import Event from './Event.vue';

export default {
    name: 'Watch',
    components: {
        Event
    },
    computed: {
        entries() {
            return this.$store.getters.EVENTS
                .filter(e => e.key.indexOf(this.expr) !== -1)
                .filter(e => e.type.toLowerCase().indexOf(this.type.toLowerCase()) !== -1);
        }
    },
    data () {
        return {
            expr: '',
            type: ''
        };
    },
    methods: {
        reset() {
            this.$store.dispatch('reset')
        }
    },
    mounted() {
    }
}

</script>