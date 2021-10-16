import React from 'react';
import { gql, useQuery } from 'urql';
import { ActivityIndicator, Header, Questions } from '../components';
import Result from '../components/Result';
import styles from '../styles/Index.module.css';

const query = gql`
  query {
    result {
      weight
    }
  }
`;

export default function Index() {
  const [{ data }] = useQuery({ query });

  return (
    <div className={styles.index}>
      <Header title="Homework" />

      {data === undefined ? (
        <div className={styles.activityIndicatorWrapper}>
          <ActivityIndicator />
        </div>
      ) : (
        <>
          {data.result.length === 0 ? (
            <Questions />
          ) : (
            <Result weight={data.result[0].weight} />
          )}
        </>
      )}
    </div>
  );
}
