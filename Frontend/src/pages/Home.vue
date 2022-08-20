<template>
  <title>Студенческая инициатива</title>
  <header class="header"><Header></Header> </header>
  <HeaderQuote></HeaderQuote>
  <petition-list :petitions="petitions"></petition-list>
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
      loaded: false,
      petitions: [],
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
        this.loaded = true;
      }, error => {
        console.log(error)
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
</style>