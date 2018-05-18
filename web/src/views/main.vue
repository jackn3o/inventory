<template>
    <v-layout row
              fill-height>
        <v-flex xl10
                offset-xl1
                xs12>
            <v-card height="100%"
                    class="relative elevation-8">
                <v-progress-linear v-model="progress"
                                   :active="progress>0"
                                   indeterminate
                                   height="4"
                                   color="primary"
                                   style="position: absolute; top:0; left:0; margin:0;  z-index:6"></v-progress-linear>
                <v-navigation-drawer dark
                                     permenant
                                     absolute
                                     hide-overlay
                                     :mini-variant="mini"
                                     v-model="drawer"
                                     style="display:flex; flex-direction:column; z-index:5; padding-bottom:0;">
                    <v-list class="pt-0">
                        <v-list-tile>
                            <v-list-tile-action>
                                <v-btn icon
                                       @click.native.stop="mini = !mini">
                                    <v-icon>menu</v-icon>
                                </v-btn>
                            </v-list-tile-action>
                            <v-list-tile-content>
                                <v-list-tile-title>Menu</v-list-tile-title>
                            </v-list-tile-content>
                            <v-list-tile-action>
                                <v-btn icon
                                       @click.native.stop="mini = !mini">
                                    <v-icon>chevron_left</v-icon>
                                </v-btn>
                            </v-list-tile-action>
                        </v-list-tile>
                        <v-divider></v-divider>

                        <v-list-tile>
                            <v-list-tile-action>
                                <v-btn icon
                                       @click.native="$router.push({name: 'inventory.list'})">
                                    <v-icon>home</v-icon>
                                </v-btn>
                            </v-list-tile-action>
                            <v-list-tile-content>
                                <v-list-tile-title>Inventory</v-list-tile-title>
                            </v-list-tile-content>
                        </v-list-tile>
                        <v-list-tile>
                            <v-list-tile-action>
                                <v-btn icon
                                       @click.native="$router.push({name: 'settings'})">
                                    <v-icon>settings</v-icon>
                                </v-btn>
                            </v-list-tile-action>
                            <v-list-tile-content>
                                <v-list-tile-title>Setting</v-list-tile-title>
                            </v-list-tile-content>
                        </v-list-tile>
                    </v-list>
                    <v-spacer></v-spacer>
                    <v-list>
                        <v-menu :nudge-width="200"
                                v-model="menu"
                                :nudge-left="60"
                                offset-x>
                            <v-list-tile avatar
                                         slot="activator">
                                <v-list-tile-avatar>
                                    <v-avatar size="40px"
                                              color="teal">
                                        <v-icon color="white">person</v-icon>
                                    </v-avatar>
                                </v-list-tile-avatar>
                                <v-list-tile-content>
                                    <v-list-tile-title>Admin</v-list-tile-title>
                                </v-list-tile-content>
                            </v-list-tile>
                            <v-card>
                                <v-card-title>
                                    <v-avatar size="40px"
                                              color="teal">
                                        <v-icon color="white">person</v-icon>
                                    </v-avatar>
                                    <div class="subheading ml-3">
                                        {{currentUser}}
                                    </div>
                                </v-card-title>
                                <v-divider></v-divider>
                                <v-list dense>
                                    <v-list-tile avatar
                                                 @click="openDialog('password')">
                                        <v-list-tile-avatar>
                                            <v-icon>vpn_key</v-icon>
                                        </v-list-tile-avatar>
                                        <v-list-tile-content>
                                            <v-list-tile-title>Change password</v-list-tile-title>
                                        </v-list-tile-content>
                                    </v-list-tile>
                                    <v-list-tile avatar
                                                 @click="openDialog('logout')">
                                        <v-list-tile-avatar>
                                            <v-icon>exit_to_app</v-icon>
                                        </v-list-tile-avatar>
                                        <v-list-tile-content>
                                            <v-list-tile-title>Logout</v-list-tile-title>
                                        </v-list-tile-content>
                                    </v-list-tile>
                                </v-list>
                            </v-card>
                        </v-menu>
                    </v-list>
                </v-navigation-drawer>
                <v-dialog v-model="dialog"
                          max-width="400px">
                    <v-card v-if="view=='password'">
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
                    </v-card>
                    <v-card v-if="view=='logout'">
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
                </v-dialog>
                <router-view style="position:absolute; top:0; left:80px; width:calc(100% - 80px); height: 100%; z-index:4;"></router-view>
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
            mini: true,
            drawer: true,
            dialog: false,
            menu: false,
            view: null,
            model: {
                oldPassword: null,
                newPassword: null,
                confirmPassword: null
            }
        }
    },
    computed: {
        progress() {
            return this.$store.getters.progress
        },
        currentUser() {
            return this.$store.getters.currentUser
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
