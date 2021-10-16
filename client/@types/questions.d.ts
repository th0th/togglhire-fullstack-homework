type Question = {
  body: string,
  id: string,
  weight: number,
};

type Option = {
  __typename: string,
  id: string,
  body: string,
  weight: number,
}

type ChoiceQuestion = Overwrite<Question, {
  __typename: 'ChoiceQuestion',
  options: Array<Option>,
  answer: ChoiceQuestionAnswer | null,
}>;

type TextQuestion = Overwrite<Question, {
  __typename: 'TextQuestion',
  answer: TextQuestionAnswer | null,
}>;

type Answer = {
  id: string,
  questionId: string,
};

type ChoiceQuestionAnswer = Overwrite<Answer, {
  optionId: string,
}>;

type TextQuestionAnswer = Overwrite<Answer, {
  body: string,
}>;
