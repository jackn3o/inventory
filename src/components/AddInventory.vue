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
                   @click="addNewInventory">
                <v-icon>done</v-icon>
            </v-btn>
            <template slot="extension">
                <v-toolbar-title>Add new item</v-toolbar-title>
            </template>
        </v-toolbar>
        <v-card-text>
            <v-text-field v-model="code"
                          label="Code"
                          placeholder="Please describe"></v-text-field>
            <v-text-field v-model="description"
                          label="Description"
                          placeholder="Please describe"></v-text-field>
            <v-text-field v-model="color"
                          label="Color"
                          placeholder="Please describe"></v-text-field>
            <v-text-field v-model="openBalance"
                          label="Opening Balance"
                          placeholder="Please describe"></v-text-field>
        </v-card-text>
    </v-card>
</template>

<script>
export default {
    data() {
        return {
            code: null,
            description: null,
            color: null,
            openBalance: null
        }
    },
    methods: {
        addNewInventory() {
            let vm = this,
                MockData = require('./../data.js'),
                nb_itemCount = MockData.items.length,
                obj_new = {
                    value: false,
                    id: '000' + String(nb_itemCount + 1),
                    code: vm.code,
                    description: vm.description,
                    color: vm.color,
                    balance: vm.openBalance
                }

            MockData.addNewItem(obj_new).then(obj_response => {
                vm.$emit('close')
            })
        }
    }
}
</script>

<style>

</style>
