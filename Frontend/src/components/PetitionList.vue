<template>
  <div class="container">
    <div class="col" v-for="column in columns" :key="column.length">
    <PetitionCard class="petition-card"
      v-for="(petition) of column"
      v-bind:petition="petition"
      :key="petition.id"
    />
    </div>
  </div>
</template>



<script>
import PetitionCard from './PetitionCard.vue'
export default {
  name: 'PetitionList',
  data() {
    return {
      cols: 2
    }
  },
  props: {
    petitions: Array,
  },
  components: {
    PetitionCard
  },
  computed: {
    columns () {
      let columns = []
      let mid = Math.ceil(this.petitions.length / this.cols)
      for (let col = 0; col < this.cols; col++) {
        columns.push(this.petitions.slice(col * mid, col * mid + mid))
      }
      return columns
    },
    rows() {
      let rows = []
      let mid = Math.ceil(this.petitions.length / this.cols)
      for (let row = 0; row < this.cols; row++) {
        rows.push(this.petitions.slice(row * mid, row * mid + mid))
      }
      return rows
    }
  }

}
</script>

<style>
.container {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  margin: 0 auto;
  justify-content: center;
}
.col {
  flex-grow: 0;
  display: flex;
  flex-direction:column;
}
.petition-card {
  position: relative;
}
</style>