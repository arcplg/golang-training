<template>
    <div class="card">
        <span>Question 2: </span>
        <div class="card-title">
            <h2><input type="text" name="A"></h2>
        </div>
        <div class="card-inner">
            <div class="question">
            <input type="radio">
            <label>A)</label>
            <input type="text" name="A">
            </div>
            <div class="question">
            <input type="radio">
            <label>B)</label>
            <input type="text" name="B">
            </div>
            <div class="question">
            <input type="radio">
            <label>C)</label>
            <input type="text" name="C">
            </div>
            <div class="question">
            <input type="radio">
            <label>D)</label>
            <input type="text" name="D">
            </div>
        </div>
    </div>
</template>
<script lang="ts" setup>

</script>

<style lang="scss">
  .card {
    margin-top: 15px;
    background-color: #ccc;
    padding: 15px;
    width: 100%;
     &-inner {
      display: flex;
      width: 100%;
      padding: 15px;
      background-color: #FFF;
      .question {
        width: 50%;
        display: flex;
      }
     }
     &-title {
      display: flex;
     }
  }
</style>