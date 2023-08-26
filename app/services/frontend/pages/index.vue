<template>
  <v-card
    justify="center"
    align="center"
    min-height="calc(100vh - 24px)"
    class="d-flex flex-column"
  >
    <v-card-title class="justify-space-between"> Tasks </v-card-title>
    <v-card-text style="height: 100%; display: flex; flex: 1 1 auto">
      <lazy-tasks />
    </v-card-text>
    <v-card-actions>
      <lazy-add-task />
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  name: 'IndexPage',
  data() {
    return {
      tasks: [
        {
          ID: 'asdasd',
          name: 'Test',
          description:
            'Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.',
          type: 'Work',
          label: 'Urgent',
          checked: false,
          dueDate: new Date(),
        },
      ],
    }
  },
}
</script>
