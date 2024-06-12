<script lang="ts" setup>
import { reactive } from 'vue'
import { SetBrightness } from '../../wailsjs/go/main/App'

const data = reactive<{ brightness: number }>({
    brightness: 0,
})

function setBrightness() {
    SetBrightness(data.brightness).then(result => {
        console.log(result)
    })
}
</script>

<template>
    <input type="range" min="0" max="100" step="1" v-model="data.brightness" />
    <button @click="setBrightness">Set Brightness</button>
</template>

<style scoped></style>
