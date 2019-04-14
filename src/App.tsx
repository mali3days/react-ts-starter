import * as React from 'react';
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

import { RootStore, Todo } from './AppStore';
import './App.css';

export interface Props {
  words: Todo[];
}
const RenderTodoList: React.FC<Props> = observer(
  ({ words }): JSX.Element => {
    console.log('RenderTodoList');

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
              <ListItemText primary={w.title} secondary={String(w.completed)} />
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

const TodoList = React.memo(RenderTodoList); // ?
// const TodoList = RenderTodoList;

export const App: React.FC<{}> = observer(
  (): JSX.Element => {
    const store = React.useContext(RootStore);
    const [word, setWord] = React.useState('');

    React.useEffect((): void => {
      store.getWords();
    }, []);

    const onSubmit = (e: React.SyntheticEvent<HTMLFormElement>): void => {
      e.preventDefault();
      console.log(word);
      store.addWord(word);
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
                <TodoList words={store.words} />
              </List>
            </div>
          </Grid>
        </div>
      </div>
    );
  },
);
