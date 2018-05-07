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
            <v-tabs slot="extension"
                    centered
                    v-model="tabs"
                    slider-color="black"
                    color="transparent"
                    style="margin-top:18px;"
                    @input="changeTab">
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
                          :rows-per-page-items="pagingSetting">
                <!-- :search="search" -->
                <template slot="items"
                          slot-scope="props">
                    <td>
                        <b class="subheading">{{ props.item.documentNo }}</b>
                        <div v-if="tabs == '0'"
                             style="font-size:11px;"
                             class="grey--text lighten-2">{{ getOutletDescription(props.item.outletId) }}</div>
                    </td>
                    <td>{{ props.item.batchNo }}</td>
                    <td>{{ $moment(props.item.date).format('DD/MMM/YYYY') }}</td>
                    <td class="text-xs-right">{{ to2Decimal(props.item.in>0? props.item.unitCost:props.item.sellingPrice) }}</td>
                    <td class="text-xs-right green--text">{{ hideIfZeroValue(props.item.in) }}</td>
                    <td class="text-xs-right red--text">{{ hideIfZeroValue(props.item.out) }}</td>
                    <td v-if="tabs == '0'"
                        class="text-xs-right">{{ to2Decimal(props.item.utdBalance) }}</td>
                    <td v-if="tabs == '0'"
                        class="text-xs-right">{{ to2Decimal(props.item.utdCost) }}</td>
                    <td class="text-xs-center">
                        <v-tooltip v-if="props.item.remark"
                                   bottom>
                            <v-icon slot="activator"
                                    color="red">announcement</v-icon>
                            <span>{{ props.item.remark }}</span>
                        </v-tooltip>

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
            <v-card class="pa-3"
                    v-if="dialog">
                <v-card-title class="title">
                    <v-layout row
                              wrap>
                        Inventory Adjustment
                    </v-layout>
                </v-card-title>
                <v-card-text class="grey lighten-4"
                             style="border-radius:12px;">
                    <v-radio-group v-model="inOrOut"
                                   row
                                   hide-details
                                   class="pt-0">
                        <v-radio label="In"
                                 value="in"></v-radio>
                        <v-radio label="Out"
                                 value="out"></v-radio>
                    </v-radio-group>
                </v-card-text>
                <v-card-text>
                    <v-layout row
                              wrap>
                        <v-flex xs12
                                md6
                                class="pr-3">
                            <v-menu ref="datepicker"
                                    lazy
                                    :close-on-content-click="false"
                                    v-model="datepicker"
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
                                <v-date-picker v-model="model.date"
                                               year-icon="mdi-calendar-blank"
                                               @input="$refs.datepicker.save($date)"></v-date-picker>

                            </v-menu>
                        </v-flex>
                        <v-flex xs12
                                md6
                                class="pl-3">
                            <v-select v-model="model.outletId"
                                      label="Outlet"
                                      placeholder="Please select"
                                      autocomplete
                                      :items="outlets"
                                      item-value="_id"
                                      item-text="description"></v-select>
                        </v-flex>
                        <v-flex xs12
                                md6
                                class="pr-3">
                            <v-text-field v-model="model.documentNo"
                                          label="Document Number"
                                          placeholder="Document number"></v-text-field>

                        </v-flex>
                        <v-flex xs12
                                md6
                                class="pl-3">
                            <v-text-field v-show="inOrOut == 'in'"
                                          v-model="model.batchNo"
                                          label="Batch Number"
                                          placeholder="Batch number"></v-text-field>
                        </v-flex>
                        <v-flex xs12
                                md6
                                class="pr-3">
                            <v-text-field v-if="inOrOut == 'in'"
                                          v-model="model.in"
                                          label="IN"
                                          placeholder="In Stock"></v-text-field>
                            <v-text-field v-if="inOrOut== 'out'"
                                          v-model="model.out"
                                          label="OUT"
                                          placeholder="Out Stock"></v-text-field>
                        </v-flex>

                        <v-flex xs12
                                md6
                                class="pl-3">
                            <v-text-field v-if="inOrOut == 'in'"
                                          v-model="model.unitCost"
                                          type="number"
                                          label="Unit Cost"
                                          placeholder="Cost per unit"
                                          prefix="$"></v-text-field>
                            <v-text-field v-if="inOrOut == 'out'"
                                          v-model="model.sellingPrice"
                                          type="number"
                                          label="Selling Price"
                                          placeholder="Price per unit"
                                          prefix="$"></v-text-field>
                        </v-flex>
                        <v-flex xs12>
                            <v-text-field v-model="model.remark"
                                          multi-line
                                          label="Remark"
                                          placeholder="Please describe"></v-text-field>
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
import validator from '@/mixins/validator'
export default {
    mixins: [validator],
    data() {
        return {
            title: null,
            datepicker: false,
            tabs: '0',
            outlets: null,
            dialog: false,
            inOrOut: 'in',
            search: null,
            details: [],
            pagingSetting: [10, 20, 30, { text: 'All', value: -1 }],
            model: {
                documentNo: null,
                batchNo: null,
                outletId: null,
                date: null,
                unitCost: null,
                utpTotalCost: null,
                sellingPrice: null,
                totalSales: null,
                in: null,
                out: null,
                remark: null
            }
        }
    },
    computed: {
        headers() {
            let arr_header = [
                { text: 'Document No', align: 'left', sortable: true, value: 'documentNo' },
                { text: 'Batch No', align: 'left', sortable: true, value: 'batchNo' },
                { text: 'Date', align: 'left', sortable: true, value: 'date' },
                { text: 'Cost($)', align: 'right', sortable: true, value: 'unitCost' },
                { text: 'In', align: 'right', sortable: true, value: 'in' },
                { text: 'Out', align: 'right', sortable: true, value: 'out' },
                { text: 'Balance Quantity', align: 'right', value: 'utdBalance' },
                { text: 'Balance Cost($)', align: 'right', value: 'utdCost' },
                { text: '', align: 'center', sortable: false, value: 'remark' }
            ]
            if (this.tabs !== '0') {
                this._.pullAt(arr_header, [6, 7])
            }

            return arr_header
        },
        formattedDate() {
            if (!this.model.date) {
                return null
            }

            return this.$moment(this.model.date).format('DD/MMMM/YYYY')
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
        }
    },
    methods: {
        changeTab() {
            console.log(this.tabs)
        },
        to2Decimal(value) {
            return parseFloat(value).toFixed(2) || '0.00'
        },
        getOutletDescription(str_id) {
            let vm = this
            if (!vm.outlets || !str_id) {
                return ''
            }

            return (
                vm.outlets.find(obj_outlet => {
                    return obj_outlet._id == str_id
                }).description || ''
            )
        },
        hideIfZeroValue(value) {
            if (value == 0) {
                return ''
            }
            return value
        },
        getCurrentItem() {
            let vm = this,
                obj_result = {}

            vm.axios
                .get(`/items/${vm.currentId}`)
                .then(obj_response => {
                    if (!obj_response || !obj_response.data) {
                        return
                    }

                    vm.title = obj_response.data.code + ' - ' + obj_response.data.description
                    vm.details = obj_response.data.details || []
                })
                .catch(obj_exception => {})
        },
        addDetail() {
            let vm = this
            vm.model.date = vm.$moment(vm.model.date)
            vm.model.out = -Math.abs(vm.model.out)
            vm.model.sellingPrice = -Math.abs(vm.model.sellingPrice)
            vm
                .post(`/items/${vm.currentId}/details`, vm.model)
                .then(obj_response => {
                    if (!obj_response || !obj_response.data) {
                        return
                    }

                    vm.dialog = false
                })
                .catch(obj_exception => {})
        },
        load() {
            let vm = this

            vm.axios
                .get('/settings/outlets')
                .then(obj_response => {
                    vm.outlets = obj_response.data
                })
                .catch(obj_exception => {})

            vm.getCurrentItem()
        }
    },
    watch: {
        currentId() {
            this.getCurrentItem()
        }
    },
    mounted() {
        this.load()
    }
}
</script>

<style>
.separator_before {
    position: relative;
}
.separator_before::before {
    content: '';
    position: absolute;
    left: -12px;
    top: 10%;
    width: 1px;
    height: 80%;
    background: rgba(0, 0, 0, 0.11);
}
</style>
