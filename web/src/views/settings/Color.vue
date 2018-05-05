<template>
    <div>
        <v-list two-line>
            <template v-for="item in list">
                <v-list-tile avatar
                             :key="item._id">
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
                <v-divider :key="'d-'+item._id"></v-divider>
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
                    Add Color
                </v-card-title>
                <v-card-text>
                    <v-text-field v-model="model.description"
                                  label="Description"
                                  placeholder="Please describe"
                                  :error-messages="validations.description"></v-text-field>
                </v-card-text>
                <v-card-text class="grey lighten-4"
                             style="border-radius: 12px;">
                    <div class="body-1 grey--text text--darken-1 ml-1">Color</div>
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
                           @click.stop="addOutlet">Save</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
    </div>
</template>

<script>
import Swatches from 'vue-swatches'
// Import the styles too, globally
import 'vue-swatches/dist/vue-swatches.min.css'
import validator from '../../mixins/validator'
export default {
    mixins: [validator],
    components: {
        Swatches
    },
    data() {
        return {
            list: [],
            dialog: false,
            model: {
                description: null,
                hex: null
            },
            color: '#000000',
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
    methods: {
        addOutlet() {
            let vm = this

            vm
                .post('/settings/colors', vm.model)
                .then(obj_response => {
                    vm.dialog = false
                    vm.load()
                })
                .catch(obj_exception => {})
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
