<template>
    <component v-if="component" :is="component"/>
</template>

<script lang="ts" setup>
import { createVNode } from 'vue'
import { inputTypes, type InputAttr } from './helper';
import T10001 from './components/T10001.vue';
import T10002 from './components/T10002.vue';


const props = defineProps({
    type: {
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
        type: Object as PropType<QuestionInput>,
        default: () => {},
    },
    errors: {
        type: Object as PropType<QuestionInput>,
        default: () => {},
    },
})

const buildFun = new Map([
    [   
        'T10001',
        () => {
            return () =>
                createVNode(T10001, {
                    class: 'T1001',
                    ...props.attr,
                })
        },
    ],
    [   
        'T10002',
        () => {
            return () =>
                createVNode(T10002, {
                    class: 'T10002',
                    ...props.attr,
                })
        },
    ],
])

const action = buildFun.get(props.type)
const component =action ? (action!.call(this)) : null;

</script>