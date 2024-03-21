import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import store from "./store";
import router from "./router";
import Antd from "ant-design-vue";
import "ant-design-vue/dist/antd.css";
import * as Icons from '@ant-design/icons-vue'

// 注册

const app = createApp(App);

Object.keys(Icons).map(key => {
    app.component(key, Icons[key as keyof typeof Icons])
})



app.use(store);
app.use(router);
app.use(Antd);
app.mount("#app");
