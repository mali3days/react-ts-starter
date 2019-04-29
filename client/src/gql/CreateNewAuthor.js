import gql from 'graphql-tag';
import { useMutation } from 'react-apollo-hooks';

const CREATE_NEW_AUTHOR = gql`
  mutation {
    createAuthor(name: $id, email: $email) {
      id
      name
      email
    }
  }
`;

export const createNewAuthor = useMutation(CREATE_NEW_AUTHOR, {
  variables: { id: 1, email: 'emaaaial' },
});

// export function createNewAuthor({ id, email }) {
//   const { data, error, refetch } = useMutation(CREATE_NEW_AUTHOR, {
//     variables: { id, email },
//   });

//   console.log(data);

//   return {
//     data,
//     error,
//     refresh: refetch,
//   };
// }
