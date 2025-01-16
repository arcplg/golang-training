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
  imageUrl: String
  videoUrl: String
  youtubeUrl: String
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
  imageUrl: String
  videoUrl: String
  youtubeUrl: String
}

interface QuestionTemplate {
  _id: String
  originNumber: Int
  note: String
  text: String
  imageUrl: String
  videoUrl: String
  youtubeUrl: String
  questionTemplateItems: [QuestionTemplateItem]
  publishedAt: Date
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
}

interface QuestionTemplateItem {
  _id: String
  key: String
  originNumber: Int
  text: String
  imageUrl: String
  videoUrl: String
  youtubeUrl: String
}

interface Media {
  type: String,
  url: String,
}

interface QuestionItemInput {
  key: String
  originNumber: Int
  text: String
  media: Media
}

interface QuestionInput {
  templateKey: String,
  templateName: String,
  originNumber: Int
  text: ?String
  media: ?Media
  questionItems: QuestionItemInput[]
}

interface GroupQuestionInput {
  title: String
  description: ?String
  thumbnailUrl: ?String
  anyTime: Boolean
  startAt: ?Date
  endAt: ?Date
}