<template>
    <v-layout row
              fill-height>
        <v-flex xl10
                offset-xl1
                xs12>
            <v-card height="100%"
                    class="relative elevation-8">
                <v-progress-linear v-model="progress"
                                   :active="progress>0"
                                   indeterminate
                                   height="4"
                                   color="primary"
                                   style="position: absolute; top:0; left:0; margin:0; z-index:6"></v-progress-linear>
                <left-navigation v-model="showLeftNav"></left-navigation>
                <bottom-navigation v-model="showBottomNav"></bottom-navigation>
                <router-view :class="routerViewClasses"></router-view>
            </v-card>
        </v-flex>
    </v-layout>
</template>

<script>
import LeftNavigation from '@/views/LeftNavigation.vue'
import BottomNavigation from '@/views/BottomNavigation.vue'

export default {
    components: {
        LeftNavigation,
        BottomNavigation
    },
    data() {
        return {}
    },
    computed: {
        progress() {
            return this.$store.getters.progress
        },
        currentUser() {
            return this.$store.getters.currentUser
        },
        showLeftNav() {
            return this.$vuetify.breakpoint.mdAndUp
        },
        showBottomNav() {
            return this.$vuetify.breakpoint.smAndDown
        },
        routerViewClasses() {
            return {
                main_router_view: true,
                bottom_navigation: this.$vuetify.breakpoint.smAndDown
            }
        }
    }
}
</script>

<style lang="scss">
.main_router_view {
    position: absolute;
    top: 0;
    left: 60px;
    width: calc(100% - 60px);
    height: 100%;
    z-index: 4;
    &.bottom_navigation {
        left: 0;
        width: 100%;
        height: calc(100% - 56px);
    }
}
.navigation_drawer {
    .v-list__tile {
        padding: 0 18px;
    }
    &.v-navigation-drawer--mini-variant .v-list__tile {
        padding: 0 6px;
    }
}
.tablet {
    .v-toolbar__content,
    .v-toolbar__extension {
        padding: 0 12px;
    }
    .v-input {
        height: 48px;
        font-size: 14px;
        .v-input__prepend-outer {
            margin-top: 8px;
            margin-right: 6px;
        }
        .v-input__control {
            min-height: 40px;
        }
        .v-text-field__slot {
            height: 40px;
        }
    }
}
</style>
