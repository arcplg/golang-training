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
        <tr v-for="(item, i) in pageQuestion.questions" :key="i">
          <td>{{ item._id }}</td>
          <td>{{ item.title }}</td>
          <td>{{ item.description }}</td>
        </tr>
      </tbody>
    </table>

    <form>
      <div>
        Title
        <input type="text" name="title" v-model="pageQuestion.form.title" />
      </div>
      <div>
        Descriptions
        <textarea name="description" v-model="pageQuestion.form.description" />
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

  /** variable */
  const query = gql`
    query Questions {
      questions {
        _id
        title
        description
      }
    }
  `
 const mutation = gql`
      mutation CreateQuestion($title: String!, $description: String! ) {
        createQuestion(input: { title: $title, description: $description }) {
            _id
            title
            description
        }
      }
      
  `;

  interface PageQuestion {
    questions: Question[],
    form: NewQuestion
  }
  const pageQuestion = ref<PageQuestion>({
    questions: [],
    form: {
      title: "",
      description: "",
    }
  })

  /** list question */
  const { data } = await graphqlQueryUseFetch(query)
  pageQuestion.value.questions  = data?.questions || [];

  /** Make new question */
  const submit = async () => {
    await graphqlQueryFetch(mutation, pageQuestion.value.form);
    const { data } = await graphqlQueryFetch(query)
    pageQuestion.value.questions  = data?.questions || [];
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
