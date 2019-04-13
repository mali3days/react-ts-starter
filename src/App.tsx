import * as React from 'react';
import TextField from '@material-ui/core/TextField';
import { createStyles, Theme, withStyles, WithStyles } from '@material-ui/core/styles';

import './App.css';

// export interface Props {
//   // blabla?: string;
// }

export const App: React.FC<{}> = (): JSX.Element => {
  const [name, setName] = React.useState('');

  const onSubmit = (e: React.SyntheticEvent<HTMLFormElement>): void => {
    e.preventDefault();
  };

  const updateName = (e: React.ChangeEvent<HTMLInputElement>): void => {
    setName(e.currentTarget.value);
  };

  return (
    <div className="App">
      <form noValidate={true} autoComplete="off" onSubmit={onSubmit}>
        <TextField
          id="standard-full-width"
          label="New word"
          style={{ margin: 8 }}
          placeholder="Watermelon"
          helperText="Enter new word and press `Enter`"
          margin="normal"
          InputLabelProps={{
            shrink: true,
          }}
          value={name}
          onChange={updateName}
        />
      </form>
    </div>
  );
};
