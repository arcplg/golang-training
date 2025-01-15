<template>
    <component v-if="component" :is="component"/>
</template>

<script lang="ts" setup>
import { createVNode } from 'vue'
import { inputTypes, type InputAttr } from './helper';
import Basic from './components/basic.vue';
import Template2 from './components/template2.vue';


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
        'basic',
        () => {
            return () =>
                createVNode(Basic, {
                    class: 'basic',
                    ...props.attr,
                })
        },
    ],
    [   
        'template2',
        () => {
            return () =>
                createVNode(Template2, {
                    class: 'template2',
                    ...props.attr,
                })
        },
    ],
])

const action = buildFun.get(props.type)
const component =action ? (action!.call(this)) : null;

</script>