import { Provider } from "urql";
import { client } from "../client";
import { Questions } from "./Questions";

export function App() {
  return (
    <Provider value={client}>
      <Questions />
    </Provider>
  );
}
