<template>
  <v-card
    justify="center"
    align="center"
    max-height="calc(100vh - 24px)"
    style="height: 100vh; overflow-y: hidden"
    class="d-flex flex-column"
  >
    <v-card-title class="justify-space-between"> Tasks </v-card-title>
    <v-card-text style="overflow-y: scroll; flex: 1 1 auto">
      <lazy-tasks :tasks="tasks" />
    </v-card-text>
    <v-card-actions>
      <lazy-add-task @added="getTasks()" />
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  name: 'IndexPage',
  data() {
    return {
      tasks: [],
    }
  },
  mounted() {
    this.getTasks()
  },
  methods: {
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
