<template>
    <v-app v-resize="onResize">
        <div class="inner_body"
             :class="appClasses">
            <v-layout row
                      fill-height>
                <v-flex xl10
                        offset-xl1
                        xs12>
                    <v-card height="100%"
                            class="relative elevation-8">
                        <v-navigation-drawer dark
                                             permenant
                                             absolute
                                             hide-overlay
                                             :mini-variant="mini"
                                             v-model="drawer"
                                             style="z-index:5;">
                            <!-- <v-toolbar flat
                                       class="transparent">
                                <v-list class="pa-0">

                                </v-list>
                            </v-toolbar> -->
                            <v-list class="pt-0">
                                <v-list-tile>
                                    <v-list-tile-action>
                                        <v-btn icon
                                               @click.native.stop="mini = !mini">
                                            <v-icon>menu</v-icon>
                                        </v-btn>
                                    </v-list-tile-action>
                                    <v-list-tile-content>
                                        <v-list-tile-title>Menu</v-list-tile-title>
                                    </v-list-tile-content>
                                    <v-list-tile-action>
                                        <v-btn icon
                                               @click.native.stop="mini = !mini">
                                            <v-icon>chevron_left</v-icon>
                                        </v-btn>
                                    </v-list-tile-action>
                                </v-list-tile>
                                <v-divider></v-divider>
                                <v-list-tile avatar>
                                    <v-list-tile-avatar>
                                        <img src="https://randomuser.me/api/portraits/men/85.jpg">
                                    </v-list-tile-avatar>
                                    <v-list-tile-content>
                                        <v-list-tile-title>John Leider</v-list-tile-title>
                                    </v-list-tile-content>
                                </v-list-tile>
                                <v-list-tile>
                                    <v-list-tile-action>
                                        <v-btn icon>
                                            <v-icon>storage</v-icon>
                                        </v-btn>
                                    </v-list-tile-action>
                                    <v-list-tile-content>
                                        <v-list-tile-title>Inventory</v-list-tile-title>
                                    </v-list-tile-content>
                                </v-list-tile>
                                <v-list-tile>
                                    <v-list-tile-action>
                                        <v-btn icon>
                                            <v-icon>settings</v-icon>
                                        </v-btn>
                                    </v-list-tile-action>
                                    <v-list-tile-content>
                                        <v-list-tile-title>Setting</v-list-tile-title>
                                    </v-list-tile-content>
                                </v-list-tile>
                            </v-list>
                        </v-navigation-drawer>
                        <router-view style="position:absolute; top:0; left:80px; width:calc(100% - 80px); height: 100%; z-index:4;"></router-view>
                    </v-card>
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
            title: 'Vuetify.js',
            mini: true,
            drawer: true
        }
    },
    computed: {
        currentViewportSize() {
            return this.$store.getters.currentViewportSize
        },
        appClasses() {
            return {
                'py-3': this.currentViewportSize === 'xl'
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
</style>

