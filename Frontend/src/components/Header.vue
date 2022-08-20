<template>
  <header>
    <span v-if="!isLoggedIn">
    <router-link to="/sign-in"><span class="sign-in-header-button"><va-button flat class="mr-4 mb-2" color="#ffffff">ВОЙТИ</va-button></span> </router-link>
    <router-link to="/sign-up"><span class="sign-up-header-button"><va-button flat class="mr-4 mb-2" color="#ffffff">РЕГИСТРАЦИЯ</va-button></span></router-link>
    </span>
    <span v-else>
      <p class="username-header">{{getName}}</p>
    </span>
    <router-link to="/about"><span class="about-header-button"><va-button flat class="mr-4 mb-2" color="#ffffff">О ПРОЕКТЕ</va-button></span></router-link>
    <router-link :to="{name: 'create-petition'}"><span class="create-petition-header-button"><va-button flat class="mr-4 mb-2" color="#ffffff">СОЗДАТЬ ПЕТИЦИЮ</va-button></span></router-link>
    <div class="overlay">
      <p class="title-style"><router-link to="/">Студенческая<br>инициатива</router-link></p>
    </div>
  </header>

</template>

<script>
import {mapGetters} from "vuex";

export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: "Header",
  computed: {
    ...mapGetters([
        'isLoggedIn',
        'getName',
    ])
  },
  methods: {},
}
</script>

<style>
.sign-in-header-button{
  position: absolute;
  top: 0;
  right: 440px;
  font-size: medium;
  font-weight: bold;
  color: #ffffff;
}

.username-header{
  font-family: 'Rubik', sans-serif;
  position: relative;
  top: 0;
  text-align: left;
  margin-left: 0.5%;
  margin-top: 0.55%;
  font-size: 19px;
  text-shadow: black 0.05em 0.05em 0.05em;
  color: #ffffff;
}

.sign-up-header-button{
  position: absolute;
  top: 0;
  right: 300px;
  font-size: medium;
  font-weight: bold;
  color: #ffffff;
}
.about-header-button{
  position: absolute;
  top: 0;
  right: 180px;
  font-size: medium;
  font-weight: bold;
  color: #ffffff;
}
.create-petition-header-button{
  position: absolute;
  top: 0;
  right: 0;
  font-size: medium;
  font-weight: bold;
  color: #ffffff;
}
header {
  background: url("../assets/head.jpg");
  margin: 0;
  text-align: center;
  width: 100%;
  height: 50vh;
  background-size: cover;
  background-attachment: fixed;
  position: relative;
  overflow: hidden;
}
header .overlay{
  width: 100%;
  height: 100%;
  padding: 50px;
  color: #FFF;
  text-shadow: 1px 1px 1px #ffffff;
}

.title-style {
  text-transform: Uppercase;
  margin-bottom: .5em;
  font-family: 'Rubik', sans-serif;
  font-size: 2rem;
  font-weight: bold;
  font-style: italic;
  color: #E4E5E6;
  position: relative;
  background: linear-gradient(to right, #000000, #000000, #000000);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}
</style>