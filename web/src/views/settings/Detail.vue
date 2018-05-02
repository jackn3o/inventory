<template>
    <div>
        <v-toolbar color="grey lighten-4"
                   flat
                   prominent
                   extended>
            <v-btn icon
                   class="hidden-md-and-up"
                   @click.native="$router.push({name:'inventory.list'})">
                <v-icon>keyboard_backspace</v-icon>
            </v-btn>
            <v-toolbar-title class="hidden-md-and-up">{{ $route.meta.title }}</v-toolbar-title>
            <v-spacer></v-spacer>
            <v-layout slot="extension">
                <v-flex xs12
                        lg8
                        offset-lg2>
                    <v-text-field v-model="searchKey"
                                  prepend-icon="search"
                                  label="search"
                                  solo-inverted
                                  flat></v-text-field>
                </v-flex>
            </v-layout>
        </v-toolbar>
        <v-divider></v-divider>
        <v-layout>
            <v-flex xs12
                    lg8
                    offset-lg2>
                <router-view></router-view>
            </v-flex>
        </v-layout>
    </div>
</template>

<script>
export default {
    data() {
        return {
            searchKey: null
        }
    }
}
</script>

<style>

</style>
