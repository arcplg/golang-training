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
    <hr>
    <div v-for="(item, i) in pageGroupQuestion.exams" :key="i">
      <div>Id: {{ item._id }}</div>
      <div>Title: {{ item.title }}</div>
      <div>Description: {{ item.description }}</div>
      <div>Thumbnail: {{ item.thumbnailUrl }}</div>
      <div>Any Time: {{ item.anyTime }}</div>
      <div>Start At: {{ item.startAt }}</div>
      <div>End At: {{ item.endAt }}</div>
      <div>Created At: {{ item.createdAt }}</div>
      <div>Created At:{{ item.updatedAt }}</div>
      <div>
        <NuxtLink :to="'questions/'+item._id">
          Detail
        </NuxtLink>
      </div>
      <hr>
    </div>

    
  </div>
</template>

<script lang="ts" setup>
import gql from "graphql-tag"

/** variable */
const query = gql`
  query Exams {
    exams {
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
  mutation CreateExam(
      $title: String!,
      $description: String,
      $thumbnailUrl: String
      $anyTime: Boolean
      $startAt: DateTime
      $endAt: DateTime
    ) {
      createExam(input: { 
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
  exams: Exam[]
  form: Exam
}
const pageGroupQuestion = ref<PageGroupQuestion>({
  exams: [],
  form: {
    _id: "",
    title: null,
    description: null,
    thumbnailUrl: null,
    anyTime: true,
    startAt: null,
    endAt: null,
    questions: [],
    answers: [],
    publishedAt:null,
    createdAt:null,
    createdBy: null,
    updatedAt:null,
    deletedAt:null,
  },
})

/** list question */
const { data } = await graphqlQueryUseFetch(query)
pageGroupQuestion.value.exams = data?.exams || []

/** Make new question */
const submit = async () => {
  const res = await graphqlQueryFetch(mutation, pageGroupQuestion.value.form)
  if(res?.errors) {
    alert('Errors')
    return;
  }
  const { data } = await graphqlQueryFetch(query)
  pageGroupQuestion.value.exams = data?.exams || []
}

</script>

<style>
</style>
