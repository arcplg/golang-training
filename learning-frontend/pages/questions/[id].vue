<template>
    <div>
      <h1>Questions detail id: {{ pageQuestion.id }}</h1>
      <div>
        <div>Title: {{ pageQuestion.groupQuestion?.title }}</div>
        <div>Description: {{ pageQuestion.groupQuestion?.description }}</div>
        <div>Thumbnail: {{ pageQuestion.groupQuestion?.thumbnailUrl }}</div>
        <div>Any Time: {{ pageQuestion.groupQuestion?.anyTime }} </div>
        <div>Start At: {{ pageQuestion.groupQuestion?.startAt }} </div>
        <div>End At: {{ pageQuestion.groupQuestion?.endAt }}</div>
        <div>Created At: {{ pageQuestion.groupQuestion?.createdAt }}</div>
        <div>Updated At: {{ pageQuestion.groupQuestion?.updatedAt }}</div>
      </div>
      <div class="inner">
        <div class="left">
            <ul>
              <li>1) 記帳時に費目・項目が二重に表示されているユーザーがいます。</li>
              <li>2) 記帳時に費目・項目が二重に表示されているユーザーがいます。</li>
              <li>3) 記帳時に費目・項目が二重に表示されているユーザーがいます。</li>
              <li>4) 記帳時に費目・項目が二重に表示されているユーザーがいます。</li>
              <li>5) 記帳時に費目・項目が二重に表示されているユーザーがいます。</li>
              <li>6) 記帳時に費目・項目が二重に表示されているユーザーがいます。</li>
              <li>7) 記帳時に費目・項目が二重に表示されているユーザーがいます。</li>
            </ul>
        </div>
        <div class="right">
          <Questions
            v-for="item,key in pageQuestion.templates"
              :key="key"
              :type="item.templateKey.toString()"
              :data="item"
          />
        </div>
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
  templates: QuestionInput[],
  groupQuestion: GroupQuestion | null
  questions: Question[]
  form: QuestionInput
}

const route = useRoute()
const pageQuestion = ref<PageQuestion>({
  id: route.params.id,
  templates: [
    {
      templateKey: 'T10001',
      templateName: 'T10001',
      originNumber: 1,
      text: null,
      media: null,
      questionItems: [],
    },
    {
      templateKey: 'T10002',
      templateName: 'T10002',
      originNumber: 1,
      text: null,
      media: null,
      questionItems: [],
    },
  ],
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
    templateKey: 'basic',
    templateName: 'Basic',
    originNumber: 1,
    text: null,
    media: null,
    questionItems: [],
  },
})

/** list question */
const { data } = await graphqlQueryUseFetch(query, {id: pageQuestion.value.id})
pageQuestion.value.groupQuestion = data?.findGroupQuestion || null
if (!data) {
  throw createError({ statusCode: 404, statusMessage: 'Page not found' })
}

const submit = async () => {
  await graphqlQueryFetch(mutation, pageQuestion.value.form)
  const { data } = await graphqlQueryFetch(query)
  pageQuestion.value.groupQuestion = data?.groupQuestion || null
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