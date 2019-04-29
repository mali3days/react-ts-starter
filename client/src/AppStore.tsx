import { createContext } from 'react';
import { observable, action } from 'mobx';
import axios from 'axios';

export interface Word {
  id: number;
  content: string;
  author_id: number;
  created_at: string;
}

export class Store {
  // @observable public words: Word[] = [];

  // @action
  // public getWords = async (): Promise<void> => {
  //   try {
  //     const { data } = await axios.get<{ data: { list: Word[] } }>(
  //       'http://localhost:8080/graphql?query={list{id,userId,title,completed}}',
  //     );
  //     this.words = data.data.list;
  //   } catch (error) {
  //     console.error(error);
  //   }
  // };

  // @action
  // public addWord = (word: string): void => {
  //   const newWord: Word = {
  //     userId: Number(new Date()),
  //     id: Number(new Date()),
  //     title: word,
  //     completed: false,
  //   };

  //   console.log(this.words);
  //   this.words.push(newWord);
  // };
}

export const RootStore = createContext(new Store());
