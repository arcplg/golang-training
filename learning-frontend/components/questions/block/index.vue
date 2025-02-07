<template>
    <component v-if="component" :is="component"/>
</template>

<script lang="ts" setup>
import { createVNode } from 'vue'
import { inputTypes, type InputAttr } from './helperBlock';
import Title from './components/title.vue';
import Option from './components/option.vue';

const props = defineProps({
    index: {
        type: Number,
        default: 0,
    },
    label: {
        type: String,
        required: true,
        validator: (value: string) => {
            return inputTypes.includes(value)
        },
    },
    attr: {
        type: Object as PropType<InputAttr>,
        default: () => {},
    },
    data: {
        type: Object as PropType<Block>,
        default: () => {},
    },
    errors: {
        type: Object as PropType<Block>,
        default: () => {},
    },
})

const buildFun = new Map([
    [   
        'title',
        () => {
            return () =>
                createVNode(Title, {
                    class: 'title',
                    ...props,
                })
        },
    ],
    [   
        'option',
        () => {
            return () =>
                createVNode(Option, {
                    class: 'option',
                    ...props,
                })
        },
    ],
])

const action = buildFun.get(props.label)
const component =action ? (action!.call(this)) : null;

</script>