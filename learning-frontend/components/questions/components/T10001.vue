<template>
    <div class="card">
        <span>Question 1: </span>
        <div class="card-title">
            <h2><input type="text" name="A"></h2>
        </div>
        <div class="card-inner">
            <div class="">
            <input type="radio">
            <label>A)</label>
            <input type="text" name="A">
            </div>
            <div class="">
            <input type="radio">
            <label>B)</label>
            <input type="text" name="B">
            </div>
            <div class="">
            <input type="radio">
            <label>C)</label>
            <input type="text" name="C">
            </div>
            <div class="">
            <input type="radio">
            <label>D)</label>
            <input type="text" name="D">
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>



</script>

<style lang="scss" scoped>
  .card {
    margin-top: 15px;
    background-color: #ccc;
    padding: 15px;
    width: 100%;
    display: block;
    box-sizing: border-box;
     &-inner {
      display: flex;
      flex-direction: column;
      width: 100%;
      padding: 15px;
      background-color: #FFF;
      height: 200px;
      box-sizing: border-box;
     }
     &-title {
      display: flex;
     }
  }
</style>