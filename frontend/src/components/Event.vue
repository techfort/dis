<template>
    <div class="container-fluid">
        <div class="row">
            <div class="col-md-1 evt">
                <span class="badge"
                    v-bind:class="{ 'badge-primary': event.type === 'KEY', 'badge-danger': event.type === 'SET', 'badge-warning': event.type === 'HASH'}">
                    {{ event.type }}
                </span>
            </div>
            <div class="col-md-3 keycol evt">{{ event.key }}</div>
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
export default {
    name: 'Event',
    props: ['event']
}
</script>