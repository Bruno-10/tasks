<template>
  <v-list style="width: 100%">
    <v-list-item-group v-if="tasks.length">
      <v-list-item v-for="task in tasks" :key="task.ID" class="text-start my-2">
        <template>
          <v-list-item-content>
            <v-list-item-title>
              <div>
                {{ task.name }}
                <v-chip
                  class="ml-4 my-2 label"
                  :color="taskLabelColors(task.label)"
                  small
                  label
                >
                  {{ task.label }}
                </v-chip>
              </div>
            </v-list-item-title>
            <v-list-item-subtitle>
              {{ task.description }}
            </v-list-item-subtitle>
          </v-list-item-content>
        </template>
      </v-list-item>
    </v-list-item-group>
    <v-list-item-group v-else>
      <v-list-item>
        <template>
          <v-list-item-content>
            <v-list-item-title>
              <strong> No pending tasks! </strong></v-list-item-title
            >
          </v-list-item-content>
        </template>
      </v-list-item>
    </v-list-item-group>
  </v-list>
</template>
<script>
const taskLabelColors = {
  Urgent: 'red',
  Health: 'green',
  Personal: 'white',
  Other: 'grey',
  Work: 'lightblue',
}
export default {
  name: 'Tasks',
  data: () => ({
    tasks: [],
  }),
  mounted() {
    this.getTasks()
  },
  methods: {
    taskLabelColors(label) {
      return taskLabelColors[label]
    },
    getTasks() {
      this.$axios
        .get(`http://tasks-api.tasks-system.svc.cluster.local:3000/tasks`)
        .then((res) => {
          this.tasks = res.data.items
        })
        .catch((err) => {
          throw new Error(err)
        })
    },
  },
}
</script>
<style lang="scss">
.label {
  max-width: fit-content;
}
</style>
