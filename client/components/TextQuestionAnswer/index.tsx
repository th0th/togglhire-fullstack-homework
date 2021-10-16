import classNames from 'classnames';
import React, { useCallback, useState } from 'react';
import { gql, useMutation } from 'urql';
import debounce from 'lodash/debounce';
import styles from './TextQuestionAnswer.module.css';

type Props = Overwrite<Omit<React.PropsWithoutRef<JSX.IntrinsicElements['div']>, 'children'>, {
  question: TextQuestion,
}>;

const SetAnswer = gql`
  mutation ($questionID: ID!, $body: String) {
    setAnswer (input: { questionID: $questionID, body: $body}) {
      id
    }
  }
`;

export default function TextQuestionAnswer({ className, question, ...props }: Props) {
  const [, setAnswer] = useMutation(SetAnswer);
  const [body, setBody] = useState<string>(question.answer?.body || '');

  const mutate = useCallback(debounce(async (b: string) => {
    await setAnswer({
      body: b,
      questionID: question.id,
    });
  }, 1000), []);

  return (
    <div {...props} className={classNames(styles.textQuestionAnswer, className)}>
      <textarea
        className={styles.textarea}
        onChange={(event) => {
          setBody(event.target.value);
          mutate(event.target.value);
        }}
        required
        value={body}
      />
    </div>
  );
}
