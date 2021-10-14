import { gql, useQuery } from "urql";

export function Questions() {
  const [{ data, fetching, error }] = useQuery({ query });

  if (fetching) return "Loading...";
  if (error) return `Error: ${error}`;

  return (
    <code>
      <pre>{JSON.stringify(data, null, 2)}</pre>
    </code>
  );
}

const query = gql`
  query {
    questions {
      __typename
      ... on ChoiceQuestion {
        id
        body
        weight
        options {
          id
          body
          weight
        }
      }
      ... on TextQuestion {
        id
        body
        weight
      }
    }
  }
`;
