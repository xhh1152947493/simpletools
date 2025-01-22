<script setup>
defineProps({
  type: {
    type: String,
    required: true,
  },
});
</script>

<template>
  <div class="tool-component">
    工具组件: {{ type }}
  </div>
</template>

<style scoped>
.tool-component {
  padding: 1rem;
  border: 1px solid #ccc;
  border-radius: 4px;
  background-color: #f9f9f9;
}
</style>