<template>
    <div>
        <v-list two-line>
            <template v-for="item in filteredItems">
                <v-list-tile avatar
                             :key="item._id">
                    <v-list-tile-avatar>
                        <v-avatar class="accent"
                                  :size="40">
                            <span class="white--text title">{{ item.code }}</span>
                        </v-avatar>
                    </v-list-tile-avatar>
                    <v-list-tile-content>
                        <v-list-tile-title>{{item.code}}</v-list-tile-title>
                        <v-list-tile-sub-title>{{item.description}}</v-list-tile-sub-title>
                    </v-list-tile-content>
                    <v-list-tile-action>
                        <v-btn icon
                               @click.stop="viewCategory(item._id)">
                            <v-icon color="grey lighten-1">info_outline</v-icon>
                        </v-btn>
                    </v-list-tile-action>
                    <v-list-tile-action>
                        <v-btn icon
                               @click.stop="deleteCategory(item._id)">
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
                    {{ currentId == null? 'Add' : 'Edit'}} Category
                </v-card-title>
                <v-card-text>
                    <v-text-field v-model="model.code"
                                  label="Code"
                                  placeholder="Please describe"
                                  :error-messages="validations.code"></v-text-field>
                    <v-text-field v-model="model.description"
                                  label="Description"
                                  placeholder="Please describe"
                                  :error-messages="validations.description"></v-text-field>
                </v-card-text>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="primary"
                           flat
                           @click.stop="dialog=false">Close</v-btn>
                    <v-btn color="primary"
                           flat
                           @click.native.stop="save">Save</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
    </div>
</template>

<script>
import commonMethods from '@/mixins/common_methods'
import validator from '@/mixins/validator'

export default {
    mixins: [commonMethods, validator],
    props: {
        searchKey: {
            type: String,
            default: null
        }
    },
    data() {
        return {
            currentId: null,
            dialog: false,
            list: [],
            model: {
                code: null,
                description: null
            }
        }
    },
    computed: {
        filteredItems() {
            let vm = this
            if (!vm.searchKey) {
                return vm.list
            }

            return vm.list.filter(obj_category => {
                return (
                    vm.isStringContain(obj_category.code, vm.searchKey) ||
                    vm.isStringContain(obj_category.description, vm.searchKey)
                )
            })
        }
    },
    methods: {
        modelReset() {
            this.model = {
                code: null,
                description: null
            }
        },
        openDialog() {
            this.currentId = null
            this.modelReset()
            this.dialog = true
        },
        viewCategory(str_id) {
            let vm = this
            vm.resetValidation() // mixin
            vm.currentId = str_id

            vm.axios
                .get(`/settings/categories/${str_id}`)
                .then(obj_response => {
                    vm.model = obj_response.data
                    vm.dialog = true
                })
                .catch(obj_exception => {})
        },
        addCategory() {
            let vm = this

            vm
                .post('/settings/categories', vm.model)
                .then(obj_response => {
                    vm.dialog = false
                    this.$store.dispatch('addToast', obj_response.data)
                    vm.load()
                })
                .catch(obj_exception => {})
        },
        editCategory() {
            let vm = this

            vm
                .put(`/settings/categories/${vm.currentId}`, vm.model)
                .then(obj_response => {
                    vm.dialog = false
                    this.$store.dispatch('addToast', obj_response.data)
                    vm.load()
                })
                .catch(obj_exception => {})
        },
        deleteCategory(str_id) {
            let vm = this

            vm.axios
                .delete(`/settings/categories/${str_id}`)
                .then(obj_response => {
                    this.$store.dispatch('addToast', obj_response.data)
                    vm.load()
                })
                .catch(obj_exception => {})
        },
        save() {
            if (this.currentId) {
                this.editCategory()
            } else {
                this.addCategory()
            }
            this.load()
        },
        load() {
            this.axios
                .get('/settings/categories')
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
