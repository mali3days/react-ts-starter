import gql from 'graphql-tag';
import { useQuery } from 'react-apollo-hooks';

export const DELETE_WORD_BY_ID = gql`
  mutation {
    deleteWord(id: $id) {
      id
    }
  }
`;
