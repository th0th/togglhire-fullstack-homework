import classNames from 'classnames';
import React from 'react';
import styles from './Header.module.css';

type Props = Overwrite<Omit<React.PropsWithoutRef<JSX.IntrinsicElements['header']>, 'children'>, {
  title: React.ReactNode,
}>;

export default function Header({ className, title, ...props }: Props) {
  return (
    <header {...props} className={classNames(styles.header, className)}>
      <img
        alt="Logo"
        className={styles.logo}
        src="https://s3.eu-central-1.amazonaws.com/production.images.hundred5.com/public/jid72cjd6fh5h0c34db396a825ddidlh42bj48ok5ljid72cjd6g1i48ab00d74171fcoboaf92eh46km95.png"
      />

      <div className={styles.center}>
        <h1 className={styles.title}>{title}</h1>
      </div>
    </header>
  );
}
