<template>
    <v-layout fill-height>
        <v-flex style="overflow:hidden;"
                :class="sideMenuClasses"
                :hidden-sm-and-down="$route.name !== 'setting'">
            <v-card tile
                    height="100%">
                <v-toolbar color="grey lighten-4"
                           flat
                           prominent
                           extended>
                    <div slot="extension">
                        <v-toolbar-title>Settings</v-toolbar-title>
                    </div>
                </v-toolbar>
                <v-divider></v-divider>
                <v-list>
                    <v-list-tile avatar
                                 v-for="m in menus"
                                 :key="m.title"
                                 @click="$router.push({name: m.to})">
                        <v-list-tile-avatar>
                            <!-- <v-icon :class="[item.iconClass]">{{ item.icon }}</v-icon> -->
                        </v-list-tile-avatar>
                        <v-list-tile-content>
                            <v-list-tile-title>{{ m.title }}</v-list-tile-title>
                        </v-list-tile-content>
                    </v-list-tile>
                </v-list>
            </v-card>
        </v-flex>
        <v-flex :class="contentClasses">
            <transition name="slide-x-reverse-transition"
                        mode="in-out">
                <detail style="z-index:6; overflow:hidden; height:100%; border-left: 1px solid rgba(0,0,0,0.12)">
                </detail>
            </transition>
        </v-flex>
    </v-layout>
</template>

<script>
import Detail from './Detail.vue'
export default {
    components: {
        Detail
    },
    data() {
        return {
            menus: [
                {
                    title: 'Category',
                    to: 'settings.category'
                },
                {
                    title: 'Color',
                    to: 'settings.color'
                },
                // {
                //     title: 'Item',
                //     to: 'settings.item'
                // },
                {
                    title: 'Outlet',
                    to: 'settings.outlet'
                }
            ],
            isShow2ndLevel: false // for computed,
        }
    },
    computed: {
        currentViewportSize() {
            return this.$store.getters.currentViewportSize
        },
        isDualColomnLayout() {
            if (this.currentViewportSize == 'xs' || this.currentViewportSize == 'sm') {
                return false
            }
            return true
        },
        active: {
            get() {
                if (this.isDualColomnLayout) {
                    return true
                }

                return this.isShow2ndLevel
            },
            set(value) {
                this.isShow2ndLevel = value
            }
        },
        sideMenuClasses() {
            return {
                sm12: this.$route.name == 'setting',
                sm3: this.$route.name !== 'setting'
            }
        },
        contentClasses() {
            return {
                xs0: this.$route.name == 'setting',
                sm9: this.$route.name !== 'setting'
            }
        }
    }
}
</script>

<style>

</style>
