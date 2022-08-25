<template>
<!--  <title>{{petition.title}}</title>-->
  <Header/>
  <span v-if="success === true" >
    <va-alert color="success" class="mb-4">
      Петиция успешно подписана!
    </va-alert>
  </span>
  <span v-if="failure === true" >
    <va-alert color="danger" class="mb-4">
      Не удалось подписать петицию. Проверьте введенные данные. Если это не поможет, обратитесь по адресу andrey.tamplon@gmail.com
    </va-alert>
  </span>
  <span v-if="alreadySigned === true" >
    <va-alert color="info" class="mb-4">
      Вы уже подписали петицию
    </va-alert>
  </span>

  <div v-if="dataReceived === true" class="petition-content">
    <div class="petition-text">
      <h1 class="petition-title">{{petition.title}}</h1>
      <h3 class="petition-creation-date">{{petition.dateOfCreation}}</h3>
      <div class="petition-tags">
        <va-chip
            v-for="(tag, index) in petition.tags"
            :key="index"
            outline
            class="mr-4 mb-2"
        >
          {{tag}}
        </va-chip>
      </div>
      <p class="petition-text-content pre-formatted">{{petition.petitionContent}}</p>
      <span class="petition-author">Автор: {{petition.authorName}}</span>
      <br><br>
      <span class="petition-date-of-expiration">Дата заверешнения: {{petition.dateOfExpiration}}</span>
    </div>
    <div class="petition-progress-bar">
      <va-progress-bar
          class="mt-2 progress-bar"
          :model-value="value"
          size="1.5rem"
          :color="color"
          content-inside
          show-percent
      />
    </div>
    <div class="petition-signed-info">
      <span class="mr-4">{{petition.numberOfSignatures}}/{{petition.signaturesTarget}}</span>
    </div>
    <va-button size="large" v-if="isLoggedIn === true && alreadySigned === false" @click="signPetition" class="sign-petition-button"> Подписать </va-button>
  </div>
  <div v-else>
    <div class="waiting-circle">
      <va-progress-circle size="250px" indeterminate />
    </div>
  </div>
  <div v-if="sessionEnded === true">
    <va-modal
        v-model="sessionEnded"
        hide-default-actions
    >
      <div>Ваша сессия закончилась, для продолжения необходимо войти в аккаунт</div>
      <template #footer>
        <router-link :to="{name: 'sign-in'}">
          <va-button>
            Войти
          </va-button>
        </router-link>
        <span style="margin-left: 3%">
          <va-button @click="sessionEnded = !sessionEnded">
            Закрыть
          </va-button>
        </span>
      </template>
    </va-modal>
  </div>
  <br><br>
  <Footer></Footer>
</template>

<script>
import axios from "axios";
import {BASE_CONTENT_URL} from "@/main";
import Header from "@/components/Header";
import Footer from "@/components/Footer";

export default {
  components: {
    Header,
    Footer,
  },
  name: "PetitionPage",
  data() {
    return {
      petition: null,
      sessionEnded: false,
      success: false,
      failure: false,
      alreadySigned: false,
      dataReceived: false,
    };
  },
  computed: {
    isLoggedIn() {
      return localStorage.getItem("token") !== null;
    },
    value() {
      return this.petition.numberOfSignatures / this.petition.signaturesTarget * 100;
    },
    color() {
      if (this.value < 50) {
        return "danger";
      } else if (this.value < 100) {
        return "warning";
      } else {
        return "success";
      }
    },
  },
  created() {
    this.getPetitionById(this.$route.params.id);
  },
  methods: {
    getPetitionById(id){
      axios.get(BASE_CONTENT_URL + '/petitions', {
        headers: {
          idFrom: id,
          idTo: id,
        }
      }).then(response => {
        console.log(response);
        this.petition = response.data[0];
        this.petition.tags = this.petition.tags.filter(tag => tag !== "");
        this.dataReceived = true;
      }, error => {
        console.log(error)
        this.dataReceived = false;
      })
    },
    signPetition() {
      axios.post(BASE_CONTENT_URL + "api/sign-petition", {
        petitionId: this.petition.id,
        userEmail: localStorage.getItem("email"),
      }, {
        headers: {
          token: localStorage.getItem("token"),
        },
      })
          .then((response) => {
            console.log(response);
            this.success = true;
            this.failure = false;
          })
          .catch((error) => {
            if (error.response.status === 409) {
              this.alreadySigned = true;
              this.success = false;
              this.failure = false;
            }
            else {
              this.failure = true;
              this.success= false;
            }
            console.log(error);
          });
    },
  }
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Ibarra+Real+Nova:wght@600&display=swap');
@import url('https://fonts.googleapis.com/css2?family=Ibarra+Real+Nova&display=swap');
.petition-content {
  width: 70%;
  margin-left: 15%
}
.petition-title{
  font-size: 2.5rem;
  font-weight: bold;
  font-family: 'Ibarra Real Nova', serif;
  text-align: center;
  margin-bottom: 1rem;
  margin-top: 2%;
}

.waiting-circle{
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.petition-tags{
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  margin-bottom: 1rem;
}
.petition-text-content{
  margin-bottom: 2rem;
  font-family: 'Ibarra Real Nova', serif;
  font-size: 1.5rem;
}
.pre-formatted{
  white-space: pre-wrap;
}
.petition-creation-date{
  font-size: 1.2rem;
  font-family: 'Ibarra Real Nova', serif;
  text-align: center;
  margin-bottom: 1rem;
}
.petition-progress-bar{
  margin-top: 1rem;
  margin-bottom: 2rem;
}
.petition-signed-info{
  text-align: center;
  margin-top: -2.5%;
  margin-left: 2%;
  font-size: 18px;
  margin-bottom: 1rem;
}

.sign-petition-button{
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  margin: 0 auto;
}
.petition-author{
  font-size: 1rem;
  font-family: 'Ibarra Real Nova', serif;
  text-align: center;
  margin-bottom: 2%;
}

.petition-date-of-expiration{
  font-size: 1rem;
  font-family: 'Ibarra Real Nova', serif;
  text-align: center;
  margin-top: 3rem;
}
</style>