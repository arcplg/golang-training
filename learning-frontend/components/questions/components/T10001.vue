<template>
    <div class="card">
        <span>Question {{props.data.originNumber}}: </span>
        <div class="card-title">
            <h2><input type="text" name="A" v-model="props.data.text"></h2>
        </div>
        <div class="card-inner">
            <div class="" v-for="item,key in questionItems" :key="key">
                <input type="radio">
                <label>{{item.optionKey}}</label>
                <input type="text" v-model="item.text">
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>

const props = defineProps({
    data: {
        type: Object as PropType<QuestionInput>,
        required: true,
    }
})

const questionItems = ref(props.data.questionItems)

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