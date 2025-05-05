<template>
  <button @click="handleClick" class="my-button">
    {{ label }}
  </button>
</template>

<script lang="ts" setup>
import axios from "axios";

defineProps({
  label: String,
});

async function handleClick(): Promise<void> {
  try {
    await axios.patch("http://localhost:8080/api/connectdb");
    alert("Подняли дб");
  } catch (err) {
    alert("Дб не хочет ща подниматься, отъебитесь ваще");
  }
}
</script>
