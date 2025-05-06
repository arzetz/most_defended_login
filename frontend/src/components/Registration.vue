<template>
  <DbButton @update:is-connected="val => isConnected = val">{{ isConnected ? 'Выйти из бд' : 'Войти в бд' }}</DbButton>
    <div class="auth" :class="{ 'auth-leaving': isLeaving, 'page-enter': isEntered }">
      <h1>Регистрация</h1>
      <form @submit.prevent="submitRegister" class="auth_inner">
        <input v-model="email" type="email" placeholder="Email" required />
        <input v-model="password" type="password" placeholder="Пароль" required />
        <input v-model="phone" type="tel" placeholder="Телефон" required />
        <div class="button_section">
            <button type="submit" class="register_button">Зарегистрироваться</button>
            <p style="width: 130px;">Есть аккаунт? 
                <a href="#" @click.prevent="goToLogin" class="to_login_link">Войти</a>
            </p>
        </div>
      </form>
    </div>
    <div :class="{ 'auth-failed': registerFailed }">
      {{thingFailed}}
    </div>
  </template>
  
  <script lang="ts" setup>
  import { ref, onMounted  } from 'vue'
  import axios from 'axios'
  import DbButton from './DbButton.vue'

  const email = ref<string>('')
  const password = ref<string>('')
  const phone = ref<string>('')
  const thingFailed = ref<string>('')
  const registerFailed = ref(false)
  const isEntered = ref(true)
  const isLeaving = ref(false)
  const isConnected = ref(false)
  import { useRouter } from 'vue-router'
  const router = useRouter()
  async function submitRegister(): Promise<void> {
    if (!/^[^\s@]+@[^\s@]+\.[a-zA-Z]{2,}$/.test(email.value)){
      registerFailed.value = true
      thingFailed.value = "Email гавно"
      setTimeout(() => {
        registerFailed.value = false
        thingFailed.value = ""
  }, 5000)                                                                           
      return
    }

    if (!/^\+7[\d]{10}$/.test(phone.value)){
      registerFailed.value = true
      thingFailed.value = "Телефон гавно"
      setTimeout(() => {
        registerFailed.value = false
        thingFailed.value = ""
  }, 5000)                                                                           
      return
    }

    try {
      await axios.post('/api/register', { email: email.value, password: password.value, phone: phone.value })
      alert('Успешная регистрация!')
    } catch (err) {
      alert('Ошибка регистрации')
    }
  }

  onMounted(() => {
  setTimeout(() => {
    isEntered.value = false
  }, 1)
})

  function goToLogin(): void {
  isLeaving.value = true
  
  setTimeout(() => {
   
    router.push('/login')
  }, 500) 
  }
  </script>

  <style>
  .reg_p{
    margin-top: 10px;
  }
  .auth{
        text-align: center;
        max-width: 500px;
        display: flex;
        flex-direction: column;
        justify-content: center;
        margin-left: auto;
        margin-right: auto;
        margin-top: 200px;
        transition: 
        opacity 0.5s ease,
        transform 0.5s ease,
        filter 0.5s ease;
        opacity: 1;
        transform: translateY(0);
  filter: blur(0);
    }
    .auth-leaving {
  opacity: 0;
  transform: translateY(10px); 
  filter: blur(2px); 
}
.page-enter{
  opacity: 0;
  transform: translateY(10px); 
  filter: blur(2px); 
}
    .auth_inner{
        display: flex;
        flex-direction: column;
    }
    .button_section{
        justify-content: space-between;
        display: flex;
        flex-direction: row;
    }
    .register_button{
      width: 140px;
      height: 40px;
      cursor: pointer;
      color: rgb(38, 38, 38);
    }
    .to_login_link{
      text-decoration: none;
      color: rgb(116, 116, 116);
      transition: 0.5s;
      display: inline-flex;
  flex-direction: column;
  align-items: left;
    }
    .to_login_link:hover{
      text-decoration: none;
      color: rgb(194, 194, 194);
    }
    .to_login_link::after {
  content: '';  
  width: 0%;
  height: 2px; 
  background-color: azure;
  transition: width 0.5s ease;
}
.to_login_link:hover::after{
  width: 100%;
}
  </style>