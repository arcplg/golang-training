import gql from "graphql-tag"

export const queryDetailExam = gql`
    query queryGql($id: String!) {
    findExam(_id: $id) {
        _id
        title
        description
        thumbnailUrl
        anyTime
        startAt
        endAt
        publishedAt
        createdAt
        updatedAt
        deletedAt
        questions {
            _id
            code
            blocks {
                _id
                label
                text
                media {
                    _id
                    type
                    url
                }
            }
            options {
                _id
                text
            }
            correctOption {
                _id
                text
            }
            createdAt
            updatedAt
            deletedAt
        }
    }
    questionTemplates {
        code
        blocks {
            label
            text
            media {
                type
                url
            },
            blocks {
                label
                text
                media {
                    type
                    url
                },
            }
        }
        options {
            label
            text
            media {
                type
                url
            },
            blocks {
                label
                text
                media {
                    type
                    url
                },
            }
        }
    }
    }
`

export const mutationAddQuestionIntoExam = gql`
    mutation AddQuestionIntoExam(
    $id: String!,
    $input: QuestionInput!
    ) {
        addQuestionIntoExam(id: $id, input: $input) {
            _id
        }
    }
`