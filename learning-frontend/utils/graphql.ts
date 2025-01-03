const defaultHeaders = {
  "Content-Type": "application/json",
  Accept: "application/json",
}

export async function graphqlQueryFetch(
  gql: any,
  variables?: Record<string, any>,
): Promise<any> {
  return await $fetch("http://localhost:8080/graphql", {
    method: "POST",
    headers: defaultHeaders,
    body: {
      query: gql?.loc?.source.body,
      variables: variables,
    },
  })
}

export async function graphqlQueryUseFetch(
  gql: any,
  variables?: Record<string, any>,
): Promise<any> {
  const { data, error } = await useFetch("http://localhost:8080/graphql", {
    method: "POST",
    headers: defaultHeaders,
    body: {
      query: gql?.loc?.source.body,
      variables: variables,
    },
  })

  if (error.value) {
    throw new Error(error.value.message)
  }

  return data.value
}

export async function graphqlUpload(gql: any, file: any): Promise<any> {
  const operations = JSON.stringify({
    query: gql?.loc?.source.body,
    variables: { file: null },
  })

  const formData = new FormData()
  const map = JSON.stringify({
    "0": ["variables.file"],
  })

  formData.append("operations", operations)
  formData.append("map", map)
  formData.append("0", file)

  const { data, error } = await useFetch("http://localhost:8080/graphql", {
    method: "POST",
    body: formData,
  })

  if (error.value) {
    throw new Error(error.value.message)
  }

  return data.value
}
