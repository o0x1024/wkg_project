import { defineStore } from 'pinia'
// import { menuList } from "../common/menu";

export const useSearcgStore = defineStore({
    id: 'Search',
    state: () => {
        return {
            searchkey: ""
        }
    },
    persist: true

})
