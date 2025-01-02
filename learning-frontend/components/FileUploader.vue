<template>
  <form @submit.prevent="handleFileUpload">
    <input type="file" multiple @change="handleFileChange" />
    <button type="submit">Upload</button>
    <p v-if="uploadMessage">{{ uploadMessage }}</p>
  </form>
</template>

<script setup>
import { ref } from "vue";

// const mutation = gql`
//     mutation CreateQuestion {
//         createQuestion(input: { title: null, description: null }) {
//             _id
//             title
//             description
//         }
//     }

// `

const selectedFiles = ref(null);
const uploadMessage = ref("");

const handleFileChange = (event) => {
  selectedFiles.value = event.target.files;
};

const handleFileUpload = async () => {
  // const { mutate } = useMutation(mutation, { title: '111112222', description: 11222 })
  // const response = await mutate()
  // console.log(response)
  // if (!selectedFiles.value || selectedFiles.value.length === 0) {
  //   uploadMessage.value = 'Please select files to upload';
  //   return;
  // }
  // for (const file of selectedFiles.value) {
  //     const { mutate } = useMutation(mutation, { file })
  //     console.log(mutate)
  // }
};
</script>
