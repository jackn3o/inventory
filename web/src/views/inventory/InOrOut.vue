<template>
    <v-dialog v-model="value"
              max-width="600px">
        <v-card class="pa-3"
                v-if="value">
            <!-- <v-card-title class="title">
                    <v-layout row
                              wrap>
                        Inventory Adjustment
                    </v-layout>
                </v-card-title> -->
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
                                v-model="datepicker"
                                :close-on-content-click="false"
                                lazy
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
                                           @input="$refs.datepicker.save(date)"></v-date-picker>
                        </v-menu>
                    </v-flex>
                    <v-flex xs12
                            md6
                            class="pl-3">
                        <v-autocomplete v-model="model.outletId"
                                        label="Outlet"
                                        placeholder="Please select"
                                        autocomplete
                                        :items="outlets"
                                        item-value="_id"
                                        item-text="description"></v-autocomplete>
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
                        <v-textarea v-model="model.remark"
                                    label="Remark"
                                    placeholder="Please describe"></v-textarea>
                    </v-flex>
                </v-layout>
            </v-card-text>

            <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="primary"
                       flat
                       @click.stop="closeDialog">Close</v-btn>
                <v-btn color="primary"
                       flat
                       @click.stop="inOrOut=='in'?inStock():outStock()">Save</v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script>
import validator from '@/mixins/validator'

export default {
    mixins: [validator],
    props: {
        value: {
            type: Boolean,
            default: false
        },
        id: {
            type: String,
            default: ''
        },
        outlets: {
            type: Array,
            default: []
        }
    },
    data() {
        return {
            date: null,
            inOrOut: 'in',
            datepicker: false,
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
        formattedDate() {
            return this.parseDate(this.model.date)
        }
    },
    methods: {
        parseDate(date) {
            if (!date) {
                return null
            }

            return this.$moment(date).format('DD/MMM/YYYY')
        },
        inStock() {
            let vm = this
            vm.model.date = vm.$moment(vm.model.date)
            vm
                .post(`/items/${vm.id}/in`, vm.model)
                .then(obj_response => {
                    if (!obj_response || !obj_response.data) {
                        return
                    }

                    vm.closeDialog()
                })
                .catch(obj_exception => {})
        },
        outStock() {
            let vm = this
            vm.model.date = vm.$moment(vm.model.date)
            vm.model.out = -Math.abs(vm.model.out)
            vm.model.sellingPrice = -Math.abs(vm.model.sellingPrice)
            vm
                .post(`/items/${vm.id}/out`, vm.model)
                .then(obj_response => {
                    if (!obj_response || !obj_response.data) {
                        return
                    }

                    vm.closeDialog()
                })
                .catch(obj_exception => {})
        },
        closeDialog() {
            this.$emit('close-dialog')
        }
    }
}
</script>

<style>
</style>
