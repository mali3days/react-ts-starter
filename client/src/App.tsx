import * as React from 'react';
import gql from 'graphql-tag';
import { useMutation } from 'react-apollo-hooks';
// import { observer, inject } from 'mobx-react';
import { observer } from 'mobx-react-lite';
import TextField from '@material-ui/core/TextField';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemAvatar from '@material-ui/core/ListItemAvatar';
import ListItemSecondaryAction from '@material-ui/core/ListItemSecondaryAction';
import ListItemText from '@material-ui/core/ListItemText';
import Avatar from '@material-ui/core/Avatar';
import IconButton from '@material-ui/core/IconButton';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import FolderIcon from '@material-ui/icons/Folder';
import DeleteIcon from '@material-ui/icons/Delete';

import { getWordsList, QWordsList } from './gql/WordsList';

import { RootStore, Word } from './AppStore';
import './App.css';

export interface Props {
  words: Word[];
}
const RenderTodoList: React.FC<Props> = observer(
  ({ words }): JSX.Element => {
    console.log('RenderWordList');

    // https://stackoverflow.com/questions/54679118/jsx-element-type-element-is-not-a-constructor-function-for-jsx-elements
    // https://github.com/DefinitelyTyped/DefinitelyTyped/issues/20356#issuecomment-336384210
    return (
      <>
        {words.map(
          (w): JSX.Element => (
            <ListItem key={w.id}>
              <ListItemAvatar>
                <Avatar>
                  <FolderIcon />
                </Avatar>
              </ListItemAvatar>
              <ListItemText primary={w.content} secondary={String(w.created_at)} />
              <ListItemSecondaryAction>
                <IconButton aria-label="Delete">
                  <DeleteIcon />
                </IconButton>
              </ListItemSecondaryAction>
            </ListItem>
          ),
        )}
      </>
    );
  },
);

const CREATE_NEW_WORD = gql`
  mutation createWord($content: String!, $authorId: Int!) {
    createWord(content: $content, author_id: $authorId) {
      id
      content
    }
  }
`;

const TodoList = React.memo(RenderTodoList); // ?
// const TodoList = RenderTodoList;

export const App: React.FC<{}> = observer(
  (): JSX.Element => {
    const [word, setWord] = React.useState('');
    const data = getWordsList();
    const createNewWord = useMutation(CREATE_NEW_WORD);

    const onSubmit = (e: React.SyntheticEvent<HTMLFormElement>): void => {
      e.preventDefault();
      createNewWord({
        variables: { content: word, authorId: 1 },
        refetchQueries: [{ query: QWordsList }],
      });
      setWord('');
    };

    const updateName = (e: React.ChangeEvent<HTMLInputElement>): void => {
      setWord(e.currentTarget.value);
    };

    return (
      <div className="App">
        <form noValidate={true} autoComplete="off" onSubmit={onSubmit}>
          <TextField
            id="standard-full-width"
            label="New word"
            placeholder="Watermelon"
            helperText="Enter new word and press `Enter`"
            margin="normal"
            InputLabelProps={{
              shrink: true,
            }}
            value={word}
            onChange={updateName}
          />
        </form>
        <div
          style={{
            display: 'flex',
            justifyContent: 'center',
            margin: '30px',
          }}
        >
          <Grid item xs={12} md={6}>
            <Typography variant="h6">Words to learn: </Typography>
            <div>
              <List dense>
                <TodoList words={data.words} />
              </List>
            </div>
          </Grid>
        </div>
      </div>
    );
  },
);
