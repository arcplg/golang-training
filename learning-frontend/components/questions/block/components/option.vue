<template>
    <div>
        {{ block.text }}
        <div class="" v-for="block,key in block.blocks" :key="key">
            <input type="text" :name="block.label.toString()" v-model="block.text">
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
        type: Object as PropType<Block>,
        required: true,
    }
})

const block = ref(props.data)

</script>

<style lang="scss" scoped>
</style>