<template>
    <div class="auth" :class="{ 'auth-leaving': isLeaving, 'page-enter': isEntered  }">
      <h1>Вход</h1>
      <form @submit.prevent="submitLogin" class="auth_inner">
        <input v-model="email" type="email" placeholder="Email" required />
        <input v-model="password" type="password" placeholder="Пароль" required />
        <div class="button_section">
        <button type="submit" class="enter_button">Войти</button>
        <p style="width: 130px;">Нет аккаунта?
          <a href="#" @click.prevent="goToRegister" class="to_register_link">Регистрация</a>
        </p>
      </div>
      </form>
     
    </div>
  </template>
  
  <script lang="ts" setup>
  import { ref, onMounted } from 'vue'
  import axios from 'axios'
  import { useRouter } from 'vue-router'
  const router = useRouter()
  const email = ref<string>('')
  const password = ref<string>('')
  const isEntered = ref(true)
  const isLeaving = ref(false)
  async function submitLogin(): Promise<void> {
    try {
      await axios.post('/api/login', { email: email.value, password: password.value })
      alert('Успешный вход!')
    } catch (err) {
      alert('Ошибка входа')
    }
  }

  onMounted(() => {
  setTimeout(() => {
    isEntered.value = false
  }, 1)
})

  function goToRegister(): void {
  isLeaving.value = true
  setTimeout(() => {
    router.push('/register')
  }, 500) 
  }
  </script>

  <style>
    .auth{
      padding: 25px;
      border-radius: 20px;
      background-color: rgb(25, 25, 25);
        text-align: center;
        width: 300px;
        max-width: 500px;
        display: flex;
        flex-direction: column;
        justify-content: center;
        margin-left: auto;
        margin-right: auto;
        margin-top: 300px;
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
    .auth_inner{
        display: flex;
        flex-direction: column;
    }
    .button_section{
      margin-top: 10px;
      padding-bottom: 10px;
        justify-content: space-between;
        display: flex;
        flex-direction: column;
        
    }
    .enter_button{
      width: 100px;
      height: 40px;
      cursor: pointer;
      color: rgb(38, 38, 38);
    }
    .to_register_link{
      text-decoration: none;
      color: rgb(116, 116, 116);
      transition: 0.5s;
      display: inline-flex;
  flex-direction: column;
  align-items: left;
    }
    .to_register_link:hover{
      text-decoration: none;
      color: rgb(194, 194, 194);
    }
    .to_register_link::after {
  content: '';  
  width: 0%;
  height: 2px; 
  background-color: azure;
  transition: width 0.5s ease;
}
.to_register_link:hover::after{
  width: 100%;
}
  </style>