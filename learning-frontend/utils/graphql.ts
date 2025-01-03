
const defaultHeaders = {
  "Content-Type": "application/json",
  Accept: "application/json",
}

export async function graphqlQueryFetch(
  gql: any,
  variables?: Record<string, any>
): Promise<any> {
  return await $fetch('http://localhost:8080/graphql', {
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
  variables?: Record<string, any>
): Promise<any> {

  const {data, error} = await useFetch('http://localhost:8080/graphql', {
    method: "POST",
    headers: defaultHeaders,
    body: {
      query: gql?.loc?.source.body,
      variables: variables,
    },
  })

  if (error.value) {
    throw new Error(error.value.message);
  }

  return data.value;
}