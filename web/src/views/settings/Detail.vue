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
        <empty-state v-if="currentRouteName=='settings'"></empty-state>
        <v-layout v-else>
            <v-flex xs12
                    lg8
                    offset-lg2>
                <router-view :searchKey="searchKey"></router-view>
            </v-flex>
        </v-layout>
    </div>
</template>

<script>
import EmptyState from './Empty.vue'
export default {
    components: {
        EmptyState
    },
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
            switch (this.currentRouteName) {
                case 'settings':
                case 'settings.cost':
                case 'settings.empty':
                    return false
                default:
                    return true
                    break
            }
        }
    }
}
</script>

<style>

</style>
