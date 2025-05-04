<template>
    <div class="auth" :class="{ 'auth-leaving': isLeaving }">
      <h1>Регистрация</h1>
      <form @submit.prevent="submitRegister" class="auth_inner">
        <input v-model="email" type="email" placeholder="Email" required />
        <input v-model="password" type="password" placeholder="Пароль" required />
        <input v-model="name" type="text" placeholder="Имя" required />
        <div class="button_section">
            <button type="submit" class="register_button">Зарегистрироваться</button>
            <p>Есть аккаунт? 
                <a href="#" @click.prevent="goToLogin" class="to_login_link">Войти</a>
            </p>
        </div>
      </form>
      
    </div>
  </template>
  
  <script lang="ts" setup>
  import { ref } from 'vue'
  import axios from 'axios'
  
  const email = ref<string>('')
  const password = ref<string>('')
  const name = ref<string>('')
  const isLeaving = ref(false)
    import { useRouter } from 'vue-router'
    const router = useRouter()
  async function submitRegister(): Promise<void> {
    try {
      await axios.post('/api/register', { email: email.value, password: password.value, name: name.value })
      alert('Успешная регистрация!')
    } catch (err) {
      alert('Ошибка регистрации')
    }
  }
  function goToLogin(): void {
  isLeaving.value = true
  setTimeout(() => {
    router.push('/login')
  }, 500) 
  }
  </script>

  <style>
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
        cursor: pointer;
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