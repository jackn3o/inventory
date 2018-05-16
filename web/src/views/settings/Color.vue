<template>
    <div>
        <v-list two-line>
            <template v-for="item in filteredItems">
                <v-list-tile avatar
                             dense
                             :key="item._id">
                    <v-list-tile-avatar>
                        <div style="height:34px; width:34px; border:2px solid #ccc; border-radius:12px;"
                             :style="{ background: item.hex}">
                        </div>
                    </v-list-tile-avatar>
                    <v-list-tile-content>
                        <v-list-tile-sub-title>{{item.description}}</v-list-tile-sub-title>
                    </v-list-tile-content>
                    <v-list-tile-action>
                        <v-btn icon
                               @click.native.stop="viewColor(item._id)">
                            <v-icon color="grey lighten-1">info_outline</v-icon>
                        </v-btn>
                    </v-list-tile-action>
                    <v-list-tile-action>
                        <v-btn icon
                               @click.native.stop="deleteColor(item._id)">
                            <v-icon color="red lighten-1">delete</v-icon>
                        </v-btn>
                    </v-list-tile-action>
                </v-list-tile>
                <v-divider :key="'d-'+item._id"></v-divider>
            </template>
        </v-list>
        <v-fab-transition>
            <v-btn v-if="!dialog"
                   color="pink"
                   dark
                   absolute
                   fab
                   style="bottom: 16px; right: 16px;"
                   @click.native="openDialog">
                <v-icon>add</v-icon>
            </v-btn>
        </v-fab-transition>
        <v-dialog v-model="dialog"
                  max-width="500px">
            <v-card class="pa-3">
                <v-card-title class="title">
                    {{ currentId == null? 'Add' : 'Edit'}} Color
                </v-card-title>
                <v-card-text>
                    <v-text-field v-model="model.description"
                                  label="Description"
                                  placeholder="Please describe"
                                  :error-messages="validations.description"></v-text-field>
                </v-card-text>
                <v-card-text class="grey lighten-4"
                             style="border-radius: 12px;">
                    <div class="body-1 grey--text text--darken-1 ml-1">
                        Color
                        <span v-if="validations.hex.length >0"
                              class="red--text">*</span>
                    </div>
                    <swatches v-model="model.hex"
                              background-color="#F5F5F5"
                              :colors="colors"
                              :swatch-style="swatchStyle"
                              :wrapper-style="wrapperStyle"
                              inline></swatches>
                </v-card-text>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="primary"
                           flat
                           @click.stop="dialog=false">Close</v-btn>
                    <v-btn color="primary"
                           flat
                           @click.stop="save">Save</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
    </div>
</template>

<script>
import Swatches from 'vue-swatches'
// Import the styles too, globally
import 'vue-swatches/dist/vue-swatches.min.css'
import commonMethods from '@/mixins/common_methods'
import validator from '../../mixins/validator'
export default {
    mixins: [commonMethods, validator],
    components: {
        Swatches
    },
    props: {
        searchKey: {
            type: String,
            default: null
        }
    },
    data() {
        return {
            currentId: null,
            list: [],
            dialog: false,
            model: {
                description: null,
                hex: null
            },
            colors: [
                '#F44336', //red
                '#E91E63', //pink
                '#9C27B0', //purple
                '#3F51B5', //indigo
                '#2196F3', //blue
                '#00BCD4', //cyan
                '#009688', //teal
                '#4CAF50', //green
                '#8BC34A', //light-green
                '#CDDC39', //lime
                '#FFEB3B', //yellow
                '#FFFDE7', //yellow light
                '#FF9800', //orange
                '#795548', //brown
                '#607D8B', //blue-grey
                '#9E9E9E', //grey
                '#000000', //black
                '#ffffff' //white
            ],
            swatchStyle: {
                border: '2px solid #eee'
            },
            wrapperStyle: {
                width: '100%'
            }
        }
    },
    computed: {
        filteredItems() {
            let vm = this
            if (!vm.searchKey) {
                return vm.list
            }

            return vm.list.filter(obj_color => {
                return vm.isStringContain(obj_color.description, vm.searchKey)
            })
        }
    },
    methods: {
        modelReset() {
            this.model = {
                description: null,
                hex: null
            }
        },
        openDialog() {
            this.currentId = null
            this.modelReset()
            this.dialog = true
        },
        viewColor(str_id) {
            let vm = this
            vm.resetValidation() // mixin
            vm.currentId = str_id

            vm.axios
                .get(`/settings/colors/${str_id}`)
                .then(obj_response => {
                    vm.model = obj_response.data
                    vm.dialog = true
                })
                .catch(obj_exception => {})
        },
        deleteColor(str_id) {
            let vm = this

            vm.axios
                .delete(`/settings/colors/${str_id}`)
                .then(obj_response => {
                    this.$store.dispatch('addToast', obj_response.data)
                    vm.load()
                })
                .catch(obj_exception => {})
        },
        save() {
            if (this.currentId) {
                this.editColor()
            } else {
                this.addColor()
            }
            this.load()
        },
        addColor() {
            let vm = this
            vm
                .post('/settings/colors', vm.model)
                .then(obj_response => {
                    vm.dialog = false
                    this.$store.dispatch('addToast', obj_response.data)
                    vm.load()
                })
                .catch(obj_exception => {})
        },
        editColor() {
            let vm = this

            vm
                .put(`/settings/colors/${vm.currentId}`, vm.model)
                .then(obj_response => {
                    vm.dialog = false
                    this.$store.dispatch('addToast', obj_response.data)
                    vm.load()
                })
                .catch(obj_exception => {
                    vm.dialog = true
                })
        },

        load() {
            this.axios
                .get('/settings/colors')
                .then(obj_response => {
                    this.list = obj_response.data
                })
                .catch(obj_exception => {})
        }
    },
    mounted() {
        this.load()
    }
}
</script>

<style>

</style>
