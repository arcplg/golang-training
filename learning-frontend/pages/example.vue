<template>
  <div>
    <h1>Questions</h1>
    <table>
      <tbody>
        <tr>
          <td>Id</td>
          <td>Title</td>
          <td>Body</td>
        </tr>
        <tr v-for="(item, i) in questions" :key="i">
          <td>{{ item._id }}</td>
          <td>{{ item.title }}</td>
          <td>{{ item.description }}</td>
        </tr>
      </tbody>
    </table>

    <form>
      <div>
        Title
        <input type="text" name="title" v-model="formQuestion.title" />
      </div>
      <div>
        Descriptions
        <textarea name="description" v-model="formQuestion.description" />
      </div>
      <div>
        <button @click.prevent="submit">Submit</button>
      </div>
    </form>

    <h2>Upload</h2>
    <!-- <FileUploader /> -->
  </div>
</template>

<script lang="ts" setup>
import gql from "graphql-tag"

const query = gql`
  query Questions {
    questions {
      _id
      title
      description
    }
  }
`

console.log(query)

interface QueryResponse {
  questions: Question[]
}

const data = await $fetch("http://localhost:8080/graphql", {
  method: "POST",
  headers: {
    "Content-Type": "application/json",
    Accept: "application/json",
  },
  body: {
    query: query?.loc?.source.body,
  },
})

// const { data } = await useQuery(queryQuestion);
// console.log(data)
// const questions: Question[]  = data.value?.questions || [];

// const queryCreateQuestion = gql`
//     mutation CreateQuestion($title: string!, $description: string! ) {
//       createQuestion(input: { title: $title, description: $description }) {
//           _id
//           title
//           description
//       }
//   }
// `;

const formQuestion = ref({
  title: "",
  description: "",
})

const submit = async () => {
  // const value: any = formQuestion.value
  // console.log(value);
  // const { mutate: createQuestion } = await useMutation<QueryResponse>(queryCreateQuestion);
  // const response = await createQuestion(value)
  // console.log(response);
}
</script>

<style>
table {
  width: 100%;
  border-collapse: collapse;
  margin: 20px 0;
  text-align: left;
  table-layout: fixed;
}

th,
td {
  padding: 10px;
  border: 1px solid #ddd;
}

th {
  background-color: #f4f4f4;
  color: #333;
}
td {
  width: 200px;
  word-wrap: break-word;
}
</style>
