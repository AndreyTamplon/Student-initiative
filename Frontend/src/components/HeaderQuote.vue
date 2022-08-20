<template>
  <div class="rectangle">
    <p class="quote">Твой <b>голос</b> в руководстве ВУЗа</p>
  </div>
  <router-link to="/create-petition">
  <va-sidebar-item>
    <va-sidebar-item-content>
      <va-sidebar-item-title class="create-petition-text" style="text-align: center;">
        СОЗДАТЬ ПЕТИЦИЮ
        </va-sidebar-item-title>
    </va-sidebar-item-content>
  </va-sidebar-item>
  </router-link>
  <div class="line">
    <va-divider />
  </div>

</template>

<script>
export default {
  name: "HeaderQuote"
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@200;900&display=swap');
.line{
  width: 100%;
  height: 2px;
  background: #000000;
  margin-top: -10px;
  margin-bottom: 40px;
}
.quote{
  font-size: xx-large;
  font-family: Inter, cursive;
  color: #112e51;
  font-style: oblique;

}
.rectangle{
  background-color: #4da843;
  width: 100%;
  height: 50px;
  margin-bottom: 0;
  display: flex;
  justify-content:center;
  align-items: center;
}
.create-petition-text{
  text-transform: Uppercase;
  position: relative;
  left: 0px;
  font-size: x-large;
  font-weight: bold;
  margin-top: 10px;
  margin-bottom: 10px;
  color: #0a5d00;
}
</style>