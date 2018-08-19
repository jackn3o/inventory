<template>
    <v-bottom-nav :value="value"
                  :active.sync="bottomNav"
                  absolute
                  shift
                  color="black"
                  style="z-index: 9;">
        <v-btn dark
               flat
               value="inventory.list">
            <span>Inventory</span>
            <v-icon>home</v-icon>
        </v-btn>

        <v-btn dark
               flat
               value="settings">
            <span>Setting</span>
            <v-icon>settings</v-icon>
        </v-btn>

        <v-btn dark
               flat
               value="account">
            <span>Account</span>
            <v-icon>person</v-icon>
        </v-btn>
    </v-bottom-nav>
</template>

<script>
export default {
    props: {
        value: {
            type: Boolean,
            default: false
        }
    },
    data() {
        return {}
    },
    computed: {
        bottomNav: {
            get() {
                return this.$route.name
            },
            set(value) {
                this.$router.push({ name: value })
            }
        }
    }
}
</script>

<style>
</style>
