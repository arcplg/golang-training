<template>
    <div class="card" :id="props.data._id.toString()">
        <span>Question {{ index }}: </span>
        <div class="card-title">
            <h2><input type="text" name="A" v-model="props.data.text"></h2>
            <div>{{ props.data.note }}</div>
        </div>
        <div class="card-inner">
            <div class="" v-for="item,key in options" :key="key">
                <input type="radio">
                <label>{{ item.label }}</label>
                <input type="text" v-model="item.text">
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>

const props = defineProps({
    index: {
        type: Number,
        default: 1
    },
    data: {
        type: Object as PropType<Question>,
        required: true,
    }
})

const options = ref(props.data.options)

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
      flex-direction: column;
      justify-content: flex-start;
     }
  }
</style>