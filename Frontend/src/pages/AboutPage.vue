<template>
  <title>О проекте</title>
  <div>
    <Header></Header>
    <HeaderQuote></HeaderQuote>
  </div>
  <div class="about-text">
    <p>Проект создан для того, чтобы сократить дистанцию между студентами и руководством вуза. </p>

    <br>
    <br>
    <br>
  </div>
  <Footer ></Footer>
</template>

<script>
import Header from "@/components/Header";
import HeaderQuote from "@/components/HeaderQuote";
import Footer from "@/components/Footer";
import {mapGetters} from "vuex";

export default {
  name: "AboutPage",
  data() {
    return {
      petition: [],
    };
  },
  components: {
    Header,
    HeaderQuote,
    Footer,
  },
  computed: {
    ...mapGetters([
        "getToken",
    ]),
  },
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@200;600&family=Josefin+Sans:ital,wght@1,700&display=swap');
.about-text {
  font-family: 'Josefin Sans', sans-serif;
  font-size: 1.5rem;
  font-weight: 600;
  color: #000000;
  text-align: center;
  margin-left: 8%;
  margin-right: 8%;
}
</style>