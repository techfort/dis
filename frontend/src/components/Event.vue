<template>
    <div class="container-fluid">
        <div class="row">
            <div class="col-md-1 evt">
                <span class="badge"
                    v-bind:class="{ 'badge-primary': event.type === 'STRING', 'badge-danger': event.type === 'SET', 'badge-warning': event.type === 'HASH'}">
                    {{ event.type }}
                </span>
            </div>
            <div class="col-md-3 keycol evt">{{ event.key | truncate }}</div>
            <div class="col-md-8 evt">{{ event.value }}</div>
        </div>
    </div>
</template>
<style>
.keycol {
    background: #eeeeee;
}
.evt {
    border-bottom: 1px solid #ddd;
}
</style>


<script>
import filters from '../filters';

export default {
    name: 'Event',
    props: ['event'],
    filters
}
</script>