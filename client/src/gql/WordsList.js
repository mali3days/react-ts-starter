import gql from 'graphql-tag';
import { useQuery } from 'react-apollo-hooks';

// query QRandomGiphy($tag: String!) {
//   giphy {
//     random(tag: $tag) {
//       images {
//         original {
//           width
//           height
//           url
//         }
//       }
//     }
//   }
// }

export const QWordsList = gql`
  query {
    words {
      id
      content
      created_at
    }
  }
`;

export function getWordsList() {
  const {
    data: { words = [] },
    error,
    refetch,
  } = useQuery(QWordsList);

  return {
    words,
    error,
    refresh: refetch,
  };
}
