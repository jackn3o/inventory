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
                                    <v-list-tile avatar>
                                        <v-list-tile-avatar>
                                            <v-icon>vpn_key</v-icon>
                                        </v-list-tile-avatar>
                                        <v-list-tile-content>
                                            <v-list-tile-title>Change password</v-list-tile-title>
                                            <v-list-tile-sub-title></v-list-tile-sub-title>
                                        </v-list-tile-content>
                                        <v-list-tile-action>
                                        </v-list-tile-action>
                                    </v-list-tile>
                                    <v-list-tile avatar
                                                 @click="dialog=true">
                                        <v-list-tile-avatar>
                                            <v-icon>exit_to_app</v-icon>
                                        </v-list-tile-avatar>
                                        <v-list-tile-content>
                                            <v-list-tile-title>Logout</v-list-tile-title>
                                            <v-list-tile-sub-title></v-list-tile-sub-title>
                                        </v-list-tile-content>
                                        <v-list-tile-action>
                                        </v-list-tile-action>
                                    </v-list-tile>
                                </v-list>
                            </v-card>
                        </v-menu>
                    </v-list>
                </v-navigation-drawer>
                <v-dialog v-model="dialog"
                          max-width="400px">
                    <v-card>
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
export default {
    data() {
        return {
            mini: true,
            drawer: true,
            dialog: false,
            menu: false
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
        logout() {
            this.$store.dispatch('logout')
            this.$router.push({ name: 'login' })
        }
    }
}
</script>

<style>
</style>
