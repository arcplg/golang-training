interface Media {
  id: String,
  type: String,
  url: String,
}

interface Block {
  id: String,
  label: String,
  text: String,
  media: Media,
  block: Block[],
}
interface Exam {
  _id: String
  title: ?String
  description: ?String
  thumbnailUrl: ?String
  questions: ?Question[]
  answers: ?Answer[]
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
}

interface Question {
  _id: String
  name: String
  note: String
  text: String
  media: Media
  blocks: [Block]
  options: [Block]
  publishedAt: Date
  createdAt: Date
  createdBy: User
  updatedAt: Date
  deletedAt: Date
}

interface QuestionTemplate {
  _id: ?String
  name: String
  note: ?String
  text: ?String
  media: ?Media
  blocks: [Block]
  options: [Block]
}