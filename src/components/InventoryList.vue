<template>
    <v-card tile
            height="100%"
            class="transparent">
        <v-layout fill-height>
            <v-flex style="position: relative; overflow:hidden;"
                    :class="contentClasses"
                    :hidden-sm-and-down="$route.name !== 'inventory:list'">
                <v-card tile
                        height="100%">
                    <v-toolbar color="grey lighten-4"
                               flat
                               prominent
                               extended>
                        <v-spacer></v-spacer>
                        <v-btn icon
                               @click="isAddNew=true">
                            <v-icon>add</v-icon>
                        </v-btn>
                        <v-btn icon>
                            <v-icon>more_vert</v-icon>
                        </v-btn>
                        <template slot="extension">
                            <v-text-field v-model="searchKey"
                                          prepend-icon="search"
                                          label="search"
                                          solo-inverted
                                          flat
                                          class="mx-3"></v-text-field>
                        </template>
                    </v-toolbar>
                    <v-divider></v-divider>
                    <v-list>
                        <v-list-group v-model="category.active"
                                      v-for="category in filteredItems"
                                      :key="category.description"
                                      class="inventory_list_group">
                            <!-- :prepend-icon="item.action" -->
                            <!-- <v-subheader inset>
                                <div>Item(s)</div>
                                <v-spacer></v-spacer>
                                <div>Balance</div>
                            </v-subheader> -->
                            <!-- <v-divider inset></v-divider> -->
                            <v-list-tile slot="activator">
                                <v-list-tile-content>
                                    <v-list-tile-title>
                                        <div class="group_title">
                                            {{ category.description }}
                                            <div class="category_count_badge">{{ category.items.length || 0 }}</div>
                                        </div>
                                    </v-list-tile-title>

                                </v-list-tile-content>
                            </v-list-tile>
                            <v-list-tile v-for="item in category.items"
                                         :key="item.code"
                                         avatar
                                         ripple
                                         @click="select(item.id)">
                                <v-list-tile-avatar>
                                    <v-avatar size="34px"
                                              :class="getAvatarClass(item.color)">
                                        <span class="subheading">{{ (item.code)? item.code.charAt(0) : '?' }}</span>
                                    </v-avatar>
                                </v-list-tile-avatar>
                                <v-list-tile-content>
                                    <v-list-tile-title class="body-1">
                                        {{ item.code }} - {{ item.description}}
                                    </v-list-tile-title>
                                    <v-list-tile-sub-title>
                                        {{ item.outlet }}
                                    </v-list-tile-sub-title>
                                </v-list-tile-content>
                                <v-list-tile-action>
                                    {{ item.balance }}
                                </v-list-tile-action>
                            </v-list-tile>
                        </v-list-group>
                    </v-list>
                </v-card>
                <transition name="slide-x-transition"
                            enter-class="slide-x150-transition-enter"
                            leave-to-class="slide-x150-transition-leave-to">
                    <add-inventory v-if="isAddNew"
                                   @close="close()"></add-inventory>
                </transition>
            </v-flex>
            <transition name="slide-x-reverse-transition"
                        mode="in-out">
                <!-- <transition :name="$route.name == 'inventory:list'? 'slide-x-reverse-transition':'slide-y-reverse-transition'"> -->
                <router-view></router-view>
            </transition>
        </v-layout>
    </v-card>
</template>

<script>
import AddInventory from './AddInventory.vue'
export default {
    components: {
        AddInventory
    },
    data() {
        return {
            searchKey: null,
            list: [],
            isShow2ndLevel: false, // for computed,
            isAddNew: false
        }
    },
    computed: {
        filteredItems() {
            let vm = this
            if (!vm.searchKey) {
                return vm.list
            }

            return vm.list.filter(obj_item => {
                return (
                    vm.isStringContain(obj_item.code, vm.searchKey) || vm.isStringContain(obj_item.code, vm.searchKey)
                )
            })
        },
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
        contentClasses() {
            return {
                sm3: this.$route.name == 'inventory:list',
                sm3: this.$route.name !== 'inventory:list'
            }
        },
        title() {
            return 'Inventory List'
        }
    },
    methods: {
        load() {
            let vm = this,
                MockData = require('./../data.js')

            MockData.getInventory().then(obj_response => {
                vm.list = obj_response
            })
        },
        close() {
            this.load()
            this.isAddNew = false
        },
        select(str_id) {
            let vm = this

            this.$router.push({ name: 'inventory:detail', params: { id: str_id } })
        },
        getAvatarClass(color) {
            if (!color) {
                return ['grey', 'lighten-2']
            }
            let str_color = color.toLowerCase(),
                obj_class = {
                    border: str_color == 'white',
                    'white--text': str_color !== 'white'
                }

            obj_class[str_color] = str_color

            return obj_class
        },
        isStringContain(value, key) {
            if (!value || !key) {
                return false
            }

            let str_value = String(value).toLowerCase(),
                str_key = String(key).toLowerCase()

            if (str_value.search(str_key) >= 0) {
                return true
            }

            return false
        }
    },
    mounted() {
        this.load()
    }
}
</script>

<style>
.inventory_list_group .list__group__header .list__tile {
    padding-right: 0;
}
.inventory_list_group .list__group__header {
    border-bottom: 1px solid rgba(0, 0, 0, 0.12);
    -webkit-transition: 0.3s cubic-bezier(0.25, 0.8, 0.5, 1);
    transition: 0.3s cubic-bezier(0.25, 0.8, 0.5, 1);
}
.inventory_list_group .list__group__header.list__group__header--active {
    border-bottom: 0;
}
.inventory_list_group .group_title {
    display: flex;
}
.category_count_badge {
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: 12px;
    background: #ccc;
    color: #888;
    padding: 2px;
    font-size: 10px;
    margin-left: 10px;
    width: 24px;
    height: 20px;
}
</style>
