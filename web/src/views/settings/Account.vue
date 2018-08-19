<template>
    <v-layout row
              justify-center>
        <v-flex sm12
                md10
                lg6>
            <v-card flat>
                <v-card-title class="headline">Reset Password</v-card-title>
                <v-card-text>
                    <v-text-field v-model="model.oldPassword"
                                  type="password"
                                  label="Old Password"
                                  placeholder="Old Password"
                                  :error-messages="validations.oldPassword"></v-text-field>
                    <v-text-field v-model="model.newPassword"
                                  type="password"
                                  label="New Password"
                                  placeholder="New Password"
                                  :error-messages="validations.newPassword"></v-text-field>
                    <v-text-field v-model="model.confirmPassword"
                                  type="password"
                                  label="Confirm Password"
                                  placeholder="Confirm Password"
                                  :error-messages="validations.confirmPassword"></v-text-field>
                </v-card-text>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn flat="flat"
                           @click.native="dialog = false">No</v-btn>
                    <v-btn color="primary"
                           flat="flat"
                           @click.native="ResetPassword">Yes</v-btn>
                </v-card-actions>
                <v-divider></v-divider>
                <v-card-title class="headline">Logout</v-card-title>
                <v-card-text>Confirm to logout?</v-card-text>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn flat="flat"
                           @click.native="dialog = false">No</v-btn>
                    <v-btn color="primary"
                           flat="flat"
                           @click.native="logout">Yes</v-btn>
                </v-card-actions>
            </v-card>
        </v-flex>
    </v-layout>
</template>

<script>
import validator from '@/mixins/validator'
export default {
    mixins: [validator],
    data() {
        return {
            model: {
                oldPassword: null,
                newPassword: null,
                confirmPassword: null
            }
        }
    },
    methods: {
        openDialog(str_view) {
            this.view = str_view
            this.dialog = true
        },
        ResetPassword() {
            let vm = this
            vm.put(`/users/${vm.currentUser}/password`, vm.model).then(obj_response => {
                if (!obj_response || !obj_response.data) {
                    return
                }
                this.$store.dispatch('logout')

                this.$store.dispatch('addToast', 'Please login with new password')
            })
        },
        logout() {
            this.$store.dispatch('logout')
            this.$router.push({ name: 'login' })
        }
    }
}
</script>

<style>
</style>
