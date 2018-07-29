<template>
    <v-menu ref="menu"
            :close-on-content-click="false"
            v-model="menu"
            :nudge-right="nudgeRight"
            lazy
            transition="scale-transition"
            offset-y
            full-width
            min-width="290px">
        <v-btn v-if="iconButton"
               slot="activator"
               icon
               class="btnClass">
            <v-icon>{{ icon }}</v-icon>
        </v-btn>
        <v-text-field v-else
                      slot="activator"
                      v-model="date"
                      :label="label"
                      :prepend-icon="icon"
                      :readonly="textReadOnly"></v-text-field>
        <v-date-picker ref="picker"
                       v-model="date"
                       :max="new Date().toISOString().substr(0, 10)"
                       min="1950-01-01"
                       @change="save"></v-date-picker>
    </v-menu>
</template>

<script>
export default {
    props: {
        value: {},
        iconButton: { type: Boolean, default: false },
        btnClass: { type: Object, default: {} },
        label: { type: String, default: 'Date' },
        icon: { type: String, default: 'event' },
        textReadOnly: { type: Boolean, default: true },
        nudgeTop: { type: Number, default: 0 },
        nudgeBottom: { type: Number, default: 0 },
        nudgeLeft: { type: Number, default: 0 },
        nudgeRight: { type: Number, default: 0 }
    },
    data: () => ({
        date: null,
        menu: false
    }),
    watch: {
        menu(val) {
            val && this.$nextTick(() => (this.$refs.picker.activePicker = 'YEAR'))
        }
    },
    methods: {
        save(date) {
            this.$refs.menu.save(date)
        }
    }
}
</script>

<style>
</style>
