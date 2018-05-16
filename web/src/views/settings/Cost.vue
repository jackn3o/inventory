<template>
    <div>
        <div class="subheading mt-3">
            How to get your item cost
        </div>
        <v-radio-group v-model="model.costMethod"
                       :mandatory="false">
            <v-radio label="Average"
                     :value="1"></v-radio>
            <v-radio label="Coming Soon"
                     disabled
                     :value="2"></v-radio>
        </v-radio-group>
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
