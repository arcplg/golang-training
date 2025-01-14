<template>
  <div>

    <h1>Questions</h1>
    <form>
      <div>
        Title
        <input type="text" name="title" v-model="pageGroupQuestion.form.title" />
      </div>
      <div>
        Description
        <input name="description" v-model="pageGroupQuestion.form.description" />
      </div>
      <div>
        <button @click.prevent="submit">Submit</button>
      </div>
    </form>
    <table>
      <tbody>
        <tr>
          <td>Id</td>
          <td>Title</td>
          <td>Description</td>
          <td>Thumbnail</td>
          <td>Any Time</td>
          <td>Start At</td>
          <td>End At</td>
          <td>Created At</td>
          <td>Updated At</td>
          <td>Action</td>
        </tr>
        <tr v-for="(item, i) in pageGroupQuestion.groupQuestions" :key="i">
          <td>{{ item._id }}</td>
          <td>{{ item.title }}</td>
          <td>{{ item.description }}</td>
          <td>{{ item.thumbnailUrl }}</td>
          <td>{{ item.anyTime }}</td>
          <td>{{ item.startAt }}</td>
          <td>{{ item.endAt }}</td>
          <td>{{ item.createdAt }}</td>
          <td>{{ item.updatedAt }}</td>
          <td>
            <NuxtLink :to="'questions/'+item._id">
              Detail
            </NuxtLink>
          </td>
        </tr>
      </tbody>
    </table>

    
  </div>
</template>

<script lang="ts" setup>
import gql from "graphql-tag"

/** variable */
const query = gql`
  query GroupQuestions {
    groupQuestions {
      _id
      title
      description
      thumbnailUrl
      anyTime
      startAt
      endAt,
      createdAt,
      updatedAt
    }
  }
`
const mutation = gql`
  mutation CreateGroupQuestion(
      $title: String!,
      $description: String!,
      $thumbnailUrl: String
      $anyTime: Boolean
      $startAt: DateTime
      $endAt: DateTime
    ) {
      createGroupQuestion(input: { 
        title: $title, 
        description: $description,
        thumbnailUrl: $thumbnailUrl,
        anyTime: $anyTime,
        startAt: $startAt,
        endAt: $endAt,
      }) {
      _id
      title
      description
      thumbnailUrl
      anyTime
      startAt
      endAt
    }
  }
`

interface PageGroupQuestion {
  groupQuestions: GroupQuestion[]
  form: GroupQuestionInput
}
const pageGroupQuestion = ref<PageGroupQuestion>({
  groupQuestions: [],
  form: {
    title: "",
    description: null,
    thumbnailUrl: null,
    anyTime: true,
    startAt: null,
    endAt: null,
  },
})

/** list question */
const { data } = await graphqlQueryUseFetch(query)
pageGroupQuestion.value.groupQuestions = data?.groupQuestions || []

/** Make new question */
const submit = async () => {
  await graphqlQueryFetch(mutation, pageGroupQuestion.value.form)
  const { data } = await graphqlQueryFetch(query)
  pageGroupQuestion.value.groupQuestions = data?.groupQuestions || []
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
