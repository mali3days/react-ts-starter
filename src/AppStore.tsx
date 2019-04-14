import { createContext } from 'react';
import { observable, action } from 'mobx';
import axios from 'axios';

export interface Todo {
  userId: number;
  id: number;
  title: string;
  completed: boolean;
  // type: any;
  // props: any;
  // key: any;
}

export class Store {
  @observable public words: Todo[] = [];

  @action
  public getWords = async (): Promise<void> => {
    try {
      const { data } = await axios.get<Todo[]>(
        'https://jsonplaceholder.typicode.com/todos?_limit=3',
      );
      this.words = data;
    } catch (error) {
      console.error(error);
    }
  };

  @action
  public addWord = (word: string): void => {
    const newWord: Todo = {
      userId: Number(new Date()),
      id: Number(new Date()),
      title: word,
      completed: false,
    };

    console.log(this.words);
    this.words.push(newWord);
  };
}

export const RootStore = createContext(new Store());
