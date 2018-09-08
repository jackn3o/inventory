<template>
    <div class="pa-3">
        <v-card class="mb-3">
            <v-card-title class="title font-weight-bold">
                How to get your item cost
            </v-card-title>
            <v-card-text>
                <v-radio-group v-model="model.costMethod"
                               :mandatory="false">
                    <v-radio label="Average"
                             :value="1"></v-radio>
                    <v-radio label="Coming Soon"
                             disabled
                             :value="2"></v-radio>
                </v-radio-group>
            </v-card-text>
        </v-card>
        <v-card>
            <v-card-title class="title font-weight-bold">
                Locked record before
            </v-card-title>
            <v-card-text>
                to be coming
            </v-card-text>
        </v-card>
    </div>
</template>

<script>
import validator from '../../mixins/validator'
export default {
    mixins: [validator],
    data() {
        return {
            dialog: false,
            model: {
                costMethod: 1
            }
        }
    }
}
</script>

<style>
</style>
