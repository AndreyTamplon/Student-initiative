<template>
  <title>Студенческая инициатива</title>
  <header class="header"><Header></Header> </header>
  <HeaderQuote></HeaderQuote>
  <petition-list v-if="dataReceived === true" :petitions="petitions"></petition-list>
  <div v-if="dataReceived === false">
    <div class="waiting-circle">
      <va-progress-circle size="250px" indeterminate />
    </div>
  </div>
  <br><br>
  <Footer></Footer>
</template>

<script>
import PetitionList from '@/components/PetitionList.vue'
import Header from "@/components/Header";
import Footer from "@/components/Footer";
import HeaderQuote from "@/components/HeaderQuote";
import axios from "axios";
import {BASE_CONTENT_URL} from "@/main";
export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: "Home",
  components: {
    PetitionList,
    Header,
    Footer,
    HeaderQuote,
  },
  data() {
    return {
      petitions: [],
      dataReceived: false,
    }
  },
  created() {
    this.getPetitionsByIdRange(1, 20);
  },
  methods: {
    getPetitionsByIdRange(idFrom, idTo){
      axios.get(BASE_CONTENT_URL + 'petitions', {
        headers: {
          idFrom: idFrom,
          idTo: idTo,
        }
      }).then(response => {
        console.log(response);
        this.petitions = response.data;
        this.dataReceived = true;
      }, error => {
        console.log(error)
        this.dataReceived = false;
      })
    },
  }
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Source+Sans+Pro:ital,wght@0,400;1,700&display=swap');
@import url('https://fonts.googleapis.com/icon?family=Material+Icons');
.header{
  margin: 0 0 0;
}
.waiting-circle{
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>