<template>
    <!-- :timeout="toastTimeout" -->
    <v-snackbar :color="toast.type"
                :bottom="true"
                :right="true"
                auto-height
                v-model="show">
        {{ toast.message }}
        <v-btn flat
               class="white--text"
               @click.native="dismissToast">Close</v-btn>
    </v-snackbar>
</template>

<script>
export default {
    computed: {
        toast() {
            return this.$store.getters.toast || {}
        },
        toastTimeout() {
            return this.$store.getters.toastTimeout
        },
        show() {
            return this.$store.getters.toast !== null
        }
    },
    methods: {
        dismissToast() {
            this.$store.dispatch('dismissToast')
        }
    },
    created() {
        this.$store.dispatch('startToaster')
    }
}
</script>

<style>

</style>
