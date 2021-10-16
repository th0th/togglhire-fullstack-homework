import classNames from 'classnames';
import React from 'react';
import ChoiceQuestionAnswer from '../ChoiceQuestionAnswer';
import TextQuestionAnswer from '../TextQuestionAnswer';
import styles from './Question.module.css';

type Props = Overwrite<Omit<React.PropsWithoutRef<JSX.IntrinsicElements['div']>, 'children'>, {
  number: number,
  question: ChoiceQuestion | TextQuestion,
}>;

export default function Question({ className, number, question, ...props }: Props) {
  return (
    <div {...props} className={classNames(styles.question, className)}>
      <div className={styles.top}>
        <p className={styles.title}>
          {`QUESTION ${number} — `}

          {question.__typename === 'ChoiceQuestion' ? 'SINGLE CHOICE' : 'FREE TEXT'}
        </p>

        <p>{question.body}</p>
      </div>

      {question.__typename === 'ChoiceQuestion' ? (
        <ChoiceQuestionAnswer question={question} />
      ) : (
        <TextQuestionAnswer question={question} />
      )}
    </div>
  );
}
