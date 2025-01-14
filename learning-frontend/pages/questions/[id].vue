<template>
    <div>
      <h1>Questions detail id: {{ pageQuestion.id }}</h1>
      <div>
        <div>Title: {{ pageQuestion.groupQuestion.title }}</div>
        <div>Description: {{ pageQuestion.groupQuestion.description }}</div>
        <div>Thumbnail: {{ pageQuestion.groupQuestion.thumbnailUrl }}</div>
        <div>Any Time: {{ pageQuestion.groupQuestion.anyTime }} </div>
        <div>Start At: {{ pageQuestion.groupQuestion.startAt }} </div>
        <div>End At: {{ pageQuestion.groupQuestion.endAt }}</div>
        <div>Created At: {{ pageQuestion.groupQuestion.createdAt }}</div>
        <div>Updated At: {{ pageQuestion.groupQuestion.updatedAt }}</div>
      </div>
    </div>
</template>
<script lang="ts" setup>
  import gql from "graphql-tag"
  const query = gql`
    query FindGroupQuestion($id: String!) {
      findGroupQuestion(_id: $id) {
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

interface PageQuestion {
  id: String | String[],
  groupQuestion: GroupQuestion
  questions: Question[]
  form: QuestionInput
}

const route = useRoute()
const pageQuestion = ref<PageQuestion>({
  id: route.params.id,
  groupQuestion: {
    _id: "",
    title: null,
    description: null,
    thumbnailUrl: null,
    questions: [],
    answers: [],
    anyTime: true,
    startAt: null,
    endAt: null,
    publishedAt: null,
    createdAt: null,
    createdBy: null,
    updatedAt: null,
    deletedAt: null,
  },
  questions: [],
  form: {
    originNumber: 1,
    note: null,
    text: null,
    imageUrl: null,
    videoUrl: null,
    youtubeUrl: null,
    questionItems: [],
  },
})

/** list question */
const { data } = await graphqlQueryUseFetch(query, {id: pageQuestion.value.id})
pageQuestion.value.groupQuestion = data?.findGroupQuestion || []
if (!data) {
  throw createError({ statusCode: 404, statusMessage: 'Page not found' })
}
</script>