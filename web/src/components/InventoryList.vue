<template>
    <v-layout fill-height>
        <v-flex style="overflow:hidden;"
                :class="sideMenuClasses"
                :hidden-sm-and-down="$route.name !== 'inventory.list'">
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
                    <!-- <v-list-group v-for="category in categories"
                                  v-model="category.active"
                                  :key="category.description"
                                  class="inventory_list_group">
                        <v-list-tile slot="activator">
                            <v-list-tile-content>
                                <v-list-tile-title>
                                    <div class="group_title">
                                        {{ category.description }}
                                        <div class="category_count_badge">{{ list.filter(obj_item=> { return obj_item.category == category._id}).length || 0 }}</div>
                                    </div>
                                </v-list-tile-title>
                            </v-list-tile-content>
                        </v-list-tile> -->
                    <template v-for="category in categories">
                        <v-subheader :key="category._id">
                            {{ category.description }}
                            <v-spacer></v-spacer>
                            <div class="category_count_badge">{{ list.filter(obj_item=> { return obj_item.category == category._id}).length || 0 }}</div>
                        </v-subheader>
                        <v-list-tile v-for="item in list.filter(obj_item=> { return obj_item.category == category._id})"
                                     :key="item.code"
                                     avatar
                                     ripple
                                     :value="item._id == currentId"
                                     :class="[item._id == currentId? 'grey lighten-4':'']"
                                     @click="select(item._id)">
                            <v-list-tile-avatar>
                                <v-avatar size="34px"
                                          :style="{ background: getHex(item.color), border: '1px solid #cccccc'}">
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
                    </template>
                    <!-- </v-list-group> -->
                </v-list>
            </v-card>
            <transition name="slide-x-transition"
                        enter-class="slide-x150-transition-enter"
                        leave-to-class="slide-x150-transition-leave-to">
                <add-inventory v-if="isAddNew"
                               @close="close()"></add-inventory>
            </transition>
        </v-flex>
        <v-flex :class="contentClasses">
            <transition name="slide-x-reverse-transition"
                        mode="in-out">
                <!-- <transition :name="$route.name == 'inventory:list'? 'slide-x-reverse-transition':'slide-y-reverse-transition'"> -->
                <router-view style="z-index:6; overflow:hidden; height:100%; border-left: 1px solid rgba(0,0,0,0.12)">
                </router-view>
            </transition>
        </v-flex>
    </v-layout>
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
            categories: null,
            colors: null,
            list: [],
            isShow2ndLevel: false, // for computed,
            isAddNew: false
        }
    },
    computed: {
        currentId() {
            return this.$route.params.id || ''
        },

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
        sideMenuClasses() {
            return {
                sm12: this.$route.name == 'inventory.list',
                sm3: this.$route.name !== 'inventory.list'
            }
        },
        contentClasses() {
            return {
                xs0: this.$route.name == 'inventory.list',
                sm9: this.$route.name !== 'inventory.list'
            }
        },
        title() {
            return 'Inventory List'
        }
    },
    methods: {
        load() {
            this.axios
                .get('/settings/categories')
                .then(obj_response => {
                    if (!obj_response || !obj_response.data) {
                        return
                    }

                    this.categories = obj_response.data
                    // .map(obj_data => {
                    //     return { description: obj_data.description, active: false }
                    // })
                })
                .catch(obj_exception => {})

            this.axios
                .get('/settings/colors')
                .then(obj_response => {
                    if (!obj_response || !obj_response.data) {
                        return
                    }

                    this.colors = obj_response.data
                    // .map(obj_data => {
                    //     return { description: obj_data.description, active: false }
                    // })
                })
                .catch(obj_exception => {})

            this.axios
                .get('/items')
                .then(obj_response => {
                    this.list = obj_response.data || []
                    this.setActiveCategories()
                })
                .catch(obj_exception => {})

            // let vm = this,
            //     MockData = require('./../data.js')

            // MockData.getInventory().then(obj_response => {
            //     vm.list = obj_response
            // })
        },
        close() {
            this.load()
            this.isAddNew = false
        },
        select(str_id) {
            let vm = this

            this.$router.push({ name: 'inventory.detail ', params: { id: str_id } })
        },
        getHex(str_colorId) {
            let vm = this
            if (!vm.colors || !str_colorId) {
                return '#90A4AE'
            }

            let str_hex = vm.colors.find(obj_color => {
                return obj_color._id == str_colorId
            }).hex

            return str_hex
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
        },
        setActiveCategories() {
            let vm = this
            if (!vm.list || vm.currentId) {
                return ''
            }

            let obj_result = vm.list.find(obj_item => {
                obj_item.category == vm.currentId
            })
            if (obj_result) {
                vm.categories.find(obj_category => {
                    obj_category._id == vm.currentId
                })['active'] = true
            }
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
