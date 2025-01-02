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
        <tr v-for="(item, i) in data?.questions" :key="i">
          <td>{{ item._id }}</td>
          <td>{{ item.title }}</td>
          <td>{{ item.description }}</td>
        </tr>
      </tbody>
    </table>

    <h2>Upload</h2>
    <!-- <FileUploader /> -->
  </div>
</template>

<script lang="ts" setup>
const queryQuestion = gql`
  query Questions {
    questions {
      _id
      title
      description
    }
  }
`;

const { data } = await useAsyncQuery(queryQuestion);
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
