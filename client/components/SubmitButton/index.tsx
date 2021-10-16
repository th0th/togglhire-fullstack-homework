import classNames from 'classnames';
import React from 'react';
import { ReactComponent as Check } from './check.svg';
import styles from './SubmitButton.module.css';

type Props = Omit<React.PropsWithoutRef<JSX.IntrinsicElements['button']>, 'children'>;

export default function SubmitButton({ className, ...props }: Props) {
  return (
    <button {...props} className={classNames(styles.submitButton, className)} type="submit">
      <Check className={styles.check} />

      Submit test
    </button>
  );
}
