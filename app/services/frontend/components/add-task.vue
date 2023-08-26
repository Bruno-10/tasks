<template>
  <v-expansion-panels v-model="menu">
    <v-expansion-panel>
      <v-expansion-panel-header> Add task </v-expansion-panel-header>
      <v-expansion-panel-content>
        <v-form
          ref="form"
          v-model="formValid"
          style="width: 100%"
          lazy-validation
        >
          <v-row no-gutters class="d-flex flex-column px-4">
            <v-col cols="12" no-gutters class="d-flex justify-start">
              <lazy-ui-text-field
                v-model="form.name"
                label="Name"
                class="mr-4"
                :rules="[requiredRule]"
              />
              <lazy-ui-text-field
                v-model="form.type"
                class="mr-4"
                label="Type"
                :rules="[requiredRule]"
              />
              <lazy-ui-input-date-picker
                label="Due Date"
                :min="today"
                @change="form.dueDate = $event"
              />
            </v-col>
            <v-col cols="12" no-gutters class="d-flex justify-center">
              <v-textarea
                v-model="form.description"
                outlined
                auto-gUTrow
                rows="3"
                row-height="32"
                label="Description"
                :rules="[requiredRule]"
              />
            </v-col>
            <v-btn color="#323639" class="white--text" @click="addTask">
              Add task
            </v-btn>
          </v-row>
        </v-form>
      </v-expansion-panel-content>
    </v-expansion-panel>
  </v-expansion-panels>
</template>
<script>
import differenceInDays from 'date-fns/differenceInDays'
import utils from '~/utils'
export default {
  name: 'AddTask',
  data: () => ({
    menu: undefined,
    formValid: true,
    form: {
      name: '',
      type: '',
      description: '',
      dueDate: utils.today(),
    },
  }),
  computed: {
    today() {
      return utils.today()
    },
    tomorrow() {
      return utils.tomorrow()
    },
  },
  methods: {
    requiredRule(v) {
      return !!v || 'This field is required'
    },
    getTaskLabel() {
      const pickedDate = new Date(this.form.dueDate)
      const todayFormated = new Date(this.today)
      const workCondition =
        this.form.name.includes('PLO') || this.form.name.includes('GJL')
      const difference = differenceInDays(todayFormated, pickedDate)
      const type = this.form.type.toLowerCase()
      if (
        differenceInDays(pickedDate, new Date(this.tomorrow)) < 1 &&
        type.includes('work')
      ) {
        return 'Urgent'
      }

      if (difference <= 7 && type.includes('personal')) {
        return 'Can be postponed'
      }

      if (difference <= 5 && type.includes('personal')) {
        return 'Can be postponed'
      }

      if (difference <= 31 && type.includes('work') && workCondition) {
        return 'Can be postponed'
      }

      return 'Not important'
    },
    addTask() {
      if ('form' in this.$refs && this.$refs.form.validate()) {
        const label = this.getTaskLabel()

        const data = {
          ...this.form,
          dueDate: `${this.form.dueDate}T00:00:00Z`,
          label,
        }

        this.$axios({
          url: 'http://tasks-api.tasks-system.svc.cluster.local:3000/tasks',
          data,
          method: 'POST',
        })
          .then((_res) => {
            this.menu = false
            this.$refs.form.reset()
            this.$emit('added')
          })
          .catch((err) => {
            throw new Error(err)
          })
      }
    },
  },
}
</script>
