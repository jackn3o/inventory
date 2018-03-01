<template>
    <v-app v-resize="onResize">
        <div class="inner_body"
             :class="appClasses">
            <v-layout row
                      fill-height>
                <v-flex lg10
                        offset-lg1
                        xs12>
                    <router-view></router-view>
                </v-flex>
            </v-layout>
        </div>
    </v-app>
</template>

<script>
export default {
    name: 'App',
    data() {
        return {
            title: 'Vuetify.js'
        }
    },
    computed: {
        currentViewportSize() {
            return this.$store.getters.currentViewportSize
        },
        appClasses() {
            return {
                'pt-3':
                    this.currentViewportSize !== 'xs' &&
                    this.currentViewportSize !== 'sm' &&
                    this.currentViewportSize !== 'md',
                'pb-3':
                    this.currentViewportSize !== 'xs' &&
                    this.currentViewportSize !== 'sm' &&
                    this.currentViewportSize !== 'md'
            }
        }
    },
    methods: {
        onResize() {
            this.$store.commit('setViewportSize', window.innerWidth)
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
.inner_body {
    height: 100vh;
    width: 100vw;
    background: url('./assets/background.jpg')no-repeat center center fixed;
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
</style>

