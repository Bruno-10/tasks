<template>
  <v-list>
    <v-list-item-group v-if="tasks.length">
      <v-list-item v-for="task in tasks" :key="task.ID" class="text-start my-2">
        <template>
          <v-list-item-content>
            <v-list-item-title>
              <div>
                {{ task.name }}
                <v-chip class="ml-4 my-2 label" small label>
                  {{ task.type }}
                </v-chip>
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
              <strong> No pending tasks! </strong>
            </v-list-item-title>
          </v-list-item-content>
        </template>
      </v-list-item>
    </v-list-item-group>
  </v-list>
</template>
<script>
const taskLabelColors = {
  Urgent: 'red',
  'Can be postponed': 'green',
  'Not important': 'grey',
}
export default {
  name: 'Tasks',
  props: {
    tasks: {
      type: Array,
      default: () => [],
    },
  },
  methods: {
    taskLabelColors(label) {
      return taskLabelColors[label]
    },
  },
}
</script>
<style lang="scss">
.label {
  max-width: fit-content;
}
</style>
