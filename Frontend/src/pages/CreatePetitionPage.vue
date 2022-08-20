<template>
  <title>Создать петицию</title>
  <span v-if="creationError === true" >
    <va-alert color="danger" class="mb-4">
      Не удалось создать петицию. Проверьте введенные данные. Если это не поможет, обратитесь по адресу andrey.tamplon@gmail.com
    </va-alert>
  </span>
  <span v-if="success === true" >
    <va-alert color="success" class="mb-4">
      Петиция успешно создана!
    </va-alert>
  </span>
  <span style="display: inline-block"><router-link to="/"><va-button icon="home" class="mr-4 mb-2" /></router-link></span>
  <div class="create-petition-box">
    <div class="create-petition-content" v-if="sessionEnded === false">
      <h1 class="create-petition-text">Создать петицию</h1>
      <va-form
          style="width: 100%;"
          tag="form"
      >
        <va-input
            style="margin-top: 5%"
            v-model="title"
            :max-length="70"
            :maxlength="70"
            :rules="[titleRule || 'Название не может быть пустым']"
            counter
            class="mb-4"
            placeholder="Название петиции"

        >
        </va-input>
        <span>
        <va-input
            style="margin-top: 5%; margin-right: 4%; width: 30%; display: inline-block"
            v-model="tags[0]"
            :max-length="40"
            :maxlength="40"
            counter
            class="mb-4"
            placeholder="Тег"

        >
        </va-input>
        <va-input
            style="margin-top: 5%; margin-right: 4%; width: 30%; display: inline-block"
            v-model="tags[1]"
            :max-length="40"
            :maxlength="40"
            counter
            class="mb-4"
            placeholder="Тег"
        >
        </va-input>
        <va-input
            style="margin-top: 5%; width: 30%; display: inline-block"
            v-model="tags[2]"
            :max-length="40"
            :maxlength="40"
            counter
            class="mb-4"
            placeholder="Тег"
        >
        </va-input>
        </span>
        <va-input
            style="margin-top: 5%"
            v-model="petitionContent"
            :max-length="5000"
            :maxlength="5000"
            :rules="[petitionContentRule  || 'Текст петиции не может быть пустым']"
            type="textarea"
            autosize
            :min-rows="8"
            :max-rows="17"
            counter
            class="mb-4"
            placeholder="Текст петиции"
        >
        </va-input>
      </va-form>
      <div class="sign-up-button" style="margin-left: 43%">
        <va-button v-if="allCorrect === false" class="mr-4 mb-2" disabled>Создать</va-button>
        <va-button v-if="allCorrect === true" class="mr-4 mb-2" @click="submitPetition">Создать</va-button>
      </div>
    </div>
    <div v-else>
      <va-modal
          v-model="sessionEnded"
          hide-default-actions
      >
        <div>Ваша сессия закончилась, для продолжения необходимо войти в аккаунт</div>
        <template #footer>
          <router-link :to="{name: 'sign-in'}">
            <va-button>
              Понятно
            </va-button>
          </router-link>
        </template>
      </va-modal>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import {BASE_CONTENT_URL} from "@/main";
import {mapGetters} from "vuex";

export default {
  name: "CreatePetitionPage",
  data() {
    return {
      title: "",
      tags: ["", "", ""],
      petitionContent: "",
      sessionEnded: false,
      creationError: false,
      success: false,
    };
  },
  computed: {
    ...mapGetters([
      "getToken",
    ]),
    titleRule() {
      return this.title.length > 0;
    },
    petitionContentRule() {
      return this.petitionContent.length > 0;
    },
    allCorrect() {
      return this.titleRule && this.petitionContentRule;
    },
  },
  methods: {
      // Later there will be an opportunity to add expiration dateOfCreation and signatures target.
      submitPetition() {
        let dateOfCreation = new Date();
        let dateOfExpiration = new Date(dateOfCreation);
        dateOfExpiration.setDate(dateOfExpiration.getDate() + 180);
        let signaturesTarget = 100;
        let authorName = localStorage.getItem("name");
        let authorEmail = localStorage.getItem("email");
        axios.post(BASE_CONTENT_URL + "api/create-petition", {
          title: this.title,
          authorName: authorName,
          authorEmail: authorEmail,
          dateOfCreation: dateOfCreation.toLocaleDateString(),
          dateOfExpiration: dateOfExpiration.toLocaleDateString(),
          tags: this.tags,
          petitionContent: this.petitionContent,
          signaturesTarget: signaturesTarget,
        }, {
          headers: {
            token: this.getToken
          }
        }).then(response => {
          this.title = "";
          this.tags = ["", "", ""];
          this.petitionContent = "";
          this.success = true;
          this.creationError = false;
          console.log(response);
        }).catch(error => {
          if (error.status === 401) {
            this.sessionEnded = true;
          }
          else {
            this.success = false;
            this.creationError=true;
          }
          console.log(error);
        });
      }
  },
}
</script>

<style scoped>

.create-petition-box{
  position: relative;
  margin: 5% auto;
  width: 70%;
  height: auto;
  border-radius: 2px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.4);
}
.create-petition-content{
  position: relative;
  left: 16.5%;
  box-sizing: border-box;
  padding: 40px;
  width: 70%;
  height: auto;
}

.create-petition-text{
  font-family: 'Rubik', sans-serif;
  text-align: center;
  font-style: normal;
  margin-left: 1.5%;
  font-weight: normal;
  font-size: 28px;
}
</style>