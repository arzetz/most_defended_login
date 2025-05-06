<template>
  <button @click="handleClick" class="my-button">
    <slot />
  </button>
</template>

<script lang="ts" setup>
import axios from "axios";
import { ref } from "vue";
const isConnected = ref(false)
defineProps({
  label: String,
  modelValue: Boolean,
});
const emit = defineEmits<{
  (e: 'update:is-connected', value: boolean): void
}>()

async function handleClick(): Promise<void> {
  if (!isConnected.value){
  try {
    await axios.patch("http://localhost:8080/api/connectdb");
    alert("Подняли дб");
    isConnected.value = true
    emit('update:is-connected', isConnected.value)
  } catch (err) {
    alert("Дб не хочет ща подниматься");
  }
  }else{
    try {
    await axios.patch("http://localhost:8080/api/disconnectdb");
    alert("Оустили дб");
    isConnected.value = false
    emit('update:is-connected', isConnected.value)
  } catch (err) {
    alert("Дб положен");
    }
  }
}
</script>
