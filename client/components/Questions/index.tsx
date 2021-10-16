import { sortBy } from 'lodash';
import React from 'react';
import { gql, useMutation, useQuery } from 'urql';
import Question from '../Question';
import SubmitButton from '../SubmitButton';
import styles from './Questions.module.css';

type Query = {
  questions: Array<ChoiceQuestion | TextQuestion>,
};

const questionsQuery = gql`
  query {
    questions {
      ... on ChoiceQuestion {
        id
        body
        weight
        options {
          id
          body
          weight
        }
        answer {
          id
          optionId
          weight
        }
      }
      ... on TextQuestion {
        id
        body
        weight
        answer {
          id
          body
          weight
        }
      }
    }
  }
`;

const Submit = gql`
  mutation {
    submit {
      weight
    }
  }
`;

export default function Questions() {
  const [{ data }] = useQuery<Query>({ query: questionsQuery });
  const [submitResult, submit] = useMutation(Submit);

  return (
    <div className={styles.questions}>
      <form
        onSubmit={async (event) => {
          event.preventDefault();

          const result = await submit();

          if (result.error === undefined) {
            window.location.reload();
          }
        }}
      >
        <div className={styles.questionsWrapper}>
          {data === undefined ? null : (
            <>
              {sortBy(data.questions, ['weight']).map((q, i) => (
                <Question key={q.id} number={i + 1} question={q} />
              ))}
            </>
          )}
        </div>

        <div className={styles.submitButtonWrapper}>
          {submitResult?.error?.graphQLErrors === undefined ? null : (
            <div>
              {submitResult?.error?.graphQLErrors.map((e) => (
                <p key={e.message}>{e.message}</p>
              ))}
            </div>
          )}

          <SubmitButton />
        </div>
      </form>
    </div>
  );
}
