<template>
    <v-card tile
            height="100%"
            style="z-index:5; position:absolute; left:0;top:0;width:100%;">
        <v-toolbar color="cyan"
                   flat
                   dark
                   prominent
                   extended>
            <v-btn icon
                   @click="$emit('close')">
                <v-icon>keyboard_backspace</v-icon>
            </v-btn>
            <v-spacer></v-spacer>
            <v-btn icon
                   @click.native="addNewInventory">
                <v-icon>done</v-icon>
            </v-btn>
            <template slot="extension">
                <v-toolbar-title>Add new item</v-toolbar-title>
            </template>
        </v-toolbar>
        <v-card-text>
            <v-text-field v-model="model.code"
                          label="Code"
                          placeholder="Please describe"
                          :error-messages="validations.code"></v-text-field>
            <v-text-field v-model="model.description"
                          label="Description"
                          placeholder="Please describe"
                          :error-messages="validations.description"></v-text-field>
            <v-text-field v-model="model.openingBalance"
                          type="number"
                          label="Opening Balance"
                          placeholder="Please describe"
                          :error-messages="validations.openingBalance"></v-text-field>
            <v-select v-model="model.color"
                      label="Color"
                      placeholder="Please select"
                      autocomplete
                      :items="colors"
                      item-value="_id"
                      item-text="description"></v-select>
            <v-select v-model="model.category"
                      label="Category"
                      placeholder="Please select"
                      autocomplete
                      :items="categories"
                      item-value="_id"
                      item-text="description"></v-select>
            <v-select v-model="model.outlet"
                      label="Outlet"
                      placeholder="Please select"
                      autocomplete
                      :items="outlets"
                      item-value="_id"
                      item-text="description"></v-select>
        </v-card-text>
    </v-card>
</template>

<script>
import validator from '../mixins/validator'
export default {
    mixins: [validator],
    data() {
        return {
            colors: null,
            categories: null,
            outlets: null,
            model: {
                code: null,
                description: null,
                category: null,
                outlet: null,
                color: null,
                openingBalance: null
            }
        }
    },
    methods: {
        addNewInventory() {
            let vm = this

            vm
                .post('/settings/items', this.model)
                .then(obj_response => {
                    vm.$emit('close')
                })
                .catch(obj_exception => {
                    console.log(obj_exception)
                })
        },
        load() {
            this.axios
                .get('/settings/categories')
                .then(obj_response => {
                    this.categories = obj_response.data
                })
                .catch(obj_exception => {})

            this.axios
                .get('/settings/outlets')
                .then(obj_response => {
                    this.outlets = obj_response.data
                })
                .catch(obj_exception => {})

            this.axios
                .get('/settings/colors')
                .then(obj_response => {
                    this.colors = obj_response.data
                })
                .catch(obj_exception => {})
        }
    },
    mounted() {
        this.load()
    }
}
</script>

<style></style>
