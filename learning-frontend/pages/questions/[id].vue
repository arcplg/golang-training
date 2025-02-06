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
            <!-- <ul>
              <li v-for="item,key in pageQuestion.exam?.questions">
                  <div>{{ item.text }}</div>
              </li>
            </ul> -->
        </div>
        <div class="right">
          <Questions
            v-for="item,key in pageQuestion.exam?.questions"
              :key="key"
              :index="key+1"
              :type="item.code.toString()"
              :data="item"
          />
          <div>
            <h2>Add question</h2>
            <form action="">
              <div v-for="template in pageQuestion.templates">
                <input type="radio" :id="template.code.toString()" :value="template" v-model="pageQuestion.questionTemplate" />
                <label for="one">{{template.code}}</label>
              </div>
              <div>
                <button @click.prevent="submit">Submit</button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
</template>
<script lang="ts" setup>
interface PageQuestion {
  id: String | String[],
  templates: QuestionTemplate[],
  exam: Exam | null
  questionTemplate: QuestionTemplate
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
  questionTemplate: {
    code: '',
    blocks: [],
    options: [],
  },
})

/** detail exam */
const { data } = await graphqlQueryUseFetch(queryDetailExam, {id: pageQuestion.value.id})
pageQuestion.value.exam = data?.findExam || null
pageQuestion.value.templates = data?.questionTemplates || null
if (!data) {
  throw createError({ statusCode: 404, statusMessage: 'Page not found' })
}

const submit = async () => {
  const res = await graphqlQueryFetch(mutationAddQuestionIntoExam, {id: pageQuestion.value.id, input: pageQuestion.value.questionTemplate})
   if(res?.errors) {
    alert('Errors')
    return;
  }
  const { data } = await graphqlQueryFetch(queryDetailExam, {id: pageQuestion.value.id})
  pageQuestion.value.exam = data?.findExam || null
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