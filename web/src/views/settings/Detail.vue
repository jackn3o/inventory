<template>
    <div>
        <v-toolbar color="grey lighten-4"
                   flat
                   prominent
                   extended>
            <!-- <v-btn icon
                   class="hidden-md-and-up">
                <v-icon>keyboard_backspace</v-icon>
            </v-btn> -->
            <v-toolbar-title class="hidden-md-and-up">{{ $route.meta.title }}</v-toolbar-title>
            <v-spacer></v-spacer>
            <v-layout slot="extension">
                <v-flex xs12
                        lg8
                        offset-lg2>
                    <v-text-field v-if="isShowSearch"
                                  v-model="searchKey"
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
                <router-view :searchKey="searchKey"></router-view>
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
    },
    computed: {
        currentRouteName() {
            return this.$route.name
        },
        isShowSearch() {
            if (this.currentRouteName == 'settings.cost') {
                return false
            }
            return true
        }
    }
}
</script>

<style>

</style>
