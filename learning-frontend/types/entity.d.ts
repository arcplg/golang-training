interface Media {
  id: String,
  type: String,
  url: String,
}
interface GroupQuestion {
  _id: String
  title: ?String
  description: ?String
  thumbnailUrl: ?String
  questions: Question[]
  answers: Answer[]
  anyTime: Boolean
  startAt: ?Date
  endAt: ?Date
  publishedAt: ?Date
  createdAt: ?Date
  createdBy: ?User
  updatedAt: ?Date
  deletedAt: ?Date
}

interface Answer {
  _id: String
  user: User
  lock: Boolean
  questions: [Question]
  startAt: Date
  endAt: Date
  publishedAt: Date
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
}

interface Question {
  _id: String
  originNumber: Int
  note: String
  text: String
  media: Media
  questionItems: [QuestionItem]
  publishedAt: Date
  createdAt: Date
  createdBy: User
  updatedAt: Date
  deletedAt: Date
}

interface QuestionItem {
  _id: String
  key: String
  originNumber: Int
  text: String
  media: Media
}

interface QuestionTemplate {
  _id: ?String
  originNumber: Int
  note: ?String
  text: ?String
  media: ?Media
  questionTemplateItems: [QuestionTemplateItem]
  publishedAt: Date
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
}

interface QuestionTemplateItem {
  _id: ?String
  key: ?String
  originNumber: Int
  text: ?String
  media: ?Media
}


/** Input */
interface QuestionInput {
  _id: ?String,
  templateKey: String,
  templateName: String,
  originNumber: Int
  text: ?String
  media: ?Media
  questionItems: QuestionItemInput[]
}
interface QuestionItemInput {
  _id: ?String,
  originNumber: Int
  optionKey: String,
  text: ?String
  media: ?Media
}
interface GroupQuestionInput {
  _id: ?String,
  title: String
  description: ?String
  thumbnailUrl: ?String
  anyTime: Boolean
  startAt: ?Date
  endAt: ?Date
}