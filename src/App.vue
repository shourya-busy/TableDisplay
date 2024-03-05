<script setup>
import tableComponent from './components/TableComponent.vue'

</script>

<template>
  <main>
    <header>
      <h1>Data Table</h1>
    </header>
      <tableComponent></tableComponent>
  </main>
</template>


<style scoped>

main{
  display: flex;
  flex-direction: column;
  margin: 10%;
}

header {
  font-size: xx-large;
  background-color: rgb(1, 38, 38);
  padding: 10px;
}
</style>