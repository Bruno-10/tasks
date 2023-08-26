<template>
  <v-text-field
    class="ui-text-field"
    v-bind="$attrs"
    outlined
    dense
    v-on="$listeners"
  >
    <template #prepend-inner>
      <slot name="prepend-inner" />
    </template>
    <template #prepend>
      <slot name="prepend" />
    </template>
    <template #append-outer>
      <slot name="append-outer" />
    </template>
    <template #append>
      <slot name="append" />
    </template>
  </v-text-field>
</template>

<script>
export default {
  name: 'UiTextField',
}
</script>
<style>
.ui-text-field {
  max-width: 220px;
}
</style>
