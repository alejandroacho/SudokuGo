<template>
  <div id="app"> 
    <img alt="SudokGo logo" src="./assets/images/comba.gif" class="logo">
    <Sudoku/>
  </div>
</template>

<script>
import Sudoku from "./components/Sudoku";
import "./assets/css/main.css";

export default {
  name: "app",
  components: {
    Sudoku
  }
};
</script>
