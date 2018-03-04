<template>
    <v-flex lg9
            md9
            sm12
            style="position:relative;">
        <v-card tile
                height="100%"
                style="z-index:6; overflow:hidden;">
            <v-toolbar color="grey lighten-4"
                       flat
                       prominent
                       extended>
                <v-btn icon
                       class="hidden-md-and-up"
                       @click.native="$router.push({name:'inventory:list'})">
                    <v-icon>keyboard_backspace</v-icon>
                </v-btn>
                <v-spacer></v-spacer>
                <v-btn icon>
                    <v-icon class="blue-grey--text text--darken-3">delete</v-icon>
                </v-btn>
                <v-btn icon>
                    <v-icon class="blue-grey--text text--darken-3">more_vert</v-icon>
                </v-btn>
                <div slot="extension">
                    <v-toolbar-title>{{ title }}</v-toolbar-title>
                </div>
            </v-toolbar>
            <v-divider></v-divider>
            <v-card-text>
                <v-data-table :headers="headers"
                              :items="details"
                              :search="search"
                              :rows-per-page-items="pagingSetting">
                    <template slot="items"
                              slot-scope="props">
                        <td>{{ props.item.documentNo }}</td>
                        <td>{{ props.item.batchNo }}</td>
                        <td>{{ $moment(props.item.date).format('DD/MMM/YYYY') }}</td>
                        <td>{{ props.item.unitCost.toFixed(2) }}</td>
                        <td class="text-xs-right green--text">{{ props.item.in }}</td>
                        <td class="text-xs-right red--text">{{ props.item.out }}</td>
                        <td class="text-xs-right">{{ props.item.balanceQuantity }}</td>
                    </template>
                    <v-alert slot="no-results"
                             :value="true"
                             color="error"
                             icon="warning">
                        Your search for "{{ search }}" found no results.
                    </v-alert>
                </v-data-table>
            </v-card-text>
            <v-btn color="pink accent-1"
                   dark
                   absolute
                   fab
                   style="bottom: 16px; right: 16px;">
                <v-icon>add</v-icon>
            </v-btn>
        </v-card>
    </v-flex>
</template>

<script>
export default {
    data() {
        return {
            search: null,
            headers: [
                { text: 'Document No', align: 'left', sortable: true, value: 'documentNo' },
                { text: 'Batch No', align: 'left', sortable: true, value: 'batchNo' },
                { text: 'Date', align: 'left', sortable: true, value: 'date' },
                { text: 'Cost', align: 'left', sortable: true, value: 'unitCost' },
                { text: 'In', align: 'right', sortable: true, value: 'in' },
                { text: 'Out', align: 'right', sortable: true, value: 'out' },
                { text: 'Balance Quantity', align: 'right', sortable: true, value: 'balanceQuantity' }
            ],
            details: [],
            currentItem: {},
            pagingSetting: [10, 20, 30, { text: 'All', value: -1 }]
        }
    },
    computed: {
        currentViewportSize() {
            return this.$store.getters.currentViewportSize
        },
        isDualColomnLayout() {
            if (this.currentViewportSize == 'xs' || this.currentViewportSize == 'sm') {
                return false
            }
            return true
        },
        currentId() {
            return this.$route.params.id
        },
        title() {
            return this.currentItem.code + ' - ' + this.currentItem.description
        }
    },
    methods: {
        getCurrentItem() {
            let vm = this,
                MockData = require('./../data.js'),
                obj_result = {}

            MockData.getInventoryById(vm.currentId).then(obj_response => {
                vm.currentItem = obj_response
            })
        },
        load() {
            let vm = this,
                MockData = require('./../data.js')

            this.getCurrentItem()
            MockData.getInventoryDetailById(vm.currentId).then(obj_response => {
                vm.details = obj_response
            })
        }
    },
    watch: {
        currentId() {
            this.load()
        }
    },
    mounted() {
        this.load()
    }
}
</script>

<style>

</style>
