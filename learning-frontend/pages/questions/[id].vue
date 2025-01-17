<template>
    <div>
      <h1>Questions detail id: {{ pageQuestion.id }}</h1>
      <div>
        <div>Title: {{ pageQuestion.exam?.title }}</div>
        <div>Description: {{ pageQuestion.exam?.description }}</div>
        <div>Thumbnail: {{ pageQuestion.exam?.thumbnailUrl }}</div>
        <div>Any Time: {{ pageQuestion.exam?.anyTime }} </div>
        <div>Start At: {{ pageQuestion.exam?.startAt }} </div>
        <div>End At: {{ pageQuestion.exam?.endAt }}</div>
        <div>Created At: {{ pageQuestion.exam?.createdAt }}</div>
        <div>Updated At: {{ pageQuestion.exam?.updatedAt }}</div>
      </div>
      <div class="inner">
        <div class="left">
            <ul>
              <li v-for="item,key in pageQuestion.questions">
                
              </li>
            </ul>
        </div>
        <div class="right">
          <Questions
            v-for="item,key in pageQuestion.templates"
              :key="key"
              :type="item.name.toString()"
              :data="item"
          />
        </div>
      </div>
    </div>
</template>
<script lang="ts" setup>
  import gql from "graphql-tag"
  const query = gql`
    query queryGql($id: String!) {
      findExam(_id: $id) {
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
      questionTemplates {
          _id
          name
          note
          text
          media {
              _id
              type
              url
          }
          blocks {
              _id
              text
              media {
                  _id
                  type
                  url
              }
          }
      }
    }
  `

  const mutation = gql`
    mutation AddQuestion(
      $id: String!,
      $originNumber: String!,
      $text: String!,
      $imageUrl: String
      $videoUrl: String
      $questionItems: Array
      $youtubeUrl: String
    ) {
      addQuestion(_id: $id, input: { 
        originNumber: $originNumber
        note: $note
        text: $text
        imageUrl: $imageUrl
        videoUrl: $videoUrl
        questionItems: $questionItems
        youtubeUrl: $youtubeUrl
      }) {
        originNumber
        note
        text
        imageUrl
        videoUrl
        questionItems
        youtubeUrl
    }
  }
  `

interface PageQuestion {
  id: String | String[],
  templates: QuestionTemplate[],
  exam: Exam | null
  questions: Question[]
  form: QuestionTemplate
}

const route = useRoute()
const pageQuestion = ref<PageQuestion>({
  id: route.params.id,
  templates: [],
  exam: {
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
    _id: null,
    name: '',
    note: null,
    text: null,
    media: null,
    blocks: null,
  },
})

/** list question */
const { data } = await graphqlQueryUseFetch(query, {id: pageQuestion.value.id})
pageQuestion.value.exam = data?.findExam || null
pageQuestion.value.templates = data?.questionTemplates || null
if (!data) {
  throw createError({ statusCode: 404, statusMessage: 'Page not found' })
}

const submit = async () => {
  await graphqlQueryFetch(mutation, pageQuestion.value.form)
  const { data } = await graphqlQueryFetch(query)
  pageQuestion.value.exam = data?.exam || null
}
</script>

<style lang="scss">
  .inner {
    display: flex;
    .left {
      display: flex;
      flex-direction: column;
      width: 25%;
      height: 100%;
      background-color: #ccc;
      padding: 15px;
    }
    .right {
      width: 75%;
      padding: 15px;
    }
  }
</style>