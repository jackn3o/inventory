<template>
    <v-navigation-drawer dark
                         permenant
                         stateless
                         absolute
                         hide-overlay
                         :mini-variant="true"
                         :mini-variant-width="60"
                         v-model="value"
                         class="navigation_drawer"
                         style="display: flex; flex-direction: column; z-index: 5; padding-bottom: 0;">
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
            <v-list-tile avatar>
                <v-list-tile-avatar>
                    <v-btn icon
                           color="teal"
                           @click.native="$router.push({name: 'account'})">
                        <v-icon color="white">person</v-icon>
                    </v-btn>
                </v-list-tile-avatar>
                <v-list-tile-content>
                    <v-list-tile-title>Admin</v-list-tile-title>
                </v-list-tile-content>
            </v-list-tile>
        </v-list>
    </v-navigation-drawer>
</template>

<script>
export default {
    props: {
        value: {
            type: Boolean,
            default: true
        }
    },
    data() {
        return {
        }
    }
}
</script>

<style>
</style>
