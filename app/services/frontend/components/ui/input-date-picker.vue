<template>
  <v-menu v-model="menu" :close-on-content-click="false" max-width="290">
    <template v-slot:activator="{ on, attrs }">
      <lazy-ui-text-field
        :value="value"
        clearable
        :label="dateRangeLabel"
        readonly
        v-bind="attrs"
        v-on="on"
        @click:clear="date = null"
      />
    </template>
    <v-date-picker
      v-model="date"
      :min="min"
      @change="
        $emit('change', date)
        value = date
        menu = false
      "
    ></v-date-picker>
  </v-menu>
</template>

<script>
import utils from '~/utils'
export default {
  name: 'UiInputDatePicker',
  props: {
    label: {
      type: String,
      default: () => '',
    },
    min: {
      type: [String, undefined],
      default: () => undefined,
    },
    max: {
      type: [String, undefined],
      default: () => undefined,
    },
  },
  data() {
    return {
      date: '',
      menu: false,
      value: utils.today(),
      dateRangeAlert: false,
    }
  },
  computed: {
    dateRangeLabel() {
      return this.label ? this.label : 'Date Range'
    },
  },
}
</script>
