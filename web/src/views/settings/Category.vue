<template>
    <div>
        <v-list two-line>
            <template v-for="item in list">
                <v-list-tile avatar
                             :key="item.id">
                    <v-list-tile-avatar>
                        <v-avatar class="red">
                            <span class="white--text headline">{{ item.code }}</span>
                        </v-avatar>
                    </v-list-tile-avatar>
                    <v-list-tile-content>
                        <v-list-tile-title>{{item.code}}</v-list-tile-title>
                        <v-list-tile-sub-title>{{item.description}}</v-list-tile-sub-title>
                    </v-list-tile-content>
                    <v-list-tile-action>
                        <v-btn icon>
                            <v-icon>info_outline</v-icon>
                        </v-btn>
                    </v-list-tile-action>
                    <v-list-tile-action>
                        <v-btn icon>
                            <v-icon>delete</v-icon>
                        </v-btn>
                    </v-list-tile-action>
                </v-list-tile>
                <v-divider :key="'d-'+item.id"></v-divider>
            </template>
        </v-list>
        <v-btn v-if="!dialog"
               color="pink"
               dark
               absolute
               fab
               style="bottom: 16px; right: 16px;"
               @click.native="dialog=true">
            <v-icon>add</v-icon>
        </v-btn>
        <v-dialog v-model="dialog"
                  max-width="500px">
            <v-card class="pa-3">
                <v-card-title class="title">
                    Add Category
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
                           @click.stop="addCategory">Save</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
    </div>
</template>

<script>
import validator from '../../mixins/validator'
export default {
    mixins: [validator],
    data() {
        return {
            title: '123',
            list: [],
            dialog: false,
            model: {
                code: null,
                description: null
            }
        }
    },
    methods: {
        addCategory() {
            let vm = this

            vm
                .post('/settings/categories', vm.model)
                .then(obj_response => {
                    vm.dialog = false
                    vm.load()
                })
                .catch(obj_exception => {
                })
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
