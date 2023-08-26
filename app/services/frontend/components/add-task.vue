<template>
  <v-expansion-panels>
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
                :value="form.dueDate"
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
    dialog: false,
    formValid: true,
    form: {
      name: 'Test',
      type: 'Work',
      description: 'tasdd',
      dueDate: utils.today(),
    },
  }),
  mounted() {
    this.addTask()
  },
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
      if (
        differenceInDays(pickedDate, new Date(this.tomorrow)) < 1 &&
        this.form.type.includes('Work')
      ) {
        return 'Urgent'
      }

      if (difference <= 7 && this.form.type.includes('Personal')) {
        return 'Can be postponed'
      }

      if (difference <= 5 && this.form.type.includes('Personal')) {
        return 'Can be postponed'
      }

      if (
        difference <= 31 &&
        this.form.type.includes('Work') &&
        workCondition
      ) {
        return 'Can be postponed'
      }

      return 'Not important'
    },
    addTask() {
      // if ('form' in this.$refs && this.$refs.form.validate()) {
      const label = this.getTaskLabel()

      const data = {
        ...this.form,
        dueDate: `${this.form.dueDate}T00:00:00Z`,
        label,
      }

      console.log(data)

      this.$axios({
        url: 'http://tasks-api.tasks-system.svc.cluster.local:3000/tasks',
        data,
        method: 'POST',
      })
        .then((_res) => {
          this.$emit('added')
        })
        .catch((err) => {
          throw new Error(err)
        })
    },
  },
  // },
}
</script>
