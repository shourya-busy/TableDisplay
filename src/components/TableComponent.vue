<script setup>
import axios from 'axios'
import { ref, computed, onMounted } from 'vue';
    const data = ref([]);
    const columns = ref([]);
    const sortColumn = ref(null);
    const sortDirection = ref('asc');

    const sortedData = computed(() => {
      if (!sortColumn.value) return data.value;
      return [...data.value].sort((a, b) => {
        const valueA = a[sortColumn.value];
        const valueB = b[sortColumn.value];
        if (valueA < valueB) return sortDirection.value === 'asc' ? -1 : 1;
        if (valueA > valueB) return sortDirection.value === 'asc' ? 1 : -1;
        return 0;
      });
    });

    const fetchData = async () => {
      try {
        const response = await axios.get('http://localhost:8080/leads');
        data.value = response.data;
        columns.value = Object.keys(data.value[0] || {});
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    };

    const toggleSort = (column) => {
      if (sortColumn.value === column) {
        sortDirection.value = sortDirection.value === 'asc' ? 'desc' : 'asc';
      } else {
        sortColumn.value = column;
        sortDirection.value = 'asc';
      }
    };

    onMounted(fetchData);
</script>

<template>
    <table>
      <thead>
        <tr>
          <th v-for="column in columns" :key="column">
            {{ column.toUpperCase() }}
            <button @click="toggleSort(column)">
              <span v-if="sortColumn === column">
                {{ sortDirection === 'asc' ? '▼' : '▲' }}
              </span>
              <span v-else>▲</span>
            </button>
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="row in sortedData" :key="row.id">
          <td v-for="column in columns" :key="column">{{ row[column] }}</td>
        </tr>
      </tbody>
    </table>
</template>

<style scoped>
table{
  padding: 5%;
  border: 1px rgb(129, 129, 129) solid;
}

th{
  color: aliceblue;
}
td{
  border: 0.5px rgb(129, 129, 129) solid;
}

button{
  background-color: rgb(47, 47, 47);
  text-align: justify;
  padding: 3px;
  color: aliceblue;
}
</style>