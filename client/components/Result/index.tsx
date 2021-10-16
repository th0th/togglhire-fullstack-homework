import classNames from 'classnames';
import React from 'react';
import styles from './Result.module.css';

type Props = Overwrite<Omit<React.PropsWithoutRef<JSX.IntrinsicElements['div']>, 'children'>, {
  weight: number,
}>;

export default function Result({ className, weight, ...props }: Props) {
  return (
    <div {...props} className={classNames(styles.result, className)}>
      <p>You have completed!</p>

      <p>{`Your score: ${weight * 100}/100`}</p>
    </div>
  );
}
