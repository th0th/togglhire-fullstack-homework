import classNames from 'classnames';
import React, { useState } from 'react';
import { gql, useMutation } from 'urql';
import styles from './ChoiceQuestionAnswer.module.css';

type Props = Overwrite<Omit<React.PropsWithoutRef<JSX.IntrinsicElements['div']>, 'children'>, {
  question: ChoiceQuestion,
}>;

const SetAnswer = gql`
  mutation ($questionID: ID!, $optionID: ID) {
    setAnswer (input: { questionID: $questionID, optionID: $optionID}) {
      id
    }
  }
`;

export default function ChoiceQuestionAnswer({ className, question, ...props }: Props) {
  const [, setAnswer] = useMutation(SetAnswer);
  const [optionId, setOptionId] = useState<string | null>(question.answer?.optionId || null);

  return (
    <div {...props} className={classNames(styles.choiceQuestionAnswer, className)}>
      <div className={styles.options}>
        {question.options.map((o) => (
          <div className={styles.option} key={o.id}>
            <label className={styles.label} htmlFor={`option-${o.id}`}>
              <input
                checked={optionId === o.id}
                className={styles.radioInput}
                id={`option-${o.id}`}
                onChange={async () => {
                  setOptionId(o.id);
                  await setAnswer({
                    questionID: question.id,
                    optionID: o.id,
                  });
                }}
                type="radio"
              />

              <div>{o.body}</div>
            </label>
          </div>
        ))}
      </div>
    </div>
  );
}
