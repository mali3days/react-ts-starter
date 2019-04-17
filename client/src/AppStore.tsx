import { createContext } from 'react';
import { observable, action } from 'mobx';
import axios from 'axios';

export interface Todo {
  userId: number;
  id: number;
  title: string;
  completed: boolean;
}

export class Store {
  @observable public words: Todo[] = [];

  @action
  public getWords = async (): Promise<void> => {
    try {
      const { data } = await axios.get<{ data: { list: Todo[] } }>(
        'http://localhost:8080/graphql?query={list{id,userId,title,completed}}',
      );
      this.words = data.data.list;
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
