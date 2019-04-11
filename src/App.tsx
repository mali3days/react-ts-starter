import * as React from 'react';
import './App.css';

export interface Props {
  blabla?: string;
}

export const App: React.FC<Props> = ({ blabla }: Props): JSX.Element => {
  const [name, setName] = React.useState(blabla);

  const onSubmit = (e: React.SyntheticEvent<HTMLFormElement>): void => {
    e.preventDefault();
  };

  const updateName = (e: React.SyntheticEvent<HTMLInputElement>): void => {
    setName(e.currentTarget.value);
  };

  return (
    <div className="App">
      <form noValidate={true} autoComplete="off" onSubmit={onSubmit}>
        <input type="text" value={name} onChange={updateName} />
      </form>
    </div>
  );
};
