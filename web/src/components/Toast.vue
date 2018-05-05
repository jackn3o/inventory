<template>
    <v-snackbar :timeout="toastTimeout"
                :color="toast.type"
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
    data() {
        return {
            test: false
        }
    },
    computed: {
        toast() {
            return this.$store.getters.toast || {}
        },
        toastTimeout() {
            return this.$store.getters.toastTimeout
        },
        show() {
            if (!this.$store.getters.toastQueue && !this.toast) {
                return false
            }

            this.showToast()
            return true
        }
    },
    methods: {
        showToast() {
            this.$store.dispatch('showToast')
        },
        dismissToast() {
            this.$store.dispatch('dismissToast')
        }
    }
}
</script>

<style>

</style>
