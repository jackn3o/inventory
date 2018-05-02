<template>
    <div>
        <v-toolbar color="grey lighten-4"
                   flat
                   prominent
                   extended>
            <v-btn icon
                   class="hidden-md-and-up"
                   @click.native="$router.push({name:'inventory.list'})">
                <v-icon>keyboard_backspace</v-icon>
            </v-btn>
            <v-toolbar-title>{{ title }}</v-toolbar-title>
            <v-spacer></v-spacer>
            <v-btn icon>
                <v-icon class="blue-grey--text text--darken-3">delete</v-icon>
            </v-btn>
            <v-btn icon>
                <v-icon class="blue-grey--text text--darken-3">more_vert</v-icon>
            </v-btn>
            <v-tabs slot="extension"
                    centered
                    v-model="tabs"
                    slider-color="black"
                    color="transparent"
                    style="margin-top:18px;">
                <v-tab key="0">
                    All
                </v-tab>
                <v-tab v-for="outlet in outlets"
                       :key="outlet._id">
                    {{ outlet.description }}
                </v-tab>
            </v-tabs>
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
                    <td>{{ props.item.outlet }}</td>
                    <td>{{ $moment(props.item.date).format('DD/MMM/YYYY') }}</td>
                    <td>{{ props.item.unitCost.toFixed(2) }}</td>
                    <td class="text-xs-right green--text">{{ props.item.in }}</td>
                    <td class="text-xs-right red--text">{{ props.item.out }}</td>
                    <td class="text-xs-right">{{ props.item.balanceQuantity }}</td>
                    <td class="text-xs-center">
                        <v-icon color="red"
                                v-if="props.item.remark">announcement
                        </v-icon>
                    </td>
                </template>
                <v-alert slot="no-results"
                         :value="true"
                         color="error"
                         icon="warning">
                    Your search for "{{ search }}" found no results.
                </v-alert>
            </v-data-table>
        </v-card-text>
        <v-btn v-if="!dialog"
               color="pink"
               dark
               absolute
               fab
               style="bottom: 16px; right: 16px;"
               @click.native="dialog=true">
            <v-icon>add</v-icon>
        </v-btn>
        <v-dialog v-model="dialog"
                  max-width="500px">
            <v-card class="pa-3">
                <v-card-title class="title">
                    Inventory Adjustment
                    <v-spacer></v-spacer>
                    <v-radio-group v-model="inOrOut"
                                   row
                                   hide-details
                                   class="pt-0">
                        <v-radio label="In"
                                 value="in"></v-radio>
                        <v-radio label="Out"
                                 value="out"></v-radio>
                    </v-radio-group>
                </v-card-title>
                <v-card-text>
                    <v-layout row
                              wrap>
                        <v-flex xs12
                                md6
                                class="pr-3">
                            <v-menu ref="menu2"
                                    lazy
                                    :close-on-content-click="false"
                                    v-model="menu2"
                                    transition="scale-transition"
                                    offset-y
                                    full-width
                                    :nudge-right="40"
                                    min-width="290px"
                                    :return-value.sync="date">
                                <v-text-field slot="activator"
                                              label="Date"
                                              placeholder="Select Date"
                                              v-model="formattedDate"
                                              prepend-icon="event"
                                              readonly></v-text-field>
                                <v-date-picker v-model="date"
                                               @input="$refs.menu2.save(date)"></v-date-picker>

                            </v-menu>
                        </v-flex>
                        <v-flex xs12
                                md6
                                class="pl-3">
                            <v-text-field v-model="model.outlet"
                                          label="Outlet"
                                          placeholder="Select Outlet"
                                          hint="optional"></v-text-field>
                        </v-flex>
                        <v-flex xs12
                                md6
                                class="pr-3">
                            <v-text-field v-model="model.documentNo"
                                          label="Document Number"></v-text-field>

                        </v-flex>
                        <v-flex xs12
                                md6
                                class="pl-3">
                            <v-text-field v-show="inOrOut == 'in'"
                                          v-model="model.batchNo"
                                          label="Batch Number"></v-text-field>
                        </v-flex>
                        <v-flex xs12
                                md6
                                class="pr-3">
                            <v-text-field v-if="inOrOut == 'in'"
                                          label="IN"
                                          placeholder="In Stock"></v-text-field>
                            <v-text-field v-if="inOrOut== 'out'"
                                          label="OUT"
                                          placeholder="Out Stock"></v-text-field>
                        </v-flex>

                        <v-flex xs12
                                md6
                                class="pl-3">
                            <v-text-field v-model="model.unitCost"
                                          type="number"
                                          label="Cost"></v-text-field>
                        </v-flex>

                    </v-layout>
                </v-card-text>

                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="primary"
                           flat
                           @click.stop="dialog=false">Close</v-btn>
                    <v-btn color="primary"
                           flat
                           @click.stop="addDetail">Save</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
    </div>
</template>

<script>
export default {
    data() {
        return {
            date: null,
            menu2: false,
            tabs: null,
            outlets: null,
            dialog: false,
            inOrOut: 'in',
            search: null,
            headers: [
                { text: 'Document No', align: 'left', sortable: true, value: 'documentNo' },
                { text: 'Batch No', align: 'left', sortable: true, value: 'batchNo' },
                { text: 'Outlet', align: 'left', sortable: true, value: 'outlet' },
                { text: 'Date', align: 'left', sortable: true, value: 'date' },
                { text: 'Cost', align: 'left', sortable: true, value: 'unitCost' },
                { text: 'In', align: 'right', sortable: true, value: 'in' },
                { text: 'Out', align: 'right', sortable: true, value: 'out' },
                { text: 'Balance Quantity', align: 'right', sortable: true, value: 'balanceQuantity' },
                { text: '', align: 'center', sortable: false, value: 'remark' }
            ],
            details: [],
            currentItem: {},
            pagingSetting: [10, 20, 30, { text: 'All', value: -1 }],
            model: {
                documentNo: null,
                batchNo: null,
                outlet: null,
                date: null,
                unitCost: null,
                in: null,
                out: null,
                remark: null
            }
        }
    },
    computed: {
        formattedDate() {
            if (!this.date) {
                return null
            }

            return this.$moment(this.date).format('DD/MMMM/YYYY')
        },
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
            return ''
            // return this.currentItem.code + ' - ' + this.currentItem.description
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
        addDetail() {},
        load() {
            let vm = this,
                MockData = require('./../data.js')

            this.axios
                .get('/settings/outlets')
                .then(obj_response => {
                    this.outlets = obj_response.data
                })
                .catch(obj_exception => {})

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
