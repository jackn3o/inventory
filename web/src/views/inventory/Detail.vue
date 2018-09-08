<template>
    <div>
        <v-toolbar color="grey lighten-4"
                   flat
                   :dense="$vuetify.breakpoint.mdAndDown"
                   extended>
            <v-btn icon
                   class="hidden-md-and-up"
                   @click.native="$router.push({name:'inventory.list'})">
                <v-icon>keyboard_backspace</v-icon>
            </v-btn>
            <v-toolbar-title>{{ title }}</v-toolbar-title>
            <datepicker iconButton></datepicker>
            <v-spacer></v-spacer>
            <v-tabs slot="extension"
                    :height="$vuetify.breakpoint.mdAndDown? 36: 48"
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
                <template slot="no-data">
                    <v-layout align-center
                              justify-center
                              @click="dialog=true">
                        <v-flex xs12
                                sm8
                                md4
                                class="text-xs-center pa-3">
                            <v-icon color="blue-grey lighten-4"
                                    style="font-size: 100px;">tab</v-icon>
                            <div class="headline grey--text">Nothing here yet</div>
                            <div class="body-1 grey--text">Click to add some item</div>
                        </v-flex>
                    </v-layout>
                </template>
                <v-alert slot="no-results"
                         :value="true"
                         color="error"
                         icon="warning">
                    Your search for "{{ search }}" found no results.
                </v-alert>
            </v-data-table>
        </v-card-text>
        <v-fab-transition>
            <v-btn v-if="!dialog"
                   color="pink"
                   dark
                   absolute
                   fab
                   style="bottom: 16px; right: 16px;"
                   @click.native="dialog=true">
                <v-icon>add</v-icon>
            </v-btn>
        </v-fab-transition>
        <in-or-out v-model="dialog"
                   :id="currentId"
                   :outlets="outlets"
                   @close-dialog="closeDialog()"></in-or-out>
    </div>
</template>

<script>
import datepicker from '@/components/datepicker'
import InOrOut from './InOrOut.vue'

export default {
    components: {
        datepicker,
        InOrOut
    },
    data() {
        return {
            title: null,
            tabs: '0',
            outlets: null,
            dialog: false,
            search: null,
            details: [],
            pagingSetting: [10, 20, 30, { text: 'All', value: -1 }]
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
        isDualColomnLayout() {
            if (this.$vuetify.breakpoint.smAndDown) {
                return false
            }
            return true
        },
        currentId() {
            return this.$route.params.id
        }
    },
    methods: {
        changeTab() {},
        closeDialog() {
            this.dialog = false
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
            this.load()
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
