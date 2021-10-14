# Toggl Hire full-stack homework

The goal of this assignment is to see how familiar you are with Go, JavaScript and React. We tried to pick a task that is similar to what you would do at Toggl Hire, while keeping it minimal so you can finish it in a short time.

You are provided with a simple application that consists of two projects. The server is a GraphQL API that returns a list of test questions. Questions can be either choice questions or text questions. Choice questions have a number of options, while text questions do not. You can see the full schema in [schema.graphqls](server/graph/schema.graphqls). The client is a React app that uses the API to fetch the questions and renders them as JSON. You can find the details in each project's README file ([server](server/README.md), [client](client/README.md)).

Your task is to expand this app to not only show the questions, but let the user answer them and record the answers on the server. Questions are answered by selecting one of the options (for choice questions) or by entering text (for text questions). To do this, you need to change the GraphQL schema on the server, implement all resolvers and then implement the client UI. 

Create a new repository on GitHub. You can [use this one as a template](https://github.com/togglhire/fullstack-homework/generate). Commit your solution to your repository and send us a link to it. If you prefer having the repository private, please add [miscer](https://github.com/miscer) as a collaborator, so we can review it.

## Basic requirements

Your solution should meet all these requirements.

### Server

- [x] Return at least one choice and one text question from the API

- [ ] Implement a mutation for submitting the answers.

  For each question, the selected option ID or the entered text is submitted.

- [ ] Validate the submitted answers (question and option IDs are correct, question types match, ...)

- [ ] Print the answers to stdout in a machine-readable format (JSON, CSV, ...)

  For simplicity we can assume that stdout is written to a log file, which is then later processed.

### Client

- [x] Show the questions and options to the user in some way.

- [ ] Show the questions and options sorted by their `weight`.

- [ ] Implement UI to answer choice and text questions.

    This can be similar to [how Toggl Hire does it](https://apply.hire.toggl.com/61VJ1QV7E1Z9ODJU2P2/practice). Note that there is no need to implement time limits or rating the questions.

- [ ] Send the answer to the API and show a success message if everything goes well.

## Bonus requirements

These requirements are not required, but feel free to complete some of them if they seem interesting, or to come up with your own :)

- [ ] Load questions from and save answers to a SQLite database.

- [ ] Evaluate the answers in some way and show a score to the candidate.

- [ ] Add automated tests for the server resolvers and client components.

## Additional notes

You can use any libraries and frameworks, but all dependencies should be defined in the package.json and go.mod files. The code should be formatted with Prettier and go fmt (or similar). 
