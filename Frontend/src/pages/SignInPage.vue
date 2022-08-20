<template>
  <title>Войти</title>
  <span v-if="unknownError === true">
    <va-alert color="danger" class="mb-4">
      Произошла неизвестная ошибка :(.
    </va-alert>
  </span>
  <span v-if="badRequestError === true">
    <va-alert color="danger" class="mb-4">
      Введенные данные некорректны.
    </va-alert>
  </span>
  <span v-if="noAccountError === true">
    <va-alert color="danger" class="mb-4">
      Пользователя с такими данными не существует. Проверьте правильность введенной почты и пароля.
    </va-alert>
  </span>
  <div >
    <span ><router-link to="/"><va-button icon="home" class="mr-4 mb-2" /></router-link></span>
    <div class="login-box">
      <div class="login-content">
        <h1 class="login-text">Авторизация</h1>
        <div class="login-forms">
          <va-form
              style="width: 300px;"
              tag="form"
          >
            <va-input
                class="mt-3"
                label="Электронная почта"
                v-model="email"
                :rules="[emailRule  || `Электронная почта должна быть из НГУ`]"
            />
            <va-input
                class="mt-3"
                :type="passwordFieldType"
                label="Пароль"
                v-model="password"
                :rules="[passwordRule || 'Ваш пароль длиннее']"
            />
          </va-form>
        </div>
        <va-button color="#838383" size="small" class="mr-4 mb-2 lg-show-password-button" @click="switchVisibility">Показать пароль</va-button>
        <div class="login-button">
          <va-button v-if="allCorrect === false" class="mr-4 mb-2" disabled>Войти</va-button>
          <va-button v-if="allCorrect === true" class="mr-4 mb-2" @click="login">Войти</va-button>
        </div>
        <router-link to="/sign-up"><va-button size="small" color="warning" :rounded="false" class="mr-4 mb-2 no-account">У меня нет аккаунта</va-button></router-link>
      </div>

    </div>
  </div>
</template>

<script>
import axios from "axios";
import {BASE_AUTH_URL} from "@/main";

export default {
  name: "SignInPage",
  data() {
    return {
      email: "",
      password: "",
      passwordFieldType: "password",
      noAccountError: false,
      unknownError: false,
      badRequestError: false,
    };
  },
  computed: {
    // eslint-disable-next-line vue/no-dupe-keys
    emailRule() {
      return (
          (this.email.indexOf("@g.nsu.ru") !== -1)
          || (this.email.indexOf("@nsu.ru") !== -1)
          || (this.email.indexOf("@alumni.nsu") !== -1)
      );
    },
    // eslint-disable-next-line vue/no-dupe-keys
    passwordRule() {
      return this.password.length >= 6;
    },
    // eslint-disable-next-line vue/no-dupe-keys
    allCorrect() {
      return this.emailRule && this.passwordRule;
    },

  },

  methods: {
    switchVisibility() {
      this.passwordFieldType = this.passwordFieldType === "password" ? "text" : "password";
    },
    login() {
      axios.post(BASE_AUTH_URL + "sign-in", {
        email: this.email,
        password: this.password,
      }).then((response) => {
        if (response.status === 200) {
          console.log(response);
          localStorage.setItem("token", response.data.token);
          localStorage.setItem("name", response.data.name);
          localStorage.setItem("email", this.email);
          this.unknownError = false;
          this.$router.push("/");
        }
      }
      ).catch((error) => {
        if (error.response.status === 400) {
          this.badRequestError = true;
        } else if (error.response.status === 401) {
          this.noAccountError = true;
        } else {
          this.unknownError = true;
        }
        console.log(error);
      });
    }
  }
}
</script>

<style scoped>
.login-box {
  position: relative;
  margin: 5% auto;
  width: 500px;
  height: 370px;
  border-radius: 2px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.4);
}
.login-content {
  position: relative;
  left: 16.5%;
  box-sizing: border-box;
  padding: 40px;
  width: 500px;
  height: 350px;
}
.login-forms{
  position: relative;
  width: 400px;
  margin-left: -6%;
  margin-bottom: 5%;
  border-radius: 2px;
}
.lg-show-password-button {
  position: relative;
  left: 11%;
}
.login-button {
  margin-left: 17%;
  margin-top: 3%;
}
.login-text{
  font-family: 'Rubik', sans-serif;
  font-style: normal;
  margin-left: 8%;
  font-weight: normal;
  font-size: 28px;
}
.no-account{
  position: relative;
  left: 9%;
  margin-top: 1.5%;
}
</style>