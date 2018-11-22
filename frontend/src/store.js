import Vuex from 'vuex';
import Vue from 'vue';

const addEntry = (state, e) => {
    if (e === null || e === undefined) {
        return;
    }
    e.id = state.events.length + 1;
    state.events.unshift(e);
};

const clear = (state) => {
    state.events = [];
};

Vue.use(Vuex);
export const store = new Vuex.Store({
    state: {
        entries: {
            KEY: {},
            SET: {},
            HASH: {}
        },
        events: [],
        socket: {
            isConnected: true,
            message: '',
            reconnectError: false
        } 
    },
    mutations: {
        ADD_EVENT(state, e) {
            addEntry(state, e);
        },
        CLEAR(state) {
            clear(state);
        },
        SOCKET_ONOPEN(state, e) {
            Vue.prototype.$socket = e.currentTarget;
            state.socket.isConnected = true;
        },
        SOCKET_RECONNECT(state) {
            state.socket.isConnected = true;
        },
        SOCKET_RECONNECT_ERROR(state) {
            state.socket.isConnected = false;
        },
        SOCKET_ONMESSAGE(state, e) {
            // console.log('MESSAGE', e);
            addEntry(state, JSON.parse(e.data));
        },
        /*
        SOCKET_ONERROR(state, e) {
            console.log(e);
        },
        */
        SOCKET_ONCLOSE(state) {
            state.socket.isConnected = false;
        },
    },
    actions: {
        addEvent({commit}, e) {
            commit('ADD_EVENT', e);
        },
        queryKey({commit}, e) {
            commit('ADD_EVENT', e);
        },
        reset({commit}) {
            commit('CLEAR');
        }
    },
    getters: {
        EVENTS: state => state.events
    }
});