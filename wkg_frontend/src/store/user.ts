import { defineStore } from 'pinia'
// import { menuList } from "../common/menu";

export const useMainStore = defineStore({
    id:'mainStore',
    state: () => {
        return {
            menu: [],
            token: "",
            role: ""
        }
    },
    persist: true

})
