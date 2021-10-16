import { createClient } from 'urql';

const client = createClient({
  url: 'http://localhost:8080/query',
});

export default client;
