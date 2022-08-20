<template>
  <title>Регистрация</title>
  <span v-if="alreadyRegisteredError === true">
    <va-alert color="danger" class="mb-4">
      Данный адрес уже зарегистрирован.
    </va-alert>
  </span>
  <span v-if="unknownErrorHappened === true">
    <va-alert color="danger" class="mb-4">
      Не удалось зарегистрировать. Проверьте введенные данные. Пароль должен быть от 6 символов и больше,
      допустимыми почтовыми адресами являются только адреса университетской почты.
    </va-alert>
  </span>
  <span v-if="confirmationError === true">
    <va-alert color="danger" class="mb-4">
      Произошла ошибка при подтверждении адреса.
    </va-alert>
  </span>
  <div v-if="registrationHappened === false && registrationConfirmation === false">
    <span ><router-link to="/"><va-button icon="home" class="mr-4 mb-2" /></router-link></span>
    <div  id="sign-up-box">
      <div class="sign-up-input-box">
        <h1 class="sign-up-text">Регистрация</h1>
        <div class="sign-up-forms">
          <va-form
              style="width: 300px;"
              tag="form"
          >
            <va-input
                class="mt-3"
                label="Имя"
                v-model="userName"
                :rules="[nameRule || 'Введите имя']"
            />
            <va-input
                class="mt-3"
                label="Электронная почта"
                v-model="email"
                :rules="[emailRule  || `Используйте почту из НГУ`]"
            />
            <va-input
                class="mt-3"
                :type="passwordFieldType"
                label="Пароль"
                v-model="password"
                :rules="[passwordRule || 'Пароль должен быть от 6 символов и больше']"
            />
            <va-input
                class="mt-3"
                :type="passwordFieldType"
                label="Повторите пароль"
                v-model="passwordConfirmation"
                :rules="[passwordConfirmationRule || 'Пароли не совпадают']"
            />
          </va-form>
        </div>
        <va-button color="#838383" size="small" class="mr-4 mb-2 show-password-button" @click="switchVisibility">Показать пароль</va-button>
        <div class="sign-up-button">
          <va-button v-if="allCorrect === false" class="mr-4 mb-2" disabled>Зарегистрироваться</va-button>
          <va-button v-if="allCorrect === true" class="mr-4 mb-2" @click="register">Зарегистрироваться</va-button>
        </div>
        <router-link to="/sign-in"><va-button style="margin-left: 6%; margin-top: 2%" size="small" color="success" :rounded="false" class="mr-4 mb-2 no-account">У меня уже есть аккаунт</va-button></router-link>
      </div>
    </div>
  </div>
  <div v-if="registrationConfirmation === true">
    <div id="sign-up-box">
      <div class="sign-up-input-box">
        <h1 class="sign-up-text">Регистрация</h1>
        <h4 class="explanatory-text">Вам на почту был отправлен код. Для завершения регистрации вам необходимо ввести его здесь</h4>
        <va-input
            class="mb-4 cc-input-form"
            v-model="confirmationCode"
            :rules="[(v) => v.length === 6 || `Не совпадает длина кода`]"
            label="Код"
            placeholder="Введите код подтверждения"
        />
        <va-button @click="confirmRegistration" type="submit" class="confirmation-code-button">
          Подтвердить
        </va-button>
      </div>
    </div>
  </div>
  <div v-if="registrationHappened === true">
    <va-alert color="success" class="mb-4">
      Вы успешно зарегестрировались.
    </va-alert>
    <div class="navigation-buttons">
      <span><router-link :to="{name:'home'}"><va-button color="success" size="large" class="mr-4 mb-2">Вернуться на главную</va-button></router-link></span>
      <span><router-link :to="{name:'sign-in'}"><va-button color="success" size="large" class="mr-4 mb-2">Войти</va-button></router-link></span>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import {BASE_AUTH_URL} from "@/main";
export default {
  name: "SignUpPage",
  data() {
    return {
      userName: "",
      email: "",
      password: "",
      passwordConfirmation: "",
      passwordFieldType: "password",
      registrationHappened: false,
      alreadyRegisteredError: false,
      unknownErrorHappened: false,
      registrationConfirmation: false,
      confirmationCode: "",
      confirmationError: false,
      savedEmail: "",
    };
  },
  computed: {
    nameRule() {
      return this.userName.length >= 1;
    },
    emailRule() {
      return (
        (this.email.indexOf("@g.nsu.ru") !== -1)
        || (this.email.indexOf("@nsu.ru") !== -1)
        || (this.email.indexOf("@alumni.nsu") !== -1)
      );
    },
    passwordRule() {
      return this.password.length >= 6;
    },
    passwordConfirmationRule() {
      return this.passwordConfirmation === this.password;
    },
    allCorrect() {
      return this.nameRule && this.emailRule && this.passwordRule && this.passwordConfirmationRule;
    },
  },
  methods: {
    switchVisibility() {
      this.passwordFieldType = this.passwordFieldType === "password" ? "text" : "password";
    },
    register() {
      if (this.userName && this.email && this.password && this.passwordConfirmation) {
        if (this.password === this.passwordConfirmation) {
          this.savedEmail = this.email;
          axios.post(BASE_AUTH_URL + 'sign-up', {
            name: this.userName,
            email: this.email,
            password: this.password,
          }).then(response => {
            console.log(response);
            if (response.status === 200) {
              this.registrationHappened = false;
              this.registrationConfirmation = true;
              this.unknownErrorHappened= false;
              this.alreadyRegisteredError = false;
            }
          }).catch(error => {
            if(error.response.status === 409) {
              this.alreadyRegisteredError = true;
              this.unknownErrorHappened= false;
            }
            else{
              this.unknownErrorHappened = true;
              this.alreadyRegisteredError = false;
            }
            console.log(error);
          });

        } else {
          alert("ПАРОЛИ НЕ СОВПАДАЮТ");
        }
      } else {
        alert("ЗАПОЛНЕНЫ НЕ ВСЕ ПОЛЯ");
      }
    },
    confirmRegistration() {
      axios.post(BASE_AUTH_URL + 'sign-up-confirmation', {
        email: this.email,
        code: this.confirmationCode,
      }).then(response => {
        console.log(response);
        if (response.status === 201) {
          this.registrationConfirmation = false;
          this.confirmationError = false;
          this.registrationHappened = true;
        }
      }).catch(error => {
        this.confirmationError = true;
        console.log(error);
      });
    },
  },
}
</script>

<style scoped>

@import url(https://fonts.googleapis.com/css?family=Roboto:400,300,500);
*:focus {
  outline: none;
}
html, body {
  margin: 0;
  height: 100vm;
  padding: 0;
  font-size: 16px;
  color: #222;
  font-family: 'Roboto', sans-serif;
  font-weight: 300;
}

#sign-up-box {
  position: relative;
  margin: 5% auto;
  width: 600px;
  height: 510px;
  border-radius: 2px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.4);
}

.sign-up-forms {
  position: relative;
  width: 400px;
  margin-left: -6%;
  margin-bottom: 5%;
  border-radius: 2px;
}

.cc-input-form {
  position: relative;
  margin-top: -5%;
  width: 250px;
  height: 40px;
  background: #fff;
  border-radius: 2px;
}
.navigation-buttons {
  position: relative;
  align-items: center;
  margin-top: 20%;
  margin-left: 38%;
}

.confirmation-code-button {
  margin-left: 10%;
}
.sign-up-button {
  margin-left: 5%;
  margin-top: 2%;
}
.sign-up-input-box {
  position: relative;
  left: 22.5%;
  box-sizing: border-box;
  padding: 40px;
  width: 500px;
  height: 400px;
}
.explanatory-text {
  width: 400px;
  margin-left: -15%;
  font-size: 18px;
  font-weight: normal;
  margin-top: 15%;
  margin-bottom: 10%;
}
.show-password-button {
  position: relative;
  left: 11%;
}

.sign-up-text {
  font-family: 'Rubik', sans-serif;
  font-style: normal;
  margin-left: 8%;
  font-weight: normal;
  font-size: 28px;
}
</style>