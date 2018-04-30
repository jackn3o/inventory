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
            <v-text-field v-model="model.color"
                          label="Color"
                          placeholder="Please describe"
                          :error-messages="validations.color"></v-text-field>
            <v-text-field v-model="model.openingBalance"
                          type="number"
                          label="Opening Balance"
                          placeholder="Please describe"
                          :error-messages="validations.openingBalance"></v-text-field>
        </v-card-text>
    </v-card>
</template>

<script>
import validator from '../mixins/validator'
export default {
    mixins: [validator],
    data() {
        return {
            model: {
                code: null,
                description: null,
                color: null,
                openingBalance: null
            }
        }
    },
    methods: {
        addNewInventory() {
            let vm = this

            vm
                .post('/setting/items', this.model)
                .then(obj_response => {
                    vm.$emit('close')
                })
                .catch(obj_exception => {
                    console.log(obj_exception)
                })
        }
    }
}
</script>

<style></style>
