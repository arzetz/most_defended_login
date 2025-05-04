<template>
  <router-view />
</template>

<style>
  html{    
    background-color: rgb(32, 32, 32);
    color: rgb(224, 224, 224);
    font-family: Arial, Helvetica, sans-serif;
  }
  input {
  margin-bottom: 20px;
  height: 50px;
  border-radius: 4px;
  border: 0px solid transparent; /* прозрачная граница по умолчанию */
  outline: none;
  transition: border 0.3s ease, box-shadow 0.3s ease;
}

input:focus {
  border: 0px solid #929594; /* нужный цвет границы */
  box-shadow: 0 0 15px rgba(255, 255, 255, 0.5);
}

  p{
    margin: 0px;

  }
</style>