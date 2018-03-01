<template>
  <v-app v-resize="onResize">
    <v-navigation-drawer v-model="drawer"
                         :permanent="miniVariant||persistent"
                         :persistent="persistent"
                         :temporary="temporary"
                         :mini-variant="miniVariant"
                         clipped
                         fixed
                         app>
      <v-list dense>
        <v-list-tile value="true"
                     v-for="(item, i) in navMenus"
                     :key="i">
          <v-list-tile-avatar>
            <v-icon v-html="item.icon"></v-icon>
          </v-list-tile-avatar>
          <v-list-tile-content>
            <v-list-tile-title v-text="item.title"></v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
    </v-navigation-drawer>
    <v-toolbar app>
      <v-toolbar-side-icon @click.stop="drawer = !drawer"></v-toolbar-side-icon>
      <v-toolbar-title v-text="title"></v-toolbar-title>
      <v-spacer></v-spacer>
    </v-toolbar>
    <v-content>
      <router-view/>
    </v-content>
    <v-footer app>
      <span>&copy; 2018</span>
    </v-footer>
  </v-app>
</template>

<script>
export default {
    name: 'App',
    data() {
        return {
            navMenus: [
                { icon: 'bubble_chart', title: 'Lastest' },
                { icon: 'list', title: 'List' },
                { icon: 'add', title: 'Add' }
            ],
            title: 'Vuetify.js',
            drawer: true
        }
    },
    computed: {
        currentViewportSize() {
            return this.$store.getters.currentViewportSize
        },
        temporary() {
            if (this.currentViewportSize == 'xs') {
                this.drawer = false
            }
            return this.currentViewportSize == 'xs'
        },
        miniVariant() {
            if (this.currentViewportSize == 'sm' || this.currentViewportSize == 'md') {
                this.drawer = true
            }
            return this.currentViewportSize == 'sm' || this.currentViewportSize == 'md'
        },
        persistent() {
            if (this.currentViewportSize == 'lg' || this.currentViewportSize == 'xl') {
                this.drawer = true
            }
            return this.currentViewportSize == 'lg' || this.currentViewportSize == 'xl'
        }
    },
    methods: {
        onResize() {
            this.$store.commit('setViewportSize', window.innerWidth)
        }
    },
    mounted() {
        this.onResize()
    }
}
</script>
