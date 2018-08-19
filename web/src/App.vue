<template>
    <v-app v-resize="onResize">
        <div :class="appClasses">
            <router-view></router-view>
        </div>
        <toast></toast>
    </v-app>
</template>

<script>
import Toast from './components/Toast.vue'
export default {
    name: 'App',
    components: {
        Toast
    },
    computed: {
        appClasses() {
            return {
                inner_body: true,
                'py-3': this.$vuetify.breakpoint.xlOnly,
                mobile: this.$vuetify.breakpoint.smAndDown,
                tablet: this.$vuetify.breakpoint.lgAndDown
            }
        }
    },
    mounted() {
        this.onResize()
    }
}
</script>
<style>
html {
    overflow: hidden;
}
html .application {
    font-family: 'Noto Sans', 'Roboto', sans-serif;
}
.inner_body {
    height: 100vh;
    width: 100vw;
    background: #888;
    /* background: url('./assets/background.jpg')no-repeat center center fixed; */
}
.relative {
    position: relative;
}

.list__tile .list__tile__action > .avatar,
.list__tile .avatar > .avatar {
    min-width: auto;
    -webkit-box-pack: center;
    -ms-flex-pack: center;
    justify-content: center;
}

.list__tile .list__tile__action > .avatar.white,
.list__tile .avatar > .avatar.white {
    border-color: #ddd !important;
    border: 1px solid;
}

/* transition */
.list-enter-active,
.list-leave-active {
    transition: all 1s;
}
.list-enter,
.list-leave-to {
    opacity: 0;
    transform: translateY(30px);
}

/* slide x  */

.slide-x150-transition-enter,
.slide-x150-transition-leave-to {
    opacity: 0;
    -webkit-transform: translateX(-150px);
    transform: translateX(-150px);
}
/*
.application--wrap .inner_body.tablet  {
    
}
*/
</style>

